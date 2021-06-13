// Code generated by sqlc. DO NOT EDIT.
// source: transaction.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  from_account_id, to_account_id, amount 
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransactionParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.queryRow(ctx, q.createTransactionStmt, createTransaction, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTransactionStmt, deleteTransaction, id)
	return err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.queryRow(ctx, q.getTransactionStmt, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransaction = `-- name: ListTransaction :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transactions
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTransactionParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransaction(ctx context.Context, arg ListTransactionParams) ([]Transaction, error) {
	rows, err := q.query(ctx, q.listTransactionStmt, listTransaction, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
