package main

import (
	"fmt"
	"github.com/ttdung/astra-go-sdk/client"
	"github.com/ttdung/astra-go-sdk/config"
)

func main() {
	cfg := &config.Config{
		ChainId:       "chain-id",
		Endpoint:      "http://localhost:26657",
		CoinType:      60,
		PrefixAddress: "astra",
		TokenSymbol:   "aastra",
	}

	astraClient := client.NewClient(cfg)

	keyWallet := astraClient.NewAccountClient()
	acc, err := keyWallet.CreateAccount()
	if err != nil {
		panic(err)
	}

	data, _ := acc.String()
	fmt.Println(data)
}
