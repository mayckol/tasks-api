-- name: StoreUser :execresult
INSERT INTO users (first_name, email, password, role_id) VALUES (?, ?, ?, ?);

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?;