-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: AuthByToken :one
SELECT
  *
FROM 
 users
WHERE
 api_key = $1;

-- name: GetUsers :many
SELECT
  id, 
  name,
  created_at,
  updated_at
FROM 
  users;

-- name: GetUser :one
SELECT
  *
FROM 
 users
WHERE
 id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ClearUsers :exec
DELETE FROM users;