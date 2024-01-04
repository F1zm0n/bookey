package usecase

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func (u *Usecase) CreateBook(ctx *gin.Context, book *entity.BookDTO) (sqlc.Book, error) {
	bookForFunc := &entity.Book{
		BookID:      uuid.New(),
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Genre:       book.Genre,
		Price:       book.Price,
		UserID:      book.UserID,
		CreatedAt:   time.Now(),
	}

	newBook, err := u.Repository.CreateBook(ctx, bookForFunc)
	if err != nil {
		return sqlc.Book{}, err
	}
	return newBook, err
}

func (u *Usecase) GetBooksByAuthor(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error) {
	Books, err := u.Repository.GetBooksByAuthor(ctx, book.Author)
	if err != nil {
		return nil, err
	}
	return Books, nil
}

func (u *Usecase) GetBooksByGenre(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error) {
	Books, err := u.Repository.GetBooksByGenre(ctx, book.Genre)
	if err != nil {
		return nil, err
	}
	return Books, nil
}

func (u *Usecase) GetBooksByUserID(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error) {
	Books, err := u.Repository.GetBooksByUserID(ctx, book.UserID)
	if err != nil {
		return nil, err
	}
	return Books, nil
}
func (u *Usecase) GetBooksByTitle(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error) {
	Books, err := u.Repository.GetBooksByTitle(ctx, book.Title)
	if err != nil {
		return nil, err
	}
	return Books, nil
}
