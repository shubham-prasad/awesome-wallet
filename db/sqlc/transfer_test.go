package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	custA := createMockUser(t)
	userA, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custA.ID,
		Balance:  1000,
		Currency: "USD",
	})
	require.NoError(t, err)
	custB := createMockUser(t)
	userB, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		Owner:    custB.ID,
		Balance:  100,
		Currency: "USD",
	})
	require.NoError(t, err)
	_, err0 := bankDb.InitiateTransfer(context.Background(), CreateTransactionParams{
		FromAccountID: userA.ID,
		ToAccountID:   userB.ID,
		Amount:        100,
	})
	require.NoError(t, err0)
	userA2, getErr := testQueries.GetWallet(context.Background(), userA.ID)
	require.NoError(t, getErr)
	require.Equal(t, userA2.Balance, int64(900))
	userB2, getErr := testQueries.GetWallet(context.Background(), userB.ID)
	require.NoError(t, getErr)
	require.Equal(t, userB2.Balance, int64(200))

	_, err2 := bankDb.InitiateTransfer(context.Background(), CreateTransactionParams{
		FromAccountID: userB.ID,
		ToAccountID:   userA.ID,
		Amount:        500,
	})
	require.Equal(t, err2.Error(), "insufficent balance")

}
