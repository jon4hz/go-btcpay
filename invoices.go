package btcpay

// WIP

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvoicesResponse struct {
	Amount           string          `json:"amount,omitempty"`
	Currency         string          `json:"currency,omitempty"`
	Metadata         InvoiceMetadata `json:"metadata,omitempty"`
	Checkout         InvoiceCheckout `json:"checkout,omitempty"`
	ID               string          `json:"id"`
	CheckoutLink     string          `json:"checkoutLink"`
	CreatedTime      int64           `json:"createdTime"`
	ExpirationTime   int64           `json:"expirationTime"`
	MonitoringTime   int64           `json:"monitoringTime"`
	Status           string          `json:"status"`
	AdditionalStatus string          `json:"additionalStatus"`
}

type InvoiceMetadata struct {
	OrderID string `json:"orderId,omitempty"`
}

type SpeedPolicy string

const (
	HighSpeed      SpeedPolicy = "HighSpeed"
	MediumSpeed    SpeedPolicy = "MediumSpeed"
	LowMediumSpeed SpeedPolicy = "LowMediumSpeed"
	LowSpeed       SpeedPolicy = "LowSpeed"
)

type InvoiceCheckout struct {
	SpeedPolicy       SpeedPolicy `json:"speedPolicy,omitempty"`
	PaymentMethods    []string    `json:"paymentMethods,omitempty"`
	ExpirationMinutes int         `json:"expirationMinutes,omitempty"`
	MonitoringMinutes int         `json:"monitoringMinutes,omitempty"`
	PaymentTolerance  float64     `json:"paymentTolerance,omitempty"`
	RedirectURL       string      `json:"redirectURL,omitempty"`
	DefaultLanguage   string      `json:"defaultLanguage,omitempty"`
}

type GetInvoicesRequest struct {
	StoreID string `json:"storeId"`
}

func (c *Client) GetInvoices(invoiceReq *GetInvoicesRequest, ctx context.Context) (*InvoicesResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, invoiceReq)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data InvoicesResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
