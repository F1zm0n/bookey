package usecase

import (
	"bookey/internal/bookey/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

func (u *Usecase) TransactUserToUser(ctx *gin.Context, transact *entity.TransactionDTO) (*entity.Transaction, error) {
	u.Logger.Info("releasing transaction")

	transaction, err := u.Repository.TransactUserToUser(ctx, transact)
	if err != nil {
		u.Logger.Error("error transacting")
		return &entity.Transaction{}, fmt.Errorf("error transacting")
	}

	amount, err := strconv.ParseFloat(transaction.Amount, 64)

	newTransaction := &entity.Transaction{
		TransactionID: transaction.TransactionID,
		ToUserID:      transaction.ToUserID,
		FromUserID:    transaction.FromUserID,
		Sum:           amount,
		Description:   transaction.Description,
		Date:          transaction.TransactionDate,
	}

	return newTransaction, nil
}

func (u *Usecase) GetUsersTransactions(ctx *gin.Context, UserID uuid.UUID) ([]*entity.Transaction, error) {

	transactions, err := u.Repository.FindUsersTransactions(ctx, UserID)
	if err != nil {
		return nil, err
	}
	var TransactionEntity *entity.Transaction
	var Transactions []*entity.Transaction
	for _, transaction := range transactions {
		floatTransaction, _ := strconv.ParseFloat(transaction.Amount, 64)
		TransactionEntity = &entity.Transaction{
			TransactionID: transaction.TransactionID,
			FromUserID:    transaction.FromUserID,
			ToUserID:      transaction.ToUserID,
			Sum:           floatTransaction,
			Description:   transaction.Description,
			Date:          transaction.TransactionDate,
		}
		Transactions = append(Transactions, TransactionEntity)
	}
	return Transactions, err
}
