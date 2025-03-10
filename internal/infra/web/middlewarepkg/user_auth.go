package middlewarepkg

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"tasks-api/internal/auth/jwtpkg"
)

type AuthKey struct{}

func AuthenticationMiddlewareFunc(tokenMaker jwtpkg.TokenServiceInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)
			if err != nil {
				http.Error(w, fmt.Sprintf("error verifying token: %v", err), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), AuthKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthorizeMiddlewareFunc(roleID int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(AuthKey{}).(*jwtpkg.UserClaims)
			if !ok {
				http.Error(w, "claims not found", http.StatusInternalServerError)
				return
			}

			if claims.RoleID != roleID {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker jwtpkg.TokenServiceInterface) (*jwtpkg.UserClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := fields[1]
	claims, err := tokenMaker.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
