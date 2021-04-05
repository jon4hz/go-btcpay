package btcpay

// WIP

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvoiceID string

type InvoiceResponse struct {
	Amount           string                  `json:"amount,omitempty"`
	Currency         string                  `json:"currency,omitempty"`
	Metadata         InvoiceMetadata         `json:"metadata,omitempty"`
	Checkout         InvoiceCheckout         `json:"checkout,omitempty"`
	ID               InvoiceID               `json:"id"`
	CheckoutLink     string                  `json:"checkoutLink"`
	CreatedTime      int64                   `json:"createdTime"`
	ExpirationTime   int64                   `json:"expirationTime"`
	MonitoringTime   int64                   `json:"monitoringTime"`
	Status           InvoiceStatus           `json:"status"`
	AdditionalStatus InvoiceAdditionalStatus `json:"additionalStatus"`
}

type InvoiceMetadata map[string]string

type SpeedPolicy string

const (
	HighSpeed      SpeedPolicy = "HighSpeed"
	MediumSpeed    SpeedPolicy = "MediumSpeed"
	LowMediumSpeed SpeedPolicy = "LowMediumSpeed"
	LowSpeed       SpeedPolicy = "LowSpeed"
)

type InvoiceStatus string

const (
	New        InvoiceStatus = "New"
	Processing InvoiceStatus = "Processing"
	Expired    InvoiceStatus = "Expired"
	Invalid    InvoiceStatus = "Invalid"
	Settled    InvoiceStatus = "Settled"
)

type InvoiceAdditionalStatus string

const (
	None        InvoiceAdditionalStatus = "None"
	PaidLate    InvoiceAdditionalStatus = "PaidLate"
	PaidPartial InvoiceAdditionalStatus = "PaidPartial"
	Marked      InvoiceAdditionalStatus = "Marked"
	AddInvalid  InvoiceAdditionalStatus = "Invalid"
	PaidOver    InvoiceAdditionalStatus = "PaidOver"
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

func (c *Client) GetInvoices(storeID *StoreID, ctx context.Context) (*[]InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, *storeID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return &dataRes, statusCode, nil
}

type CreateInvoiceRequest struct {
	Amount          string          `json:"amount"`
	Currency        string          `json:"currency,omitempty"`
	Metadata        InvoiceMetadata `json:"metadata,omitempty"`
	InvoiceCheckout InvoiceCheckout `json:"checkout,omitempty"`
}

func (c *Client) CreateInvoice(storeID *StoreID, invoiceRequest *CreateInvoiceRequest, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, *storeID)
	dataReq, err := json.Marshal(invoiceRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return &dataRes, statusCode, nil
}

func (c *Client) GetInvoice(storeID *StoreID, invoiceID *InvoiceID, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return &dataRes, statusCode, nil
}

func (c *Client) ArchiveInvoice(storeID *StoreID, invoiceID *InvoiceID, ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := c.doRequest(req)
	if err != nil {
		return statusCode, err
	}

	return statusCode, nil
}

type InvoiceUpdate struct {
	Metadata InvoiceMetadata `json:"metadata,omitempty"`
}

func (c *Client) UpdateInvoice(storeID *StoreID, invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	dataReq, err := json.Marshal(invoiceUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return &dataRes, statusCode, nil
}
