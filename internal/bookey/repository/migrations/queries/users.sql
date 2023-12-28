-- name: CreateUser :one
INSERT INTO Users (user_id,username,email,password_hash,balance)
VALUES ($1, $2, $3, $4,$5)
RETURNING *;




