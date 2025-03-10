package web

import (
	"encoding/json"
	"net/http"
	"strconv"
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

// UpdateTask godoc
// @Summary update task.
// @Description update task.
// @Tags Technician
// @Accept */*
// @Produce json
// @Param task_id query string true "task id"
// @Param request body usecase.TechnicianUpdateTaskInputDTO true "task"
// @Success 201 {object} usecase.TechnicianUpdateTaskOutputDTO
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 404 {string} {object} "not found"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /technician/task/{task_id} [patch]
func (a *TechnicianHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
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

	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.NotFoundError)
		return
	}

	tID, err := strconv.Atoi(taskID)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.TechnicianUpdateTaskInputDTO
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	input.TaskID = tID

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.TechnicianUpdateTaskUseCase{
		TechnicianRepository: a.technicianRepository,
	}

	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONPresenter(w, http.StatusCreated, output, nil)
}

// FindTask godoc
// @Summary find task.
// @Description find task.
// @Tags Technician
// @Accept */*
// @Produce json
// @Param task_id query string true "task id"
// @Success 201 {object} usecase.TechnicianFindTaskOutputDTO
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 404 {string} {object} "not found"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /technician/task/{task_id} [get]
func (a *TechnicianHandler) FindTask(w http.ResponseWriter, r *http.Request) {
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

	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.NotFoundError)
		return
	}

	tID, err := strconv.Atoi(taskID)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.TechnicianFindTaskInputDTO
	input.ID = tID

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.TechnicianFindTaskUseCase{
		TechnicianRepository: a.technicianRepository,
	}

	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONSingleResPresenter(w, http.StatusOK, output, nil)
}
