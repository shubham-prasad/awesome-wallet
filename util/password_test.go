package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPasswordHashing(t *testing.T) {
	pass := RandomString(10)
	hash, err := CreatePasswordHash(pass)
	require.NoError(t, err)
	cmp := ComparePassword(hash, pass)
	require.NoError(t, cmp)
	cmp2 := ComparePassword(hash, RandomString(15))
	require.Error(t, cmp2)
}
