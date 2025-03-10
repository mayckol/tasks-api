// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: task.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const findTaskByID = `-- name: FindTaskByID :one
SELECT id,
       user_id,
       summary,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE id = ? and user_id = ? and deleted_at is null
`

type FindTaskByIDParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
}

type FindTaskByIDRow struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Summary   string    `json:"summary"`
	UpdatedBy int32     `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) FindTaskByID(ctx context.Context, arg FindTaskByIDParams) (FindTaskByIDRow, error) {
	row := q.queryRow(ctx, q.findTaskByIDStmt, findTaskByID, arg.ID, arg.UserID)
	var i FindTaskByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Summary,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findTasksByUserID = `-- name: FindTasksByUserID :many
SELECT id,
       user_id,
       summary,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE user_id = ? and deleted_at is null
`

type FindTasksByUserIDRow struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Summary   string    `json:"summary"`
	UpdatedBy int32     `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) FindTasksByUserID(ctx context.Context, userID int32) ([]FindTasksByUserIDRow, error) {
	rows, err := q.query(ctx, q.findTasksByUserIDStmt, findTasksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindTasksByUserIDRow
	for rows.Next() {
		var i FindTasksByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Summary,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const storeTask = `-- name: StoreTask :execresult
INSERT INTO tasks (user_id,
                   summary,
                   updated_by)
VALUES (?,
        ?,
        ?)
`

type StoreTaskParams struct {
	UserID    int32  `json:"user_id"`
	Summary   string `json:"summary"`
	UpdatedBy int32  `json:"updated_by"`
}

func (q *Queries) StoreTask(ctx context.Context, arg StoreTaskParams) (sql.Result, error) {
	return q.exec(ctx, q.storeTaskStmt, storeTask, arg.UserID, arg.Summary, arg.UpdatedBy)
}

const updateTask = `-- name: UpdateTask :execresult
UPDATE tasks
SET updated_at = now(),
    summary = ?,
    is_done = ?,
    updated_by = ?
WHERE id = ? and deleted_at is null
`

type UpdateTaskParams struct {
	Summary   string `json:"summary"`
	IsDone    bool   `json:"is_done"`
	UpdatedBy int32  `json:"updated_by"`
	ID        int32  `json:"id"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (sql.Result, error) {
	return q.exec(ctx, q.updateTaskStmt, updateTask,
		arg.Summary,
		arg.IsDone,
		arg.UpdatedBy,
		arg.ID,
	)
}
