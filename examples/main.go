package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jon4hz/go-btcpay"
	"github.com/jon4hz/go-btcpay/examples/config"
)

func main() {
	ctx := context.Background()
	config, err := config.ReadConf("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	client := btcpay.NewBasicClient(config.BTCPay.URL, config.BTCPay.Username, config.BTCPay.Password)
	fmt.Println(client.GetHealth(ctx))

	fmt.Println(client.GetServerInfo(ctx))

	// create new APIKey
	apiKey, err := client.CreateAPIKey(&btcpay.CreateAPIKeyRequest{
		Permissions: []btcpay.Permission{client.CreateRestrictedKey(btcpay.BTCPayStoreCanviewinvoices, "66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn")},
	}, ctx)
	if err != nil {
		panic(err)
	}

	// delete the new APIKey
	err = client.RevokeAPIKey(apiKey.APIKey, ctx)
	if err != nil {
		panic(err)
	}

	// create new APIKey
	apiKey, err = client.CreateAPIKey(&btcpay.CreateAPIKeyRequest{
		Permissions: []btcpay.Permission{client.CreateRestrictedKey(btcpay.BTCPayStoreCanviewinvoices, "66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn")},
	}, ctx)
	if err != nil {
		panic(err)
	}

	// add APIKey to client
	client.APIKey = apiKey.APIKey
	fmt.Println(client.GetCurrentAPIKey(ctx))
	time.Sleep(10 * time.Second)
	client.RevokeCurrentAPIKey(ctx)

}
