package dto

import (
	"database/sql"
)

type User struct {
	Username     string `json:"username"`
	Email        string `bson:"email"`
	PasswordHash string `json:"password_hash"`
}

type Book struct {
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Description sql.NullString `json:"description"`
	Genre       sql.NullString `json:"genre"`
	Price       string         `json:"price"`
}
