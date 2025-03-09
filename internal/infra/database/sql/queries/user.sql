-- name: StoreUser :execresult
INSERT INTO users (first_name, email, password, role_id) VALUES (?, ?, ?, ?);

-- name: DeleteUser :execresult
UPDATE users SET deleted_at = NOW() WHERE id = ?;

-- name: UserByEmail :one
SELECT * FROM users WHERE email = ?;