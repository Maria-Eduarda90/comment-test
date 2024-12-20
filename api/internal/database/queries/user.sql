-- name: GetUserById :one
SELECT * FROM users u WHERE u.id = $1;

-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindUserByEmail :one
SELECT u.id, u.name, u.email FROM users u WHERE u.email = $1;

-- name: FindUserByID :one
SELECT id, name, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    updated_at = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: FindManyUsers :many
SELECT id, name, email, created_at, updated_at 
FROM users
ORDER BY created_at DESC;

-- name: UpdatePassword :exec
UPDATE users SET password = $2, updated_at = $3 WHERE id = $1;

-- name: GetUserPassword :one
SELECT u.password FROM users u WHERE u.id = $1;
