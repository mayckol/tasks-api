package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
)

// swagger:model ManagerDeleteTaskInputDTO
type ManagerDeleteTaskInputDTO struct {
	ID int `json:"id,omitempty"`
}

type ManagerDeleteTaskUseCase struct {
	ManagerRepository entity.ManagerRepository
}

func (n *ManagerDeleteTaskUseCase) Execute(input ManagerDeleteTaskInputDTO, userID int) *errorpkg.AppError {
	err := n.ManagerRepository.DeleteTask(input.ID, userID)
	if err != nil {
		return errorpkg.Wrap("failed to delete task", http.StatusInternalServerError, err)
	}

	return nil
}
