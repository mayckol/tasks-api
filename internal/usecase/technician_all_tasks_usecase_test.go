package usecase

import (
	"errors"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/repository"
	"testing"
)

func TestTechnicianAllTasksUseCase_Execute(t *testing.T) {
	testsSuites := []struct {
		name     string
		isFailed bool
	}{
		{
			name:     "failed to get all tasks",
			isFailed: true,
		},
		{
			name:     "failed to count tasks",
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
			if tt.isFailed {
				repoMock.On("AllTasksByUser", 1, 1).Return(&[]entity.TaskEntity{}, errors.New("error"))
				repoMock.On("CountTasksByUser", 1).Return(0, errors.New("error"))
			} else {
				repoMock.On("AllTasksByUser", 1, 1).Return(&[]entity.TaskEntity{}, nil)
				repoMock.On("CountTasksByUser", 1).Return(0, nil)
			}

			uc := TechnicianAllTasksUseCase{
				TechnicianRepository: repoMock,
			}

			_, err := uc.Execute(TechnicianAllTasksInputDTO{
				UserID: 1,
				Page:   1,
			}, 1)

			if tt.isFailed && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
