package handlers

import (
	"database/sql"
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/F1zm0n/bookey/internal/bookey/usecase"
	logging "github.com/F1zm0n/bookey/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase usecase.UsecaseTransport
	Logger  *logging.Logger
}

type HandlerTransport interface {
	Register(engine *gin.Engine)
}

func NewHandler(logger *logging.Logger, DB *sql.DB, DBTX *sqlc.Queries) HandlerTransport {
	return &Handler{
		Usecase: usecase.NewUsecase(logger, DB, DBTX),
		Logger:  logger,
	}
}

type searchFunc func(ctx *gin.Context, Book *entity.BookDTO) ([]sqlc.Book, error)

func (h *Handler) searchTempl(ctx *gin.Context, BookDTO *entity.BookDTO, ourFunc searchFunc) *[]sqlc.Book {

	NewBooksDTO, err := ourFunc(ctx, BookDTO)
	if err != nil {
		h.Logger.Error("error getting books data: ", err)
		ctx.AbortWithError(400, err)
	}
	return &NewBooksDTO
}
