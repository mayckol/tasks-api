-- name: StoreTask :execresult
INSERT INTO tasks (user_id,
                   summary,
                   updated_by)
VALUES (?,
        ?,
        ?);

-- name: FindTasksByUserID :many
SELECT id,
       user_id,
       summary,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE user_id = ? and deleted_at is null;

-- name: FindTaskByID :one
SELECT id,
       user_id,
       summary,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE id = ? and deleted_at is null;

-- name: UpdateTask :execresult
UPDATE tasks
SET summary = ?,
    updated_by = ?
WHERE id = ?;