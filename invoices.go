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

type Invoice struct {
	Store  *Store
	Client *Client
	ID     InvoiceID
}

type InvoiceResponse struct {
	Amount           string                        `json:"amount,omitempty"`
	Currency         string                        `json:"currency,omitempty"`
	Metadata         InvoiceMetadata               `json:"metadata,omitempty"`
	Checkout         InvoiceCheckout               `json:"checkout,omitempty"`
	ID               InvoiceID                     `json:"id"`
	CheckoutLink     string                        `json:"checkoutLink"`
	CreatedTime      int64                         `json:"createdTime"`
	ExpirationTime   int64                         `json:"expirationTime"`
	MonitoringTime   int64                         `json:"monitoringTime"`
	Status           BTCPayInvoiceStatus           `json:"status"`
	AdditionalStatus BTCPayInvoiceAdditionalStatus `json:"additionalStatus"`
}

type InvoiceMetadata map[string]interface{}

// Enums SpeedPolicy
type BTCPaySpeedPolicy string

type SpeedPolicy struct {
	HighSpeed      BTCPaySpeedPolicy
	MediumSpeed    BTCPaySpeedPolicy
	LowMediumSpeed BTCPaySpeedPolicy
	LowSpeed       BTCPaySpeedPolicy
}

func GetSpeedPolicy() *SpeedPolicy {
	return &SpeedPolicy{
		HighSpeed:      "HighSpeed",
		MediumSpeed:    "MediumSpeed",
		LowMediumSpeed: "LowMediumSpeed",
		LowSpeed:       "LowSpeed",
	}
}

// Enums InvoiceStatus
type BTCPayInvoiceStatus string

type InvoiceStatus struct {
	New        BTCPayInvoiceStatus
	Processing BTCPayInvoiceStatus
	Expired    BTCPayInvoiceStatus
	Invalid    BTCPayInvoiceStatus
	Settled    BTCPayInvoiceStatus
}

func GetInvoiceStatus() *InvoiceStatus {
	return &InvoiceStatus{
		New:        "New",
		Processing: "Processing",
		Expired:    "Expired",
		Invalid:    "Invalid",
		Settled:    "Settled",
	}
}

// Enums InvoiceAdditionalStatus
type BTCPayInvoiceAdditionalStatus string

type InvoiceAdditionalStatus struct {
	None        BTCPayInvoiceAdditionalStatus
	PaidLate    BTCPayInvoiceAdditionalStatus
	PaidPartial BTCPayInvoiceAdditionalStatus
	Marked      BTCPayInvoiceAdditionalStatus
	AddInvalid  BTCPayInvoiceAdditionalStatus
	PaidOver    BTCPayInvoiceAdditionalStatus
}

func GetInvoiceAdditionalStatus() *InvoiceAdditionalStatus {
	return &InvoiceAdditionalStatus{
		None:        "None",
		PaidLate:    "PaidLate",
		PaidPartial: "PaidPartial",
		Marked:      "Marked",
		AddInvalid:  "Invalid",
		PaidOver:    "PaidOver",
	}

}

// Enums InvoiceStatusMark
type BTCPayInvoiceStatusMark string

type InvoiceStatusMark struct {
	MarkInvalid  BTCPayInvoiceStatusMark
	MarkComplete BTCPayInvoiceStatusMark
}

func GetInvoiceStatusMark() *InvoiceStatusMark {
	return &InvoiceStatusMark{
		MarkInvalid:  "Invalid",
		MarkComplete: "Complete",
	}
}

type InvoiceCheckout struct {
	SpeedPolicy       BTCPaySpeedPolicy `json:"speedPolicy,omitempty"`
	PaymentMethods    []string          `json:"paymentMethods,omitempty"`
	ExpirationMinutes int               `json:"expirationMinutes,omitempty"`
	MonitoringMinutes int               `json:"monitoringMinutes,omitempty"`
	PaymentTolerance  float64           `json:"paymentTolerance,omitempty"`
	RedirectURL       string            `json:"redirectURL,omitempty"`
	DefaultLanguage   string            `json:"defaultLanguage,omitempty"`
}

// Get an array of all Invoices from a single store.
func (c *Client) GetInvoices(storeID *StoreID, ctx context.Context) ([]*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, *storeID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetInvoices(ctx context.Context) ([]*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", s.Client.URL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*InvoiceResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

type InvoiceRequest struct {
	Amount          string          `json:"amount"`
	Currency        string          `json:"currency,omitempty"`
	Metadata        InvoiceMetadata `json:"metadata,omitempty"`
	InvoiceCheckout InvoiceCheckout `json:"checkout,omitempty"`
}

// Create an invoice for a certain store
func (c *Client) CreateInvoice(storeID *StoreID, invoiceRequest *InvoiceRequest, ctx context.Context) (*InvoiceResponse, int, error) {
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

func (s *Store) CreateInvoice(invoiceRequest *InvoiceRequest, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", s.Client.URL, s.ID)
	dataReq, err := json.Marshal(invoiceRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
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

// Get a signle invoice from a single store.
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

func (s *Store) GetInvoice(invoiceID *InvoiceID, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
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

func (i *Invoice) GetInvoice(ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
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

// Archive a single invoice of a store
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

func (s *Store) ArchiveInvoice(invoiceID *InvoiceID, ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (i *Invoice) ArchiveInvoice(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := i.Client.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

type InvoiceUpdate struct {
	Metadata InvoiceMetadata `json:"metadata,omitempty"`
}

// Update the metadata from an existing invoice
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

func (s *Store) UpdateInvoice(invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	dataReq, err := json.Marshal(invoiceUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
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

func (i *Invoice) UpdateInvoice(invoiceUpdate *InvoiceUpdate, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	dataReq, err := json.Marshal(invoiceUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
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

type MarkInvoiceStatusRequest struct {
	Status BTCPayInvoiceStatusMark `json:"status"`
}

// Mark an invoice as invalid or settled.
func (c *Client) MarkInvoiceStatus(storeID *StoreID, invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", c.URL, *storeID, *invoiceID)
	dataReq, err := json.Marshal(markInvoiceStatusRequest)
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

func (s *Store) MarkInvoiceStatus(invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", s.Client.URL, s.ID, *invoiceID)
	dataReq, err := json.Marshal(markInvoiceStatusRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
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

func (i *Invoice) MarkInvoiceStatus(markInvoiceStatusRequest *MarkInvoiceStatusRequest, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", i.Client.URL, i.Store.ID, i.ID)
	dataReq, err := json.Marshal(markInvoiceStatusRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
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

// unarchive an invoice
func (c *Client) UnarchiveInvoice(storeID *StoreID, invoiceID *InvoiceID, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", c.URL, *storeID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
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

func (s *Store) UnarchiveInvoice(invoiceID *InvoiceID, ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", s.Client.URL, s.ID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
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

func (i *Invoice) UnarchiveInvoice(ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", i.Client.URL, i.Store.ID, i.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
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
