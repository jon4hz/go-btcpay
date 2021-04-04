package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jon4hz/go-btcpay"
	"github.com/jon4hz/go-btcpay/examples/config"
)

func main() {
	config, err := config.ReadConf("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	client := btcpay.NewBasicClient(config.BTCPay.URL, config.BTCPay.Username, config.BTCPay.Password)
	ctx := context.Background()
	fmt.Println(client.CreateAPIKey(&btcpay.CreateAPIKeyRequest{}, ctx))
}
