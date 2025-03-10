package repository

import (
	"github.com/stretchr/testify/mock"
	"tasks-api/internal/entity"
)

type TechnicianRepositoryMock struct {
	mock.Mock
}

func (t *TechnicianRepositoryMock) NewTask(input entity.TaskEntity) (int, error) {
	args := t.Called(input)
	return args.Int(0), args.Error(1)
}

func (t *TechnicianRepositoryMock) FindTask(taskID, userID int) (*entity.TaskEntity, error) {
	args := t.Called(taskID, userID)
	return args.Get(0).(*entity.TaskEntity), args.Error(1)
}

func (t *TechnicianRepositoryMock) CountTasksByUser(userID int) (int, error) {
	args := t.Called(userID)
	return args.Int(0), args.Error(1)
}

func (t *TechnicianRepositoryMock) UpdateTask(input entity.TaskEntity) (int, error) {
	args := t.Called(input)
	return args.Int(0), args.Error(1)
}

func (t *TechnicianRepositoryMock) AllTasksByUser(userID int, filter entity.PaginationFilter) (*[]entity.TaskEntity, error) {
	args := t.Called(userID, filter)
	return args.Get(0).(*[]entity.TaskEntity), args.Error(1)
}
