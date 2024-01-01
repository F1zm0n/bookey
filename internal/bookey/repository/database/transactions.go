package database

import (
	"bookey/internal/bookey/entity"
	"bookey/internal/bookey/repository/migrations/sqlc"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func (r *Repository) TransactUserToUser(context context.Context, dto *entity.TransactionDTO) (sqlc.Transaction, error) {
	tx, err := r.DB.BeginTx(context, nil)
	if err != nil {
		return sqlc.Transaction{}, err
	}
	defer tx.Rollback()

	var fromUserBalance int32

	fmt.Println(dto.FromUserPassword)
	if err := tx.QueryRowContext(context, "SELECT balance FROM Users WHERE user_id=$1",
		dto.FromUserID).
		Scan(&fromUserBalance); err != nil {
		if err == sql.ErrNoRows {
			return sqlc.Transaction{}, fmt.Errorf("no rows returned  %v", err)
		}

		return sqlc.Transaction{}, fmt.Errorf("no such user %v", err)
	}
	var BookSum int32
	var ToUserID uuid.UUID
	fmt.Println(dto.BookID)
	if err := tx.QueryRowContext(context, `SELECT price,user_id FROM Books WHERE book_id=$1`,
		dto.BookID).
		Scan(&BookSum, &ToUserID); err != nil {
		if err == sql.ErrNoRows {
			return sqlc.Transaction{}, fmt.Errorf("wrong book id %v", err)
		}
	}

	if BookSum > fromUserBalance {
		return sqlc.Transaction{}, fmt.Errorf("not enough balance")
	}

	_, err = tx.ExecContext(context, "UPDATE Users SET balance= balance - $1 WHERE user_id = $2",
		BookSum, dto.FromUserID)
	if err != nil {
		return sqlc.Transaction{}, err
	}

	_, err = tx.ExecContext(context, "UPDATE Users SET balance= balance + $1 WHERE user_id = $2",
		BookSum, ToUserID)
	if err != nil {
		return sqlc.Transaction{}, err
	}

	err = r.DBTX.DeleteBookByID(context, dto.BookID)

	if err != nil {
		return sqlc.Transaction{}, fmt.Errorf("error deleting book %v", err)
	}

	stringBook := strconv.FormatInt(int64(BookSum), 10)

	transaction, err := r.DBTX.CreateTransaction(context, sqlc.CreateTransactionParams{
		TransactionID:   uuid.New(),
		FromUserID:      dto.FromUserID,
		ToUserID:        ToUserID,
		Amount:          stringBook,
		Description:     dto.Description,
		TransactionDate: time.Now(),
	})

	if err != nil {
		return sqlc.Transaction{}, fmt.Errorf("couldn't create transaction %v", err)
	}

	if err = tx.Commit(); err != nil {
		return sqlc.Transaction{}, fmt.Errorf("error commiting transaction")
	}

	return transaction, nil

}

func (r *Repository) FindUsersTransactions(ctx context.Context, userId uuid.UUID) ([]sqlc.Transaction, error) {
	r.Logger.Info("getting users transactions")

	transactions, err := r.DBTX.GetUsersTransactions(ctx, userId)
	if err != nil {
		r.Logger.Error("error getting users transactions")
		return []sqlc.Transaction{}, fmt.Errorf("error getting users transactions %v", err)
	}
	return transactions, nil

}
