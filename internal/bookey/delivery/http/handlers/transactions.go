package handlers

import (
	"github.com/F1zm0n/bookey/internal/bookey/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TransactionUserToUser(ctx *gin.Context) {
	var Transaction *entity.TransactionDTO
	err := ctx.BindJSON(&Transaction)

	if err != nil {
		h.Logger.Error("error accepting user data: ", err)
		ctx.AbortWithError(400, err)
	}
	transaction, err := h.Usecase.TransactUserToUser(ctx, Transaction)
	if err != nil {
		h.Logger.Error("error transacting: ", err)
		ctx.AbortWithError(400, err)
	}

	ctx.JSON(200, transaction)
}
func (h *Handler) GetUsersTransactions(ctx *gin.Context) {
	var UserID *entity.UserID
	err := ctx.BindJSON(&UserID)
	if err != nil {
		h.Logger.Error("error accepting user data: ", err)
		ctx.AbortWithError(400, err)
	}
	transactions, err := h.Usecase.GetUsersTransactions(ctx, UserID.UserID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})
		return
	}

	ctx.JSON(200, transactions)
}
