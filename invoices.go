package btcpay

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
func (c *Client) GetInvoices(ctx context.Context, storeID *StoreID) ([]*InvoiceResponse, int, error) {
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
func (c *Client) CreateInvoice(ctx context.Context, storeID *StoreID, invoiceRequest *InvoiceRequest) (*InvoiceResponse, int, error) {
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

func (s *Store) CreateInvoice(ctx context.Context, invoiceRequest *InvoiceRequest) (*InvoiceResponse, int, error) {
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
func (c *Client) GetInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
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

func (s *Store) GetInvoice(ctx context.Context, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
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
func (c *Client) ArchiveInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (int, error) {
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

func (s *Store) ArchiveInvoice(ctx context.Context, invoiceID *InvoiceID) (int, error) {
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
func (c *Client) UpdateInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
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

func (s *Store) UpdateInvoice(ctx context.Context, invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
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

func (i *Invoice) UpdateInvoice(ctx context.Context, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
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

// Enums PaymentStatus
type BTCPayPaymentStatus string

type PaymentStatus struct {
	New        BTCPayPaymentStatus
	Processing BTCPayPaymentStatus
	Expired    BTCPayPaymentStatus
	Invalid    BTCPayPaymentStatus
	Settled    BTCPayPaymentStatus
}

func GetPaymentStatus() *PaymentStatus {
	return &PaymentStatus{
		Processing: "Processing",
		Invalid:    "Invalid",
		Settled:    "Settled",
	}
}

type PaymentID string

type Payment struct {
	ID           PaymentID           `json:"id"`
	ReceivedDate int64               `json:"receivedDate"`
	Value        string              `json:"value"`
	Fee          string              `json:"fee"`
	Status       BTCPayPaymentStatus `json:"status"`
	Destination  string              `json:"destination"`
}

type InvoicePaymentMethodResponse struct {
	PaymentMethod     string    `json:"paymentMethod"`
	Destination       string    `json:"destination"`
	PaymentLink       string    `json:"paymentLink,omitempty"`
	Rate              string    `json:"rate"`
	PaymentMethodPaid string    `json:"paymentMethodPaid"`
	TotalPaid         string    `json:"totalPaid"`
	Due               string    `json:"due"`
	Amount            string    `json:"amount"`
	NetworkFee        string    `json:"networkFee"`
	Payments          []Payment `json:"payments"`
}

// View information about the specified invoice's payment methods
func (c *Client) GetInvoicePaymentMethod(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", c.URL, *storeID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*InvoicePaymentMethodResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetInvoicePaymentMethod(ctx context.Context, invoiceID *InvoiceID) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", s.Client.URL, s.ID, *invoiceID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*InvoicePaymentMethodResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (i *Invoice) GetInvoicePaymentMethod(ctx context.Context) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", i.Client.URL, i.Store.ID, i.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := i.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*InvoicePaymentMethodResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

type MarkInvoiceStatusRequest struct {
	Status BTCPayInvoiceStatusMark `json:"status"`
}

// Mark an invoice as invalid or settled.
func (c *Client) MarkInvoiceStatus(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
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

func (s *Store) MarkInvoiceStatus(ctx context.Context, invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
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

func (i *Invoice) MarkInvoiceStatus(ctx context.Context, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
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
func (c *Client) UnarchiveInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
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

func (s *Store) UnarchiveInvoice(ctx context.Context, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
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
