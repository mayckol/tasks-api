package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"time"
)

// swagger:model TechnicianFindTaskInputDTO
type TechnicianFindTaskInputDTO struct {
	ID int `json:"id" validate:"required"`
}

// swagger:model TechnicianFindTaskOutputDTO
type TechnicianFindTaskOutputDTO struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Summary   string    `json:"summary"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by,omitempty"`
}

type TechnicianFindTaskUseCase struct {
	TechnicianRepository entity.TechnicianRepository
}

func (n *TechnicianFindTaskUseCase) Execute(input TechnicianFindTaskInputDTO, userID int) (*TechnicianFindTaskOutputDTO, *errorpkg.AppError) {
	task, err := n.TechnicianRepository.FindTask(input.ID, userID)
	if err != nil {
		return nil, errorpkg.Wrap("failed to create find technician", http.StatusInternalServerError, err)
	}

	if task == nil {
		return nil, errorpkg.Wrap("task not found", http.StatusNotFound, nil)
	}

	if task.UserID != userID {
		return nil, errorpkg.Wrap("not allowed", http.StatusUnauthorized, nil)
	}

	return &TechnicianFindTaskOutputDTO{
		ID:        task.ID,
		UserID:    task.UserID,
		Summary:   task.Summary,
		IsDone:    task.IsDone,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		UpdatedBy: task.UpdatedBy,
	}, nil
}
