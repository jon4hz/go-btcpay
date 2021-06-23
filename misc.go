package btcpay

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetLanguageCodes(ctx context.Context) ([]*LanguageCodesRespose, int, error) {
	endpoint := fmt.Sprintf("%s/misc/lang", c.URL)
	var dataRes []*LanguageCodesRespose
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// View the checkout page of an invoice
func (c *Client) GetInvoiceCheckoutPage(ctx context.Context, invoiceID *InvoiceID, lang ...string) (*InvoiceCheckoutPage, int, error) {
	var endpoint string
	if len(lang) > 0 {
		fmt.Println(lang[0])
		endpoint = fmt.Sprintf("%s/i/%s?lang=%s", c.URL, *invoiceID, lang[0])
	} else {
		endpoint = fmt.Sprintf("%s/i/%s", c.URL, *invoiceID)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doSimpleRequest(req)
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
	bytes, statusCode, err := i.Client.doSimpleRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceCheckoutPage
	func(bytes []byte, dataRes *InvoiceCheckoutPage) {
		dataRes.Page = bytes
	}(bytes, &dataRes)
	return &dataRes, statusCode, nil
}
