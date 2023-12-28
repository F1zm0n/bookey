-- +goose Up

CREATE TABLE Users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL
);

-- +goose Down

DROP TABLE Users;