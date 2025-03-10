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
WHERE id = ? and user_id = ? and deleted_at is null;

-- name: AllTasksByUser :many
SELECT id,
       user_id,
       summary,
       is_done,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE user_id = ?
  AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: AllTasks :many
SELECT id,
       user_id,
       summary,
       is_done,
       updated_by,
       created_at,
       updated_at
FROM tasks
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountTasksByUser :one
SELECT COUNT(*) as total
FROM tasks
WHERE user_id = ?
  AND deleted_at IS NULL;

-- name: CountTasks :one
SELECT COUNT(*) as total
FROM tasks
WHERE deleted_at IS NULL;


-- name: UpdateTask :execresult
UPDATE tasks
SET updated_at = now(),
    summary = ?,
    is_done = ?,
    updated_by = ?
WHERE id = ? and deleted_at is null;

-- name: DeleteTask :execresult
UPDATE tasks
SET deleted_at = now(),
    updated_at = now(),
    updated_by = ?
WHERE id = ? and deleted_at is null;