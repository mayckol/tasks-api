package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"tasks-api/configs"
	"tasks-api/internal/auth/jwtpkg"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/internal/infra/notify"
	"tasks-api/internal/infra/presenter"
	"tasks-api/internal/infra/web/middlewarepkg"
	"tasks-api/internal/usecase"
	"tasks-api/internal/validation"
)

type TechnicianHandler struct {
	envs                 *configs.EnvVars
	technicianRepository entity.TechnicianRepository
	validator            validation.ValidatorInterface
	notifyService        notify.NotifyInterface
}

func NewTechnicianHandler(
	envs *configs.EnvVars,
	technicianRepository entity.TechnicianRepository,
	validator validation.ValidatorInterface,
	notifyService notify.NotifyInterface,
) *TechnicianHandler {
	return &TechnicianHandler{
		envs:                 envs,
		technicianRepository: technicianRepository,
		validator:            validator,
		notifyService:        notifyService,
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
// @Param task_id path int true "Task ID"
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

	taskID := chi.URLParam(r, "id")
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
		NotifyService:        a.notifyService,
	}

	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONPresenter(w, http.StatusCreated, output, nil)
}

// AllTasks godoc
// @Summary all tasks.
// @Description all tasks.
// @Tags Technician
// @Accept */*
// @Produce json
// @Param page query string false "page"
// @Success 200 {object} usecase.TechnicianAllTasksOutputDTO
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 404 {string} {object} "not found"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /technician/task [get]
func (a *TechnicianHandler) AllTasks(w http.ResponseWriter, r *http.Request) {
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

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.TechnicianAllTasksInputDTO
	input.UserID = userID
	input.Page = p

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.TechnicianAllTasksUseCase{
		TechnicianRepository: a.technicianRepository,
	}

	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONSingleResPresenter(w, http.StatusOK, output)
}

// FindTask godoc
// @Summary find task.
// @Description find task.
// @Tags Technician
// @Accept */*
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 200 {object} usecase.TechnicianFindTaskOutputDTO
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

	taskID := chi.URLParam(r, "id")
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

	presenter.JSONSingleResPresenter(w, http.StatusOK, output)
}
