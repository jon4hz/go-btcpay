package main

// WIP

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jon4hz/go-btcpay"
	"github.com/jon4hz/go-btcpay/examples/config"
)

var ctx = context.Background()

func main() {
	config, err := config.ReadConf("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	client := btcpay.NewBasicClient(config.BTCPay.URL, config.BTCPay.Username, config.BTCPay.Password)

	/* cont, cancel := context.WithTimeout(ctx, 1)
	defer cancel()
	fmt.Println(client.GetHealth(cont)) */
	//fmt.Println(client.GetServerInfo(ctx))

	// get store id
	stores, _, err := client.GetStores(ctx)
	if err != nil {
		panic(err)
	}
	storeID := getStoreID(stores)
	createInvoiceByStoreGetAndDeleteInvoiceByID(client, storeID)

	//getInvoicesByStore(client, storeID)
	//reateAndDeleteInvoice(client, storeID)

	//createNewStore(client)

	//createNewUser(client)
	//fmt.Println(client.GetUser(ctx))

	/* fmt.Println((&btcpay.Client{URL: client.URL}).GetLanguageCodes(ctx))

	langs, _, err := client.GetLanguageCodes(ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range langs {
		fmt.Println(v)
	} */

	//createPrintPageDeleteInvoice(client, &storeID)
	testStruct()
	getPaymentRequests(client, &storeID)
}

func getPaymentRequests(c *btcpay.Client, storeID *btcpay.StoreID) {
	paymentRequests, _, err := c.GetPaymentRequests(ctx, storeID)
	if err != nil {
		panic(err)
	}
	for _, v := range paymentRequests {
		fmt.Println(v)
	}

}

func createPrintPageDeleteInvoice(c *btcpay.Client, storeID *btcpay.StoreID) {
	invoice, _, _ := c.CreateInvoice(ctx, storeID, &btcpay.InvoiceRequest{
		Amount:   "10",
		Currency: "USD",
	})
	page, _, _ := c.GetInvoiceCheckoutPage(ctx, &invoice.ID)
	fmt.Println(string(page.Page))
	invoiceC := c.Invoice
	invoiceC.ID = invoice.ID
	fmt.Println(invoiceC.ArchiveInvoice(ctx))
}

func createNewUser(c *btcpay.Client) {
	fmt.Println(c.CreateUser(ctx, &btcpay.UserRequest{
		Email:           "test@test.com",
		Password:        "asdfasdf",
		IsAdministrator: false,
	}))
}

func createNewStore(c *btcpay.Client) {
	fmt.Println(c.CreateStore(ctx, &btcpay.StoreRequest{Name: "test03"}))
}

func createInvoiceByStoreGetAndDeleteInvoiceByID(client *btcpay.Client, storeID btcpay.StoreID) {
	// create a new storeClient
	storeClient := client.Store
	// assign a store ID to the storeClient
	storeClient.ID = storeID

	// create a new invoice for the store
	invoice, _, err := storeClient.CreateInvoice(ctx, &btcpay.InvoiceRequest{Amount: "10", Currency: "USD"})
	if err != nil {
		panic(err)
	}
	// create a new invoiceClient, based on the current client
	invoiceClient := client.Invoice
	// assign a storeClient to the invoiceClient
	invoiceClient.Store = storeClient
	// assign a invoice ID to the invoiceClient
	invoiceClient.ID = invoice.ID

	fmt.Println(invoiceClient.GetInvoice(ctx))
	invoiceClient.ArchiveInvoice(ctx)
}

func getInvoicesByStore(client *btcpay.Client, storeID btcpay.StoreID) {
	storeClient := client.Store
	storeClient.ID = storeID
	fmt.Println(storeClient.GetInvoices(ctx))
}

func getStoreID(stores []*btcpay.StoreResponse) btcpay.StoreID {
	for _, v := range stores {
		if v.Name == "test01" {
			return v.ID
		}
	}
	return ""
}

func createAndDeleteInvoice(client *btcpay.Client, storeID btcpay.StoreID) {
	fmt.Println(storeID)
	invoice, _, err := client.CreateInvoice(ctx, &storeID, &btcpay.InvoiceRequest{Amount: "11", Currency: "USD", Metadata: btcpay.InvoiceMetadata{"test": "asdf", "test2": "aaaa"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(invoice)
	client.UpdateInvoice(ctx, &storeID, &invoice.ID, &btcpay.InvoiceUpdate{Metadata: btcpay.InvoiceMetadata{"test3": "ccccc"}})
	fmt.Println(client.GetInvoices(ctx, &storeID))
	client.ArchiveInvoice(ctx, &storeID, &invoice.ID)
}

func createAndDeleteAPIKey(client *btcpay.Client) {
	// create new APIKey
	apiKey, _, err := client.CreateAPIKey(ctx, &btcpay.APIKeyRequest{
		Permissions: []btcpay.BTCPayPermission{btcpay.CreateCustomPermission(btcpay.GetPermission().StoreCanviewinvoices, btcpay.StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))}})
	if err != nil {
		panic(err)
	}

	// delete the new APIKey
	_, err = client.RevokeAPIKey(ctx, &apiKey.APIKey)
	if err != nil {
		panic(err)
	}
}

func createAndDeleteCurrentAPIKey(client *btcpay.Client) {
	// create new APIKey
	apiKey, _, err := client.CreateAPIKey(ctx, &btcpay.APIKeyRequest{
		Permissions: []btcpay.BTCPayPermission{btcpay.CreateCustomPermission(btcpay.GetPermission().StoreCanviewinvoices, btcpay.StoreID("66tU3WhCAcsbocA3EmUXHE96XsoVQjWMUiTp3s6LLYAn"))}})
	if err != nil {
		panic(err)
	}

	// add APIKey to client
	client.APIKey = apiKey.APIKey
	fmt.Println(client.GetCurrentAPIKey(ctx))
	time.Sleep(10 * time.Second)
	client.RevokeCurrentAPIKey(ctx)
}

func testStruct() {
	type testStruct2 struct {
		Data  int  `json:"data"`
		Valid bool `json:"valid"`
	}

	type testStruct struct {
		X []testStruct2
	}

	data := `{"x1": {"data": 1, "valid": true}, "x2":{"data": 2, "valid": false}}`

	var test testStruct
	err := json.Unmarshal([]byte(data), &test)
	if err != nil {
		panic(err)
	}
	fmt.Println(test.X)
}
