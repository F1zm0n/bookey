package entity

import "github.com/google/uuid"

type TransactionDTO struct {
	FromUserID       uuid.UUID `json:"sender"`
	FromUserPassword string    `json:"receivers_password"`
	BookID           uuid.UUID `json:"book_id"`
	Description      string    `json:"description"`
}

//отправка
