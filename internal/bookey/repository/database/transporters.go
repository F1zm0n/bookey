package database

import (
	"bookey/internal/bookey/entity"
	"bookey/internal/bookey/repository/migrations/sqlc"
	logging "bookey/pkg/logger"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	Logger *logging.Logger
	DB     *sql.DB
	DBTX   *sqlc.Queries
}

type RepositoryTransport interface {
	CreateUser(ctx *gin.Context, user *entity.User) error
}

func NewRepository(logger *logging.Logger, DB *sql.DB) RepositoryTransport {
	return &Repository{
		Logger: logger,
		DB:     DB,
		DBTX:   sqlc.New(DB),
	}
}
