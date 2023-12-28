-- +goose Up
CREATE TABLE Books (
    book_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    genre VARCHAR(25) NOT NULL ,
    price DECIMAL(10, 2) NOT NULL,
    user_id UUID REFERENCES Users(user_id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE Books;