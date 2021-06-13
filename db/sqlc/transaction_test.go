package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestTransaction(t *testing.T, from Wallet, to Wallet) Transaction {
	transaction, err := testQueries.CreateTransaction(context.Background(), CreateTransactionParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        100,
	})
	require.NoError(t, err)
	return transaction
}

func TestCreateTransaction(t *testing.T) {
	w1 := createTestWallet(t)
	w2 := createTestWallet(t)
	t1 := createTestTransaction(t, w1, w2)
	require.Equal(t, t1.FromAccountID, w1.ID)
	require.Equal(t, t1.ToAccountID, w2.ID)
	require.NotZero(t, t1.CreatedAt)
}

func TestGetTransaction(t *testing.T) {
	w1 := createTestWallet(t)
	w2 := createTestWallet(t)
	t1 := createTestTransaction(t, w1, w2)
	t2, err := testQueries.GetTransaction(context.Background(), t1.ID)
	require.NoError(t, err)
	reflect.DeepEqual(t1, t2)
}
