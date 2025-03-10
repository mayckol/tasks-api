package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/utils"
)

// swagger:model TechnicianAllTasksInputDTO
type TechnicianAllTasksInputDTO struct {
	UserID int `json:"user_id" validate:"required"`
	Page   int `json:"page,omitempty" validate:"min=1"`
	Limit  int `json:"limit,omitempty" validate:"min=1"`
}

// swagger:model TechnicianAllTasksOutputDTO
type TechnicianAllTasksOutputDTO struct {
	Data  []entity.TaskEntity `json:"data"`
	Page  int                 `json:"page,omitempty"`
	Total int                 `json:"total,omitempty"`
}

type TechnicianAllTasksUseCase struct {
	TechnicianRepository entity.TechnicianRepository
}

func (n *TechnicianAllTasksUseCase) Execute(input TechnicianAllTasksInputDTO, userID int) (*TechnicianAllTasksOutputDTO, *errorpkg.AppError) {
	paginationFilter := entity.PaginationFilter{
		Page:  input.Page,
		Limit: input.Limit,
	}

	utils.HandlePagination(&paginationFilter)

	tasks, err := n.TechnicianRepository.AllTasksByUser(userID, paginationFilter)
	if err != nil {
		return nil, errorpkg.Wrap("failed to find tasks", http.StatusInternalServerError, err)
	}

	count, err := n.TechnicianRepository.CountTasksByUser(userID)
	if err != nil {
		return nil, errorpkg.Wrap("failed to find tasks", http.StatusInternalServerError, err)
	}

	return &TechnicianAllTasksOutputDTO{
		Data:  *tasks,
		Total: count,
		Page:  input.Page,
	}, nil
}
