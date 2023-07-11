package storage_handler_test

import (
	"testing"

	"github.com/JackalLabs/jackalgo/handlers/storage_handler"
	"github.com/JackalLabs/jackalgo/handlers/wallet_handler"
	"github.com/stretchr/testify/require"
)

func TestBuyStorage(t *testing.T) {
	r := require.New(t)

	wallet, err := wallet_handler.DefaultWalletHandler("slim odor fiscal swallow piece tide naive river inform shell dune crunch canyon ten time universe orchard roast horn ritual siren cactus upon forum")
	r.NoError(err)

	storageHandler := storage_handler.NewStorageHandler(wallet)

	_, err = storageHandler.BuyStorage("jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgc3sexc", 1, 1_000_000)
	r.NoError(err)
}
