-- name: CreateUser :one
INSERT INTO Users (user_id,username,email,password_hash,balance)
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM Users WHERE user_id=$1 AND password_hash=$2;

-- name: GetUser :one
SELECT * FROM Users WHERE user_id=$1 AND password_hash=$2;







