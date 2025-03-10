package repository

import (
	"context"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/database/queries"
)

type ManagerRepository struct {
	q *queries.Queries
}

func NewManagerRepository(q *queries.Queries) *ManagerRepository {
	return &ManagerRepository{q: q}
}

func (m ManagerRepository) DeleteTask(input entity.TaskEntity) error {
	_, err := m.q.DeleteTask(context.Background(), queries.DeleteTaskParams{
		ID: int32(input.ID),
	})
	return err
}

func (m ManagerRepository) AllTasks(page int) (*[]entity.TaskEntity, error) {
	// this is a hardcoded value, but it should be dynamic
	const pageSize = 10
	offset := (page - 1) * pageSize
	task, err := m.q.AllTasks(context.Background(), queries.AllTasksParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var tasks []entity.TaskEntity
	for _, t := range task {
		tasks = append(tasks, entity.TaskEntity{
			ID:        int(t.ID),
			UserID:    int(t.UserID),
			Summary:   t.Summary,
			IsDone:    t.IsDone,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			UpdatedBy: int(t.UpdatedBy),
		})
	}

	return &tasks, nil
}

func (m ManagerRepository) CountTasks() (int, error) {
	count, err := m.q.CountTasks(context.Background())
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
