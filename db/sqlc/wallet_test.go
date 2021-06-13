package db

import (
	"context"
	"testing"

	"github.com/shubham-prasad/awesome-wallet/util"
	"github.com/stretchr/testify/require"
)

func createTestWallet(t *testing.T) Wallet {
	cust := createMockUser(t)
	args := CreateWalletParams{
		Owner:    cust.ID,
		Currency: CurrencyUSD,
		Balance:  int64(util.RandomInt(10, 1000)),
	}
	wallet, err := testQueries.CreateWallet(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, wallet)
	require.Equal(t, args.Owner, wallet.Owner)
	require.Equal(t, args.Balance, wallet.Balance)
	require.Equal(t, args.Currency, wallet.Currency)
	require.NotZero(t, wallet.ID)
	require.NotZero(t, wallet.CreatedAt)
	return wallet
}

func TestCreateWallet(t *testing.T) {
	createTestWallet(t)
}

func TestGetWallet(t *testing.T) {
	wallet1 := createTestWallet(t)
	wallet2, err := testQueries.GetWallet(context.Background(), wallet1.ID)
	require.NoError(t, err)
	require.Equal(t, wallet1.Balance, wallet2.Balance)
	require.Equal(t, wallet1.ID, wallet2.ID)
}

func TestListWallet(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestWallet(t)
	}
	wallets, err := testQueries.ListWallet(context.Background(), ListWalletParams{Limit: 5, Offset: 5})
	require.NoError(t, err)
	require.Equal(t, len(wallets), 5)
	for _, wallet := range wallets {
		require.NotEmpty(t, wallet)
	}
}

func TestUpdateWallet(t *testing.T) {
	w1 := createTestWallet(t)
	err := testQueries.UpdateWallet(context.Background(), UpdateWalletParams{ID: w1.ID, Amount: -500})
	require.NoError(t, err)
	w2, err := testQueries.GetWallet(context.Background(), w1.ID)
	require.NoError(t, err)
	require.EqualValues(t, w2.Balance-w1.Balance, int64(-500))
	require.Greater(t, w2.UpdatedAt.Sub(w1.UpdatedAt).Milliseconds(), (int64)(0))
	require.Equal(t, w2.CreatedAt.Sub(w1.CreatedAt).Milliseconds(), int64(0))
}

func TestDeleteWallet(t *testing.T) {
	wallets, err := testQueries.ListWallet(context.Background(), ListWalletParams{Limit: 500, Offset: 0})
	require.NoError(t, err)
	for _, wallet := range wallets {
		testQueries.DeleteWallet(context.Background(), wallet.ID)
	}
}
