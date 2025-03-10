package usecase

import (
	"errors"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/repository"
	"testing"
)

func TestTechnicianUpdateTaskUseCase_Execute(t *testing.T) {
	testsSuites := []struct {
		name     string
		isFailed bool
	}{
		{
			name:     "failed to update task",
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
			repoMock.On("FindTask", 1, 1).Return(&entity.TaskEntity{
				ID: 1,
			}, nil)

			var updateTask entity.TaskEntity
			updateTask.UserID = 1
			updateTask.Summary = "summary"
			updateTask.UpdatedBy = 1
			if tt.isFailed {
				repoMock.On("UpdateTask", updateTask).Return(0, errors.New("error"))
			} else {
				repoMock.On("UpdateTask", updateTask).Return(1, nil)
			}

			uc := TechnicianUpdateTaskUseCase{
				TechnicianRepository: repoMock,
			}

			_, err := uc.Execute(TechnicianUpdateTaskInputDTO{
				Summary: updateTask.Summary,
				TaskID:  1,
			}, 1)

			if tt.isFailed && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
