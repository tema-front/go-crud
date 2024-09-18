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

-- name: EditUser :one
UPDATE
  users
SET
  name = $1,
  updated_at = $2
WHERE
  id = $3
RETURNING *;

-- name: GetUsers :many
SELECT
  *
FROM 
  users
WHERE 
  $1::text IS NULL OR name ILIKE '%' || $1 || '%'
ORDER BY 
  CASE 
    WHEN $2 = 'ASC' THEN name 
    WHEN $2 = 'DESC' THEN name
    ELSE NULL 
  END;

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