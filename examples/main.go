package main

// WIP

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

	// get store id
	stores, _, err := client.GetStores(ctx)
	if err != nil {
		panic(err)
	}
	storeID := getStoreID(stores)
	if len(storeID) > 0 {
		fmt.Println(storeID)
		invoice, _, err := client.CreateInvoice(&storeID, &btcpay.CreateInvoiceRequest{Amount: "11", Currency: "USD", Metadata: btcpay.InvoiceMetadata{"test": "asdf", "test2": "aaaa"}}, ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(invoice)
		client.UpdateInvoice(&storeID, &invoice.ID, &btcpay.InvoiceUpdate{Metadata: btcpay.InvoiceMetadata{"test3": "ccccc"}}, ctx)
		fmt.Println(client.GetInvoices(&storeID, ctx))
		client.ArchiveInvoice(&storeID, &invoice.ID, ctx)
	}

}

func getStoreID(stores []*btcpay.StoreResponse) btcpay.StoreID {
	for _, v := range stores {
		if v.Name == "test01" {
			return v.ID
		}
	}
	return ""
}

func createAndDeleteAPIKey(client *btcpay.Client) {

	// create new APIKey
	apiKey, _, err := client.CreateAPIKey(&btcpay.CreateAPIKeyRequest{
		Permissions: []btcpay.Permission{client.CreateRestrictedKey(btcpay.BTCPayStoreCanviewinvoices, btcpay.StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))},
	}, context.Background())
	if err != nil {
		panic(err)
	}

	// delete the new APIKey
	_, err = client.RevokeAPIKey(&apiKey.APIKey, context.Background())
	if err != nil {
		panic(err)
	}
}

func createAndDeleteCurrentAPIKey(client *btcpay.Client) {
	// create new APIKey
	apiKey, _, err := client.CreateAPIKey(&btcpay.CreateAPIKeyRequest{
		Permissions: []btcpay.Permission{client.CreateRestrictedKey(btcpay.BTCPayStoreCanviewinvoices, btcpay.StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))},
	}, context.Background())
	if err != nil {
		panic(err)
	}

	// add APIKey to client
	client.APIKey = apiKey.APIKey
	fmt.Println(client.GetCurrentAPIKey(context.Background()))
	time.Sleep(10 * time.Second)
	client.RevokeCurrentAPIKey(context.Background())
}
