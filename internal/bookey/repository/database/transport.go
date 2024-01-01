package database

import (
	"bookey/internal/bookey/entity"
	"bookey/internal/bookey/repository/migrations/sqlc"
	logging "bookey/pkg/logger"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Repository struct {
	Logger *logging.Logger
	DB     *sql.DB
	DBTX   *sqlc.Queries
}

type RepositoryTransport interface {
	CreateUser(ctx *gin.Context, user *entity.User) (sqlc.User, error)
	GetUserByID(ctx *gin.Context, ID uuid.UUID, password string) (sqlc.User, error)
	DeleteUserByID(ctx *gin.Context, ID uuid.UUID, password string) error
	CreateBook(ctx *gin.Context, book *entity.Book) (sqlc.Book, error)
	GetBooksByAuthor(ctx *gin.Context, author string) ([]sqlc.Book, error)
	GetBooksByGenre(ctx *gin.Context, genre string) ([]sqlc.Book, error)
	GetBooksByUserID(ctx *gin.Context, ID uuid.UUID) ([]sqlc.Book, error)
	GetBooksByTitle(ctx *gin.Context, title string) ([]sqlc.Book, error)
	TransactUserToUser(context context.Context, dto *entity.TransactionDTO) (sqlc.Transaction, error)
	FindUsersTransactions(ctx context.Context, userId uuid.UUID) ([]sqlc.Transaction, error)
}

func NewRepository(logger *logging.Logger, DB *sql.DB, DBTX *sqlc.Queries) RepositoryTransport {
	return &Repository{
		Logger: logger,
		DB:     DB,
		DBTX:   DBTX,
	}
}
