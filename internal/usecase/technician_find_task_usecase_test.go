package usecase

import (
	"errors"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/repository"
	"testing"
)

func TestTechnicianFIndTaskUseCase_Execute(t *testing.T) {
	testsSuites := []struct {
		name     string
		isFailed bool
	}{
		{
			name:     "failed to get task",
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
				repoMock.On("FindTask", 1, 1).Return(&entity.TaskEntity{}, errors.New("error"))
			} else {
				repoMock.On("FindTask", 1, 1).Return(&entity.TaskEntity{}, nil)
			}

			uc := TechnicianFindTaskUseCase{
				TechnicianRepository: repoMock,
			}

			_, err := uc.Execute(TechnicianFindTaskInputDTO{
				ID: 1,
			}, 1)

			if tt.isFailed && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
