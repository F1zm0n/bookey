package database

import (
	"bookey/internal/bookey/entity"
	"bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Repository) CreateUser(ctx *gin.Context, user *entity.User) error {
	_, err := r.DBTX.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:       uuid.New(),
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Balance:      "0",
	})
	if err != nil {
		r.Logger.Error("error creating user: ", err)
		return err
	}
	return nil
}
