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
	Price       float64
	UserID      uuid.UUID
}

type Transaction struct {
	TransactionID   uuid.UUID
	UserID          string
	Amount          string
	Description     string
	TransactionDate time.Time
}

type User struct {
	UserID       uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	Balance      float64
}
