-- name: CreateTransaction :one
INSERT INTO transactions (transaction_id,from_user_id,to_user_id,amount,description,transaction_date)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;
-- name: GetUsersTransactions :many
SELECT * FROM transactions WHERE to_user_id=$1 OR from_user_id=$1;

