package usecase

import (
	"fmt"
	"net/http"
	"tasks-api/internal/entity"
	"tasks-api/internal/errorpkg"
	"tasks-api/internal/infra/notify"
)

// swagger:model TechnicianUpdateTaskInputDTO
type TechnicianUpdateTaskInputDTO struct {
	Summary string `json:"summary,omitempty" validate:"omitempty,min=1,max=255"`
	IsDone  bool   `json:"is_done,omitempty" validate:"omitempty"`
	TaskID  int    `json:"task_id" validate:"required"`
}

// swagger:model TechnicianUpdateTaskOutputDTO
type TechnicianUpdateTaskOutputDTO struct {
	TaskID int `json:"task_id"`
}

type TechnicianUpdateTaskUseCase struct {
	TechnicianRepository entity.TechnicianRepository
	NotifyService        notify.NotifyInterface
}

func (n *TechnicianUpdateTaskUseCase) Execute(input TechnicianUpdateTaskInputDTO, userID int) (*TechnicianUpdateTaskOutputDTO, *errorpkg.AppError) {
	t, err := n.TechnicianRepository.FindTask(input.TaskID, userID)
	if err != nil {
		return nil, errorpkg.Wrap("failed to find task", http.StatusInternalServerError, err)
	}

	if t.UserID != userID {
		return nil, errorpkg.Wrap("task not found", http.StatusNotFound, nil)
	}

	t.Summary = input.Summary
	t.IsDone = input.IsDone
	t.UpdatedBy = userID

	_, err = n.TechnicianRepository.UpdateTask(*t)

	if err != nil {
		return nil, errorpkg.Wrap("failed to create update technician", http.StatusInternalServerError, err)
	}

	if t.IsDone {
		go func() {
			err := n.NotifyService.TaskPerformed(t.ID, userID)
			if err != nil {
				fmt.Printf("failed to notify task performed: %v\n", err)
			}
		}()

	}

	return &TechnicianUpdateTaskOutputDTO{
		TaskID: t.ID,
	}, nil
}
