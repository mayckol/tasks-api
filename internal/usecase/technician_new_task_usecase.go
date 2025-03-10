package usecase

import (
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
)

// swagger:model TechnicianNewTaskInputDTO
type TechnicianNewTaskInputDTO struct {
	Summary string `json:"summary" validate:"required,min=3,max=2500"`
}

// swagger:model TechnicianNewTaskOutputDTO
type TechnicianNewTaskOutputDTO struct {
	TaskID int `json:"task_id"`
}

type TechnicianNewTaskUseCase struct {
	TechnicianRepository entity.TechnicianRepository
}

func (n *TechnicianNewTaskUseCase) Execute(input TechnicianNewTaskInputDTO, userID int) (*TechnicianNewTaskOutputDTO, *errorpkg.AppError) {
	id, err := n.TechnicianRepository.NewTask(entity.TaskEntity{
		UserID:    userID,
		Summary:   input.Summary,
		UpdatedBy: userID,
	})

	if err != nil {
		return nil, errorpkg.Wrap("failed to create new technician", http.StatusInternalServerError, err)
	}

	return &TechnicianNewTaskOutputDTO{
		TaskID: id,
	}, nil
}
