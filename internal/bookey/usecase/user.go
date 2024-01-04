package usecase

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *Usecase) CreateUser(ctx *gin.Context, user *entity.UserDTO) (sqlc.User, error) {
	userForFunc := &entity.User{
		UserID:       uuid.New(),
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
		Balance:      0,
	}
	User, err := u.Repository.CreateUser(ctx, userForFunc)
	if err != nil {
		return sqlc.User{}, err
	}

	return User, nil
}

func (u *Usecase) GetUserByID(ctx *gin.Context, user *entity.UserDTO) (sqlc.User, error) {
	User, err := u.Repository.GetUserByID(ctx, user.UserID, user.Password)
	if err != nil {
		return sqlc.User{}, err
	}
	return User, nil
}

func (u *Usecase) DeleteUserByID(ctx *gin.Context, user *entity.UserDTO) error {
	err := u.Repository.DeleteUserByID(ctx, user.UserID, user.Password)
	if err != nil {
		return err
	}
	return nil
}
