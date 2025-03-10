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

func (t TechnicianRepository) UpdateTask(input entity.TaskEntity) (int, error) {
	res, err := t.q.UpdateTask(context.Background(), queries.UpdateTaskParams{
		Summary:   input.Summary,
		UpdatedBy: int32(input.UpdatedBy),
		IsDone:    input.IsDone,
		ID:        int32(input.ID),
	})

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (t TechnicianRepository) AllTasksByUser(userID, page int) (*[]entity.TaskEntity, error) {
	// this is a hardcoded value, but it should be dynamic
	const pageSize = 10
	offset := (page - 1) * pageSize
	task, err := t.q.AllTasksByUser(context.Background(), queries.AllTasksByUserParams{
		UserID: int32(userID),
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

func (t TechnicianRepository) CountTasksByUser(userID int) (int, error) {
	count, err := t.q.CountTasksByUser(context.Background(), int32(userID))
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (t TechnicianRepository) FindTask(taskID, userID int) (*entity.TaskEntity, error) {
	task, err := t.q.FindTaskByID(context.Background(), queries.FindTaskByIDParams{
		ID:     int32(taskID),
		UserID: int32(userID),
	})
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
