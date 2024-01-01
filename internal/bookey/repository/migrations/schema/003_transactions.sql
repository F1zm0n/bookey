-- +goose Up

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    from_user_id UUID NOT NULL,
    to_user_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    description VARCHAR(255) NOT NULL ,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (from_user_id) REFERENCES users(user_id),
    FOREIGN KEY (to_user_id) REFERENCES users(user_id)
);
CREATE INDEX idx_from ON transactions (from_user_id);

CREATE INDEX idx_to ON transactions (to_user_id);

CREATE INDEX idx_from_to ON transactions (from_user_id, to_user_id);
-- +goose Down
DROP TABLE transactions;