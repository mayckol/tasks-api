package repository

import (
	"context"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/database/queries"
)

type TechnicianRepository struct {
	q *queries.Queries
}

func NewTechnicianRepository(q *queries.Queries) *TechnicianRepository {
	return &TechnicianRepository{q: q}
}

func (t TechnicianRepository) NewTask(input entity.TaskEntity) (int, error) {
	res, err := t.q.StoreTask(context.Background(), queries.StoreTaskParams{
		UserID:    int32(input.UserID),
		Summary:   input.Summary,
		UpdatedBy: int32(input.UserID),
	})

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (t TechnicianRepository) FindTask(taskID int) (*entity.TaskEntity, error) {
	task, err := t.q.FindTaskByID(context.Background(), int32(taskID))
	if err != nil {
		return nil, err
	}

	return &entity.TaskEntity{
		ID:        int(task.ID),
		UserID:    int(task.UserID),
		Summary:   task.Summary,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		UpdatedBy: int(task.UpdatedBy),
	}, nil
}

func (t TechnicianRepository) UpdateTask(taskID int, task entity.TaskEntity) error {
	_, err := t.q.UpdateTask(context.Background(), queries.UpdateTaskParams{
		Summary:   task.Summary,
		UpdatedBy: int32(task.UpdatedBy),
		ID:        int32(taskID),
	})

	return err
}
