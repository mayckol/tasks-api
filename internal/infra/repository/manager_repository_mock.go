package repository

import (
	"github.com/stretchr/testify/mock"
	"tasks-api/internal/entity"
)

type ManagerRepositoryMock struct {
	mock.Mock
}

func (t *ManagerRepositoryMock) DeleteTask(taskId, updatedBy int) error {
	args := t.Called(taskId, updatedBy)
	return args.Error(0)
}

func (t *ManagerRepositoryMock) AllTasks(page entity.PaginationFilter) (*[]entity.TaskEntity, error) {
	args := t.Called(page)
	return args.Get(0).(*[]entity.TaskEntity), args.Error(1)
}

func (t *ManagerRepositoryMock) CountTasks() (int, error) {
	args := t.Called()
	return args.Int(0), args.Error(1)
}
