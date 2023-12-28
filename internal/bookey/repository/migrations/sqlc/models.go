// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	BookID      uuid.UUID
	Title       string
	Author      string
	Description string
	Genre       string
	Price       string
	UserID      uuid.NullUUID
}

type Transaction struct {
	TransactionID   uuid.UUID
	UserID          uuid.UUID
	Amount          string
	Description     string
	TransactionDate time.Time
}

type User struct {
	UserID       uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	Balance      string
}
