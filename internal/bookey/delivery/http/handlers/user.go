package handlers

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var UserDTO *entity.UserDTO
	err := ctx.BindJSON(&UserDTO)

	if err != nil {
		h.Logger.Error("error accepting user data: ", err)
		ctx.AbortWithError(400, err)
	}
	user, err := h.Usecase.CreateUser(ctx, UserDTO)
	if err != nil {
		h.Logger.Error("error inserting user in database: ", err)
		ctx.AbortWithError(400, err)
	}
	NewUserDTO := entity.UserToUserDTO(&user)
	ctx.JSON(200, NewUserDTO)
}
func (h *Handler) GetUser(ctx *gin.Context) {
	var UserDTO *entity.UserDTO
	err := ctx.BindJSON(&UserDTO)

	if err != nil {
		h.Logger.Error("error accepting user data: ", err)
		ctx.AbortWithError(400, err)
	}

	h.Logger.Info("getting user by id")
	user, err := h.Usecase.GetUserByID(ctx, UserDTO)
	if err != nil {
		h.Logger.Error("error getting user from request")
		ctx.AbortWithError(400, err)
	}

	NewUserDTO := entity.UserToUserDTO(&user)
	ctx.JSON(200, NewUserDTO)
}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	var UserDTO *entity.UserDTO
	ctx.BindJSON(&UserDTO)

	h.Logger.Info("deleting user")
	err := h.Usecase.DeleteUserByID(ctx, UserDTO)
	if err != nil {
	}
}
