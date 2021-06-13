-- name: CreateUser :one
INSERT INTO users (
  name, pwd
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserName :one
Update users
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
Update users
SET pwd = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;