package entity

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	BookID      uuid.UUID
	Title       string
	Author      string
	Description string
	Genre       string
	Price       int32
	UserID      uuid.UUID
	CreatedAt   time.Time
}

type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	FromUserID    uuid.UUID `json:"sender"`
	ToUserID      uuid.UUID `json:"receiver"`
	Sum           float64   `json:"amount"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
}

type User struct {
	UserID       uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	Balance      int32
}
