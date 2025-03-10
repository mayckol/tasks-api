package usecase

import (
	"errors"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/repository"
	"testing"
)

func TestTechnicianNewTaskUseCase_Execute(t *testing.T) {
	testsSuites := []struct {
		name     string
		isFailed bool
	}{
		{
			name:     "failed to create task",
			isFailed: true,
		},
		{
			name:     "success",
			isFailed: false,
		},
	}

	repoMock := new(repository.TechnicianRepositoryMock)

	for _, tt := range testsSuites {
		t.Run(tt.name, func(t *testing.T) {
			var newTask entity.TaskEntity
			newTask.UserID = 1
			newTask.Summary = "summary"
			newTask.UpdatedBy = 1
			if tt.isFailed {
				repoMock.On("NewTask", newTask).Return(0, errors.New("error"))
			} else {
				repoMock.On("NewTask", newTask).Return(1, nil)
			}

			uc := TechnicianNewTaskUseCase{
				TechnicianRepository: repoMock,
			}

			_, err := uc.Execute(TechnicianNewTaskInputDTO{
				Summary: newTask.Summary,
			}, 1)

			if tt.isFailed && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
