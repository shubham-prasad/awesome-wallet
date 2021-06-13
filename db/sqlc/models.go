// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"fmt"
	"time"
)

type Currency string

const (
	CurrencyUSD      Currency = "USD"
	CurrencyINR      Currency = "INR"
	CurrencyCASHBACK Currency = "CASHBACK"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported scan type for Currency: %T", src)
	}
	return nil
}

type Transaction struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"fromAccountID"`
	ToAccountID   int64     `json:"toAccountID"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Pwd       string    `json:"pwd"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Wallet struct {
	ID        int64     `json:"id"`
	Owner     int64     `json:"owner"`
	Currency  Currency  `json:"currency"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}