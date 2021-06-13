package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Storage struct {
	*Queries
	dbConn *sql.DB
}

func NewStorage(dbConn *sql.DB) *Storage {
	return &Storage{
		dbConn:  dbConn,
		Queries: New(dbConn),
	}
}

func (tm *Storage) execTx(ctx context.Context, fn func(q *Queries) (Transaction, error)) (Transaction, error) {
	tx, err := tm.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return Transaction{}, err
	}
	q := New(tx)
	transaction, err := fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return transaction, fmt.Errorf("tx err: %v, rb error %v", err, rbErr)
		}
		return Transaction{}, err
	}
	txErr := tx.Commit()
	if txErr != nil {
		return transaction, txErr
	}
	return transaction, nil
}

func (tm *Storage) InitiateTransfer(ctx context.Context, createTransactionParams CreateTransactionParams) (Transaction, error) {
	transaction, err := tm.execTx(ctx, func(q *Queries) (Transaction, error) {
		transaction := Transaction{}
		// Get lock for sender and reciever
		// fmt.Println(util.GetGoId(), "Getting lock", createTransactionParams.FromAccountID, createTransactionParams.ToAccountID)
		wallets, errR := q.lockWalletsForUpdate(ctx, []int64{createTransactionParams.FromAccountID, createTransactionParams.ToAccountID})
		if errR != nil {
			return transaction, errR
		}
		var sender Wallet
		var reciever Wallet
		foundSender := false
		foundReciever := false
		for _, w := range wallets {
			if w.ID == createTransactionParams.FromAccountID {
				sender = w
				foundSender = true
			}
		}
		for _, w := range wallets {
			if w.ID == createTransactionParams.ToAccountID {
				reciever = w
				foundReciever = true
			}
		}
		if !foundSender || !foundReciever {
			return transaction, errors.New("invalid account details")
		}

		// fmt.Println(util.GetGoId(), "sending", "[", createTransactionParams.FromAccountID, sender.Balance, "]", "[->", createTransactionParams.Amount, "->]", "[", createTransactionParams.ToAccountID, reciever.Balance, "]")
		if sender.Balance < createTransactionParams.Amount {
			return transaction, errors.New("insufficent balance")
		}

		if sender.Currency != reciever.Currency {
			return transaction, errors.New("incompatible currency")
		}

		// update balances
		err0 := q.UpdateWallet(ctx, UpdateWalletParams{ID: createTransactionParams.FromAccountID, Amount: -createTransactionParams.Amount})
		if err0 != nil {
			return transaction, err0
		}
		err1 := q.UpdateWallet(ctx, UpdateWalletParams{ID: createTransactionParams.ToAccountID, Amount: createTransactionParams.Amount})
		if err1 != nil {
			return transaction, err1
		}
		return q.CreateTransaction(ctx, createTransactionParams)
	})

	return transaction, err
}
