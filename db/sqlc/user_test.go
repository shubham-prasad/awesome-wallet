package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/shubham-prasad/awesome-wallet/util"
	"github.com/stretchr/testify/require"
)

func createMockUser(t *testing.T) User {
	pass, err := util.CreatePasswordHash(util.RandomString(16))
	require.NoError(t, err)
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Name: util.RandomString(10),
		Pwd:  pass,
	})
	require.NoError(t, err)
	return user
}

func TestCreateUser(t *testing.T) {
	createMockUser(t)
}

func TestGetUser(t *testing.T) {
	usr := createMockUser(t)
	getUsr, err := testQueries.GetUser(context.Background(), usr.ID)
	require.NoError(t, err)
	require.True(t, reflect.DeepEqual(usr, getUsr))
}

func TestListUser(t *testing.T) {
	for i := 0; i < 15; i++ {
		createMockUser(t)
	}
	users, err := testQueries.ListUser(context.Background(), ListUserParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(users), 10)
}

func TestUpdateUsername(t *testing.T) {
	usr := createMockUser(t)
	newName := util.RandomString(15)
	_, err := testQueries.UpdateUserName(context.Background(), UpdateUserNameParams{
		ID:   usr.ID,
		Name: newName,
	})
	require.NoError(t, err)
	getUsr, err := testQueries.GetUser(context.Background(), usr.ID)
	require.NoError(t, err)
	require.Equal(t, getUsr.Name, newName)
	require.Equal(t, getUsr.ID, usr.ID)
}

func TestUpdatePassword(t *testing.T) {
	usr := createMockUser(t)
	newPass, err := util.CreatePasswordHash(util.RandomString(15))
	require.NoError(t, err)
	_, err = testQueries.UpdateUserPassword(context.Background(), UpdateUserPasswordParams{
		ID:  usr.ID,
		Pwd: newPass,
	})
	require.NoError(t, err)
	getUsr, err := testQueries.GetUser(context.Background(), usr.ID)
	require.NoError(t, err)
	require.Equal(t, getUsr.Pwd, newPass)
	require.Equal(t, getUsr.ID, usr.ID)
}
