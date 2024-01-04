package handlers

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateBook(ctx *gin.Context) {
	var BookDTO *entity.BookDTO
	err := ctx.BindJSON(&BookDTO)

	if err != nil {
		h.Logger.Error("error accepting user data: ", err)
		ctx.AbortWithError(400, err)
	}

	h.Logger.Info("creating book")

	book, err := h.Usecase.CreateBook(ctx, BookDTO)
	if err != nil {
		ctx.AbortWithError(400, err)
	}

	NewBookDTO := entity.BookToBookDTO(&book)

	ctx.JSON(200, NewBookDTO)
}
func (h *Handler) SearchBook(ctx *gin.Context) {
	var BookDTO *entity.BookDTO

	err := ctx.BindJSON(&BookDTO)

	if err != nil {
		h.Logger.Error("error getting user data ", err)
		ctx.AbortWithError(400, err)
	}
	var NewBooks *[]sqlc.Book

	if BookDTO.Genre != "" {
		NewBooks = h.searchTempl(ctx, BookDTO, h.Usecase.GetBooksByGenre)

	} else if BookDTO.Author != "" {
		NewBooks = h.searchTempl(ctx, BookDTO, h.Usecase.GetBooksByAuthor)

	} else if BookDTO.UserID != uuid.Nil {
		NewBooks = h.searchTempl(ctx, BookDTO, h.Usecase.GetBooksByUserID)

	} else if BookDTO.Title != "" {
		NewBooks = h.searchTempl(ctx, BookDTO, h.Usecase.GetBooksByTitle)

	}
	var NewBooksDTO []*entity.BookDTO

	var NewBookDTO *entity.BookDTO
	for _, book := range *NewBooks {
		NewBookDTO = entity.BookToBookDTO(&book)
		NewBooksDTO = append(NewBooksDTO, NewBookDTO)
	}
	ctx.JSON(200, NewBooksDTO)
}
