-- name: GetUserById :one
SELECT * FROM users u WHERE u.id = $1;