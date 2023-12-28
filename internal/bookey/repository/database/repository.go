package database

import (
	logging "bookey/pkg/logger"
	"database/sql"
)

type UserRepository struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewUserRepository(db *sql.DB, logger *logging.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}
