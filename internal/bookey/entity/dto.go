package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"` //binding:"required min=4,max=30"
	Email    string    `json:"email"`    //validate:"required,email"
	Password string    `json:"password"` //binding:"required min=8,max=64"
	Balance  int32     `json:"balance"`
}

type BookDTO struct {
	BookID      uuid.UUID `json:"book_id"`
	UserID      uuid.UUID `json:"user_id" ` //binding:"required"
	Title       string    `json:"title" `   //binding:"required"
	Author      string    `json:"author" `  //binding:"required"
	Description string    `json:"description"`
	Genre       string    `json:"genre" ` //binding:"required"
	Price       int32     `json:"price" ` //binding:"required gte=0 lte=1000000000"
	CreatedAt   time.Time `json:"created_at"`
}
type UserID struct {
	UserID uuid.UUID `json:"user_id"`
}
