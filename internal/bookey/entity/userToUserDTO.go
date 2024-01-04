package entity

import "github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"

func UserToUserDTO(user *sqlc.User) *UserDTO {
	NewUserDTO := &UserDTO{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.PasswordHash,
		Balance:  user.Balance,
	}
	return NewUserDTO
}
