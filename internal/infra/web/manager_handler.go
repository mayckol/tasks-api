package web

import (
	"github.com/go-chi/chi/v5"
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
	"tasks-api/utils"
)

type ManagerHandler struct {
	envs              *configs.EnvVars
	managerRepository entity.ManagerRepository
	validator         validation.ValidatorInterface
}

func NewManagerHandler(
	envs *configs.EnvVars,
	managerRepository entity.ManagerRepository,
	validator validation.ValidatorInterface,
) *ManagerHandler {
	return &ManagerHandler{
		envs:              envs,
		managerRepository: managerRepository,
		validator:         validator,
	}
}

// AllTasks godoc
// @Summary all tasks.
// @Description all tasks.
// @Tags Manager
// @Accept */*
// @Produce json
// @Param page  query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} usecase.ManagerAllTasksOutputDTO
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 403 {string} {object} "forbidden"
// @Failure 404 {string} {object} "not found"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /manager/task [get]
func (a *ManagerHandler) AllTasks(w http.ResponseWriter, r *http.Request) {
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
	limit := r.URL.Query().Get("limit")

	pagFilter, err := utils.PaginationFilterByQueryParams(page, limit)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.ManagerAllTasksInputDTO
	input.Page = pagFilter.Page
	input.Limit = pagFilter.Limit

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.ManagerAllTasksUseCase{
		ManagerRepository: a.managerRepository,
	}

	output, appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONSingleResPresenter(w, http.StatusOK, output)
}

// DeleteTask godoc
// @Summary delete task.
// @Description delete task.
// @Tags Manager
// @Accept */*
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 204
// @Failure 400 {string} {object} "invalid request"
// @Failure 401 {string} {object} "unauthorized"
// @Failure 403 {string} {object} "forbidden"
// @Failure 404 {string} {object} "not found"
// @Failure 500 {string} string "internal server error"
// @Security ApiKeyAuth
// @Router /manager/task/{task_id} [delete]
func (a *ManagerHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
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

	taskID := chi.URLParam(r, "task_id")
	if taskID == "" {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	tID, err := strconv.Atoi(taskID)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.ManagerDeleteTaskInputDTO
	input.ID = tID

	invalidFields, isFailure := a.validator.Validate(input)
	if isFailure {
		presenter.JSONPresenter(w, http.StatusBadRequest, invalidFields, errorpkg.ValidateFieldsError)
		return
	}

	uc := usecase.ManagerDeleteTaskUseCase{
		ManagerRepository: a.managerRepository,
	}

	appError := uc.Execute(input, userID)
	if appError != nil {
		presenter.JSONPresenter(w, appError.StatusCode, nil, appError)
		return
	}

	presenter.JSONPresenter(w, http.StatusNoContent, nil)
}
