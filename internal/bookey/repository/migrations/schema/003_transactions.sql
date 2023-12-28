-- +goose Up

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    description VARCHAR(255) NOT NULL ,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- +goose Down
DROP TABLE transactions;