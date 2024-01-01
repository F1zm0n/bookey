package entity

import (
	"bookey/internal/bookey/repository/migrations/sqlc"
)

func BookToBookDTO(book *sqlc.Book) *BookDTO {
	NewBookDTO := &BookDTO{
		BookID:      book.BookID,
		UserID:      book.UserID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Genre:       book.Genre,
		Price:       book.Price,
		CreatedAt:   book.CreatedAt,
	}
	return NewBookDTO
}
