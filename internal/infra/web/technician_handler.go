package web

import (
	"encoding/json"
	"net/http"
	"tasks-api/configs"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/internal/infra/presenter"
	"tasks-api/internal/infra/web/middlewarepkg"
	"tasks-api/internal/usecase"
	"tasks-api/internal/validation"
)

// TechnicianHandler is a struct that contains the database connection and the environment variables it also can be AuthHandler
// I chose to name it TechnicianHandler only to focus on the functional requirements of the task
type TechnicianHandler struct {
	envs                 *configs.EnvVars
	technicianRepository entity.TechnicianRepository
	validator            validation.ValidatorInterface
}

func NewTechnicianHandler(
	envs *configs.EnvVars,
	technicianRepository entity.TechnicianRepository,
	validator validation.ValidatorInterface,
) *TechnicianHandler {
	return &TechnicianHandler{
		envs:                 envs,
		technicianRepository: technicianRepository,
		validator:            validator,
	}
}

// Task godoc
// @Summary new task.
// @Description new task.
// @Tags Technician
// @Accept */*
// @Produce json
// @Param request body usecase.TechnicianNewTaskInputDTO true "task"
// @Success 201 {object} usecase.TechnicianNewTaskOutputDTO
// @Failure 400 {string} {object} "invalid request"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /technician/task [post]
func (a *TechnicianHandler) Task(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewarepkg.AuthKey{}).(*jwtpkg.UserClaims)
	if !ok {
		presenter.JSONPresenter(w, http.StatusInternalServerError, nil, invalidTokenClaimsErr)
		return
	}

	userID := claims.UserID
	if userID == 0 {
		presenter.JSONPresenter(w, http.StatusInternalServerError, nil, invalidTokenClaimsErr)
		return
	}

	var input usecase.TechnicianNewTaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.TechnicianNewTaskUseCase{
		TechnicianRepository: a.technicianRepository,
	}
	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONPresenter(w, http.StatusCreated, output, nil)
}
