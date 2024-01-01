-- +goose Up

CREATE TABLE Users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    balance INTEGER NOT NULL
);
CREATE INDEX idx_user_id ON Users (user_id);
-- +goose Down

DROP TABLE Users;
-- DROP INDEX idx_user_id;
