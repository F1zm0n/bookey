package database

import (
	"bookey/internal/bookey/entity"
	"bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const limit = 100

func (r *Repository) CreateBook(ctx *gin.Context, book *entity.Book) (sqlc.Book, error) {
	r.Logger.Info("creating new book")
	User, err := r.DBTX.CreateBook(ctx, sqlc.CreateBookParams{
		BookID:      book.BookID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Genre:       book.Genre,
		Price:       book.Price,
		UserID:      book.UserID,
		CreatedAt:   book.CreatedAt,
	})
	if err != nil {
		r.Logger.Error("error creating book: ", err)
		return User, err
	}
	r.Logger.Info("book created successfully")
	return User, nil
}

func (r *Repository) GetBooksByAuthor(ctx *gin.Context, author string) ([]sqlc.Book, error) {
	r.Logger.Info("getting books by author")
	books, err := r.DBTX.GetBooksByAuthor(ctx, sqlc.GetBooksByAuthorParams{
		Author: author,
		Limit:  limit,
	})
	if err != nil {
		r.Logger.Error("error getting books by author ", err)
		return nil, err
	}
	r.Logger.Info("success getting books by author")
	return books, nil
}

func (r *Repository) GetBooksByGenre(ctx *gin.Context, genre string) ([]sqlc.Book, error) {
	r.Logger.Info("getting books by genre")
	books, err := r.DBTX.GetBooksByGenre(ctx, sqlc.GetBooksByGenreParams{
		Genre: genre,
		Limit: limit,
	})
	if err != nil {
		r.Logger.Error("error getting books by genre ", err)
		return nil, err
	}
	r.Logger.Info("success getting books by genre")
	return books, nil
}

func (r *Repository) GetBooksByUserID(ctx *gin.Context, ID uuid.UUID) ([]sqlc.Book, error) {
	r.Logger.Info("getting books by user id")
	books, err := r.DBTX.GetBooksByUserID(ctx, sqlc.GetBooksByUserIDParams{
		UserID: ID,
		Limit:  limit,
	})
	if err != nil {
		r.Logger.Error("error getting books by user id ", err)
		return nil, err
	}
	r.Logger.Info("success getting books by user id")
	return books, nil
}

func (r *Repository) GetBooksByTitle(ctx *gin.Context, title string) ([]sqlc.Book, error) {
	r.Logger.Info("getting books by title")
	books, err := r.DBTX.GetBooksByTitle(ctx, sqlc.GetBooksByTitleParams{
		Title: title,
		Limit: limit,
	})
	if err != nil {
		r.Logger.Error("error getting books by title ", err)
		return nil, err
	}
	r.Logger.Info("success getting books by title")
	return books, nil
}
