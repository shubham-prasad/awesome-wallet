package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParallelTransaction(t *testing.T) {
	n := 100
	custA := createMockUser(t)
	custB := createMockUser(t)
	w1, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custA.ID,
		Balance:  int64(1000),
		Currency: "USD",
	})
	require.NoError(t, err)
	w2, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custB.ID,
		Balance:  int64(1000),
		Currency: "USD",
	})
	require.NoError(t, err)

	pipe := make(chan error)

	for i := 0; i < n; i++ {
		go (func() {
			_, err := bankDb.InitiateTransfer(context.Background(), CreateTransactionParams{
				FromAccountID: w1.ID,
				ToAccountID:   w2.ID,
				Amount:        int64(10),
			})
			pipe <- err
		})()
	}

	for i := 0; i < n; i++ {
		err := <-pipe
		require.NoError(t, err)
	}

	// reconsile
	w1After, err := testQueries.GetWallet(context.Background(), w1.ID)
	require.NoError(t, err)
	w2After, err := testQueries.GetWallet(context.Background(), w2.ID)
	require.NoError(t, err)
	require.Equal(t, int64(0), w1After.Balance)
	require.Equal(t, int64(2000), w2After.Balance)
}

func TestBackForthParallelTransaction(t *testing.T) {
	n := 10
	custA := createMockUser(t)
	custB := createMockUser(t)
	w1, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custA.ID,
		Balance:  int64(100),
		Currency: "USD",
	})
	require.NoError(t, err)
	w2, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custB.ID,
		Balance:  int64(100),
		Currency: "USD",
	})
	require.NoError(t, err)

	pipe := make(chan error)

	for i := 0; i < n; i++ {
		go (func(i int) {
			var from int64
			var to int64
			if i%2 == 0 {
				from = w1.ID
				to = w2.ID
			} else {
				from = w2.ID
				to = w1.ID
			}
			_, err := bankDb.InitiateTransfer(context.Background(), CreateTransactionParams{
				FromAccountID: from,
				ToAccountID:   to,
				Amount:        int64(10),
			})
			pipe <- err
		})(i)
	}

	for i := 0; i < n; i++ {
		err := <-pipe
		require.NoError(t, err)
	}

	// reconsile
	w1After, err := testQueries.GetWallet(context.Background(), w1.ID)
	require.NoError(t, err)
	w2After, err := testQueries.GetWallet(context.Background(), w2.ID)
	require.NoError(t, err)
	require.Equal(t, int64(100), w1After.Balance)
	require.Equal(t, int64(100), w2After.Balance)
}
