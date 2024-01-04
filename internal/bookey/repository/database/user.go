package database

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Repository) CreateUser(ctx *gin.Context, user *entity.User) (sqlc.User, error) {
	r.Logger.Info("creating new user")
	User, err := r.DBTX.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:       user.UserID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Balance:      1000,
	})
	if err != nil {
		r.Logger.Error("error creating user: ", err)
		return User, err
	}
	r.Logger.Info("user created successfully")
	return User, nil
}
func (r *Repository) GetUserByID(ctx *gin.Context, ID uuid.UUID, password string) (sqlc.User, error) {
	r.Logger.Info("getting user by id")
	user, err := r.DBTX.GetUser(ctx, sqlc.GetUserParams{
		UserID:       ID,
		PasswordHash: password,
	})
	if err != nil {
		r.Logger.Error("error getting user by id ", err)
		return sqlc.User{}, err
	}
	r.Logger.Info("success getting user by id")
	return user, nil
}
func (r *Repository) DeleteUserByID(ctx *gin.Context, ID uuid.UUID, password string) error {
	r.Logger.Info("deleting user by id")
	err := r.DBTX.DeleteUser(ctx, sqlc.DeleteUserParams{
		UserID:       ID,
		PasswordHash: password,
	})
	if err != nil {
		r.Logger.Error("error deleting user by id ", err)
		return err
	}
	r.Logger.Info("success deleting user by id")
	return nil
}
