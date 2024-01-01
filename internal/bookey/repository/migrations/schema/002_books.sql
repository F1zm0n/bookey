-- +goose Up
CREATE TABLE Books (
    book_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    genre VARCHAR(25) NOT NULL ,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL REFERENCES Users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL
);
CREATE INDEX idx_author ON Books (author);
CREATE INDEX idx_genre ON Books (genre);
CREATE INDEX idx_title ON Books (title);
CREATE INDEX idx_book_user_id ON Books (user_id);

-- +goose Down
DROP TABLE Books;
-- DROP INDEX idx_author;
-- DROP INDEX idx_genre;
-- DROP INDEX idx_title;
-- DROP INDEX idx_book_user_id;