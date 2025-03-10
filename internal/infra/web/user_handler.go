package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"tasks-api/configs"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/internal/infra/presenter"
	"tasks-api/internal/usecase"
	"tasks-api/internal/validation"
)

var invalidTokenClaimsErr = errors.New("invalid token claims")

// UserHandler is a struct that contains the database connection and the environment variables it also can be AuthHandler
// I chose to name it UserHandler only to focus on the functional requirements of the task
type UserHandler struct {
	envs           *configs.EnvVars
	userRepository entity.UserRepository
	JWTService     jwtpkg.TokenServiceInterface
	validator      validation.ValidatorInterface
}

func NewUserHandler(
	envs *configs.EnvVars,
	userRepository entity.UserRepository,
	jwtService jwtpkg.TokenServiceInterface,
	validator validation.ValidatorInterface,
) *UserHandler {
	return &UserHandler{
		envs:           envs,
		userRepository: userRepository,
		JWTService:     jwtService,
		validator:      validator,
	}
}

// NewUser godoc
// @Summary new user.
// @Description new user.
// @Tags User
// @Accept */*
// @Produce json
// @Param request body usecase.NewUserInputDTO true "user"
// @Success 201 {string} string "created"
// @Failure 400 {string} {object} "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /user [post]
func (a *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var input usecase.NewUserInputDTO
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

	uc := usecase.NewUserUseCase{
		UserRepository: a.userRepository,
	}
	appError := uc.Execute(input)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONPresenter(w, http.StatusCreated, nil, nil)
}

// Signin godoc
// @Summary signin.
// @Description signin.
// @Tags User
// @Accept */*
// @Produce json
// @Param request body usecase.SigninInputDTO true "signin"
// @Success 200 {object} usecase.SigninOutputDTO "signin"
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /user/signin [post]
func (a *UserHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var input usecase.SigninInputDTO
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

	uc := usecase.SigninUseCase{
		UserRepository: a.userRepository,
		JWTService:     a.JWTService,
	}
	output, appError := uc.Execute(input)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		Value:  output.AccessToken,
		Path:   "/",
		Secure: true,
	})

	// I put the access token and role_id in the response body for testing purposes, the values would be stored only in the cookie in a real-world scenario
	presenter.JSONPresenter(w, http.StatusOK, output, nil)
}
