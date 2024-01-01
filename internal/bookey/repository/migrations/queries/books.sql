-- name: CreateBook :one
INSERT INTO Books (book_id,title,author,description,genre,price,user_id,created_at)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;
-- name: GetBooksByTitle :many
SELECT * FROM Books WHERE title=$1
ORDER BY created_at DESC
LIMIT $2;
-- name: GetBooksByAuthor :many
SELECT * FROM Books WHERE author=$1
ORDER BY created_at DESC
LIMIT $2;
-- name: GetBooksByGenre :many
SELECT * FROM Books WHERE genre=$1
ORDER BY created_at DESC
LIMIT $2;
-- name: GetBooksByUserID :many
SELECT * FROM Books WHERE user_id=$1
ORDER BY created_at DESC
LIMIT $2;
-- name: DeleteBookByID :exec
DELETE FROM Books WHERE book_id=$1;

