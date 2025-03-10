package usecase

import (
	"errors"
	"tasks-api/internal/infra/repository"
	"testing"
)

func TestManagerDeleteTaskUseCase_Execute(t *testing.T) {
	testsSuites := []struct {
		name     string
		isFailed bool
	}{
		{
			name:     "failed to delete task",
			isFailed: true,
		},
		{
			name:     "success",
			isFailed: false,
		},
	}

	repoMock := new(repository.ManagerRepositoryMock)

	for _, tt := range testsSuites {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isFailed {
				repoMock.On("DeleteTask", 1, 1).Return(errors.New("error"))
				repoMock.On("CountTasks").Return(errors.New("error"))
			} else {
				repoMock.On("DeleteTask", 1, 1).Return(nil)
				repoMock.On("CountTasks").Return(nil)
			}

			uc := ManagerDeleteTaskUseCase{
				ManagerRepository: repoMock,
			}

			err := uc.Execute(ManagerDeleteTaskInputDTO{ID: 1}, 1)

			if tt.isFailed && err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}
