package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"tasks-api/configs"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/notify"
	"tasks-api/internal/infra/repository"
	"tasks-api/internal/infra/web/middlewarepkg"
	"tasks-api/internal/validation"
	"testing"
)

func newRequestWithClaims(body string, claims *jwtpkg.UserClaims) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/technician/task", strings.NewReader(body))
	if claims != nil {
		ctx := context.WithValue(req.Context(), middlewarepkg.AuthKey{}, claims)
		req = req.WithContext(ctx)
	}
	return req
}

// func (a *TechnicianHandler) Task(w http.ResponseWriter, r *http.Request) {
func TestTechnicianHandler_Task(t *testing.T) {
	h := &TechnicianHandler{
		envs:                 &configs.EnvVars{},
		technicianRepository: new(repository.TechnicianRepositoryMock),
		validator:            validation.NewWrapper(),
		notifyService:        new(notify.SimpleNotifier),
	}

	tests := []struct {
		name           string
		requestBody    string
		claims         *jwtpkg.UserClaims
		validatorFail  bool
		repository     entity.TechnicianRepository
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "missing JWT claims",
			requestBody:    `{"summary": "test task"}`,
			claims:         nil,
			validatorFail:  false,
			repository:     h.technicianRepository,
			expectedStatus: http.StatusInternalServerError,
			expectedError:  errors.New("invalid token claims"),
		},
		{
			name:        "over then max length (2500) summary",
			requestBody: randomSummary(2501),
			claims: &jwtpkg.UserClaims{
				UserID: 1,
				RoleID: 1,
				RegisteredClaims: jwt.RegisteredClaims{
					ID: "1",
				},
			},
			validatorFail:  false,
			repository:     h.technicianRepository,
			expectedStatus: http.StatusNotAcceptable,
			expectedError:  errors.New("failed to validate fields"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := newRequestWithClaims(tt.requestBody, tt.claims)
			w := httptest.NewRecorder()

			h.Task(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status code %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedError != nil {
				if !strings.Contains(w.Body.String(), tt.expectedError.Error()) {
					t.Errorf("expected error %v, got %s", tt.expectedError, w.Body.String())
				}
			}
		})
	}
}

func randomSummary(length int) string {
	const allowedCharacters = "ABC123"
	r := rand.New(rand.NewSource(int64(rand.Intn(500))))
	b := make([]byte, length)
	for i := range b {
		b[i] = allowedCharacters[r.Intn(len(allowedCharacters))]
	}
	s := string(b)
	return fmt.Sprintf(`{"summary": "%s"}`, s)
}
