package web

import (
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
// @Param page query string false "page"
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
	if page == "" {
		page = "1"
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		presenter.JSONPresenter(w, http.StatusBadRequest, nil, errorpkg.ParseJsonError)
		return
	}

	var input usecase.ManagerAllTasksInputDTO
	input.Page = p

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
