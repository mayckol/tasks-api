package repository

import (
	"context"
	"database/sql"
	"tasks-api/internal/entity"
	"tasks-api/internal/infra/database/queries"
)

type ManagerRepository struct {
	q *queries.Queries
}

func NewManagerRepository(q *queries.Queries) *ManagerRepository {
	return &ManagerRepository{q: q}
}

func (m ManagerRepository) DeleteTask(taskId, updatedBy int) error {
	_, err := m.q.DeleteTask(context.Background(), queries.DeleteTaskParams{
		ID:        int32(taskId),
		DeletedBy: sql.NullInt32{Int32: int32(updatedBy), Valid: true},
	})
	return err
}

func (m ManagerRepository) AllTasks(filter entity.PaginationFilter) (*[]entity.TaskEntity, error) {
	offset := (filter.Page - 1) * filter.Limit
	task, err := m.q.AllTasks(context.Background(), queries.AllTasksParams{
		Limit:  int32(filter.Limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var tasks []entity.TaskEntity
	for _, t := range task {
		newTask := entity.TaskEntity{
			ID:        int(t.ID),
			UserID:    int(t.UserID),
			Summary:   t.Summary,
			IsDone:    t.IsDone,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			UpdatedBy: int(t.UpdatedBy),
		}
		if t.PerformedAt.Valid {
			newTask.PerformedAt = &t.PerformedAt.Time
		}
		tasks = append(tasks, newTask)
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
