package usecase

import (
	"database/sql"
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/database"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	logging "github.com/F1zm0n/bookey/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Usecase struct {
	Logger     *logging.Logger
	Repository database.RepositoryTransport
}

type UsecaseTransport interface {
	CreateUser(ctx *gin.Context, user *entity.UserDTO) (sqlc.User, error)
	GetUserByID(ctx *gin.Context, user *entity.UserDTO) (sqlc.User, error)
	DeleteUserByID(ctx *gin.Context, user *entity.UserDTO) error
	CreateBook(ctx *gin.Context, book *entity.BookDTO) (sqlc.Book, error)
	GetBooksByAuthor(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error)
	GetBooksByGenre(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error)
	GetBooksByUserID(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error)
	GetBooksByTitle(ctx *gin.Context, book *entity.BookDTO) ([]sqlc.Book, error)
	TransactUserToUser(ctx *gin.Context, transact *entity.TransactionDTO) (*entity.Transaction, error)
	GetUsersTransactions(ctx *gin.Context, UserID uuid.UUID) ([]*entity.Transaction, error)
}

func NewUsecase(logger *logging.Logger, DB *sql.DB, DBTX *sqlc.Queries) UsecaseTransport {
	return &Usecase{
		Logger:     logger,
		Repository: database.NewRepository(logger, DB, DBTX),
	}
}
