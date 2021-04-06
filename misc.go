package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LanguageCodesRespose struct {
	Code            string `json:"code"`
	CurrentLanguage string `json:"currentLanguage"`
}

func (c *Client) GetLanguageCodes(ctx context.Context) ([]*LanguageCodesRespose, int, error) {
	endpoint := fmt.Sprintf("%s/misc/lang", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	fmt.Println(string(bytes))
	var dataRes []*LanguageCodesRespose
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return dataRes, statusCode, nil
}

type InvoiceCheckoutPage struct {
	Page []byte
}

// View the checkout page of an invoice
func (c *Client) GetInvoiceCheckoutPage(invoiceID *InvoiceID, ctx context.Context) (*InvoiceCheckoutPage, int, error) {
	endpoint := fmt.Sprintf("%s/i/%s", c.URL, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceCheckoutPage
	func(bytes []byte, dataRes *InvoiceCheckoutPage) {
		dataRes.Page = bytes
	}(bytes, &dataRes)

	return &dataRes, statusCode, nil
}

func (i *Invoice) GetInvoiceCheckoutPage(ctx context.Context) (*InvoiceCheckoutPage, int, error) {
	endpoint := fmt.Sprintf("%s/i/%s", i.Client.URL, i.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceCheckoutPage
	func(bytes []byte, dataRes *InvoiceCheckoutPage) {
		dataRes.Page = bytes
	}(bytes, &dataRes)

	return &dataRes, statusCode, nil
}
