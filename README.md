# Go-BTCPay
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/jon4hz/go-btcpay/blob/master/LICENSE) 
[![testing](https://github.com/jon4hz/go-btcpay/actions/workflows/testing.yml/badge.svg)](https://github.com/jon4hz/go-btcpay/actions/workflows/testing.yml)

_-- Work in progress --_

A Golang SDK for the BTCPay Server [Greenfield API v1](https://docs.btcpayserver.org/API/Greenfield/v1/).

## 💡 About
This package provies full access to the Greenfield API v1 from a BTCPayServer. Every API call returns, if available, a corresponding go struct, a HTTP status code and an error.

It's possible to control the individual calls by passing a context for each function and method.

## 🚀 Getting started

### 🧑‍💻 Create a client

You can create a client either by using basic authentication or by using an API Key.

```go
package main

import (
    "context"
    "fmt"
    "github.com/jon4hz/go-btcpay"
)

func main() {
    // create empty context interface
    ctx := context.Background()

    // Create a basicAuth client
    client := btcpay.CreateBasicClient("https://mybtcpayserver.com", "myUsername", "myPassword")

    // Print informations about the server, etc
    fmt.Println(client.GetServerInfo(ctx))

    // Does the same but with an APIKey instead of basicAuth
    // Create a client with an APIKey
    client2 := btcpay.CreateBasicClient("https://mybtcpayserver.com", btcpay.APIKey("myAPIKey")

    // Print informations about the server, etc again but use the APIKey based client
    fmt.Println(client2.GetServerInfo(ctx))
}
```

### 📝 Create an invoice
You can create an invoice by using the previously created client.
```go
// assign a store to the client
client.Store.ID = btcpay.StoreID("YourStoreID")

// create the invoice
invoice, _, err := client.CreateInvoice(context.TODO(), &client.Store.ID, &btcpay.InvoiceRequest{
    Amount:   "10",
    Currency: "USD",
})
if err != nil {
   fmt.Println(err)
} else {
    fmt.Println(invoice) // invoice has type *btcpay.InvoiceResponse
}
```


Calling the method `CreateInvoice()` works for variable of type *btcpay.Store, too.
```go
// by passing the store from the previously created client, the new store (*btcpay.Store) contains 
// a pointer  back to the initial client 
store = client.Store
// assign a storeID to the store
store.ID = btcpay.StoreID("YourStoreID")

// create the invoice
invoice, _, err := store.CreateInvoice(context.TODO(), &btcpay.InvoiceRequest{
    Amount:   "10",
    Currency: "USD",
})
if err != nil {
   fmt.Println(err)
} else {
    fmt.Println(invoice) // invoice has type *btcpay.InvoiceResponse
}
```


_[more examples will follow soon]_

initialented endpoints.

Endpoint                                                                   |              Status
|:-------------------------------------------------------------------------|:-------------------|
|`/api/v1/api-keys`                                                        | ✅ Fully implemented
|`/api-keys/authorize`                                                     | ⚡️ Testing required 
|`/api/v1/health`                                                          | ✅ Fully implemented
|`/api/v1/server/info`                                                     | ✅ Fully implemented
|`/api/v1/users`                                                           | ✅ Fully implemented
|`/api/v1/users/me/notifications`                                          | ✅ Fully implemented
|`/api/v1/stores`                                                          | ⚠️ Partially implemented
|`/api/v1/stores/{storeId}/invoices`                                       | ✅ Fully implemented
|`/api/v1/stores/{storeId}/payment-requests                              ` | ✅ Fully implemented
|`/api/v1/stores/{storeId}/pull-payments`                                  | ✅ Fully implemented
|`/api/v1/stores/{storeId}/payment-methods/OnChain/{cryptoCode}/wallet`    | ⏳ Work in progress
|`/misc/lang`                                                              | ❌ Not working, [issue](https://github.com/btcpayserver/btcpayserver/issues/2437)
|`/i`                                                                      | ✅ Fully implemented
|`/api/v1/pull-payments`                                                   | ✅ Fully implemented

## 📜 Licensing
This SDK is released under the MIT-License found in the [LICENSE](https://github.com/jon4hz/go-btcpay/blob/master/LICENSE) file.