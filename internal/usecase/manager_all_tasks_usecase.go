package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
)

// swagger:model ManagerAllTasksInputDTO
type ManagerAllTasksInputDTO struct {
	Page int `json:"page,omitempty"`
}

// swagger:model ManagerAllTasksOutputDTO
type ManagerAllTasksOutputDTO struct {
	Data  []entity.TaskEntity `json:"data"`
	Page  int                 `json:"page,omitempty"`
	Total int                 `json:"total,omitempty"`
}

type ManagerAllTasksUseCase struct {
	ManagerRepository entity.ManagerRepository
}

func (n *ManagerAllTasksUseCase) Execute(input ManagerAllTasksInputDTO, userID int) (*ManagerAllTasksOutputDTO, *errorpkg.AppError) {
	tasks, err := n.ManagerRepository.AllTasks(input.Page)
	if err != nil {
		return nil, errorpkg.Wrap("failed to create find manager", http.StatusInternalServerError, err)
	}

	count, err := n.ManagerRepository.CountTasks()
	if err != nil {
		return nil, errorpkg.Wrap("failed to create find manager", http.StatusInternalServerError, err)
	}

	return &ManagerAllTasksOutputDTO{
		Data:  *tasks,
		Total: count,
		Page:  input.Page,
	}, nil
}
