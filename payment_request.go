package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PaymentRequestID string

type PaymentRequest struct {
	Store  *Store
	Client *Client
	ID     PaymentRequestID
}

// Enums PaymentRequestStatus
type BTCPayPaymentRequestStatus string

type PaymentRequestStatus struct {
	Pending   BTCPayInvoiceStatus
	Completed BTCPayInvoiceStatus
	Expired   BTCPayInvoiceStatus
}

func GetPaymentRequestStatus() *PaymentRequestStatus {
	return &PaymentRequestStatus{
		Pending:   "Pending",
		Completed: "Completed",
		Expired:   "Expired",
	}
}

type PaymentRequestRequest struct {
	Amount                    float64 `json:"amount"`
	Title                     string  `json:"title"`
	Currency                  string  `json:"currency"`
	Email                     string  `json:"email,omitempty"`
	Description               string  `json:"description,omitempty"`
	ExpiryDate                int64   `json:"expiryDate,omitempty"`
	EmbeddedCSS               string  `json:"embeddedCSS,omitempty"`
	CustomCSSLink             string  `json:"customCSSLink,omitempty"`
	AllowCustomPaymentAmounts bool    `json:"allowCustomPaymentAmounts,omitempty"`
}

type PaymentRequestResponse struct {
	ID       PaymentRequestID           `json:"id"`
	Status   BTCPayPaymentRequestStatus `json:"status"`
	Created  string                     `json:"created"`
	Archived bool                       `json:"archived"`
	PaymentRequestRequest
}

// View information about the existing payment requests
func (c *Client) GetPaymentRequests(storeID *StoreID, ctx context.Context) ([]*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", c.URL, *storeID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetPaymentRequests(ctx context.Context) ([]*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", s.Client.URL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

// Create a new payment request
func (c *Client) CreatePaymentRequest(storeID *StoreID, paymentRequestRequest *PaymentRequestRequest, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", c.URL, *storeID)
	dataReq, err := json.Marshal(paymentRequestRequest)
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
	var dataRes *PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) CreatePaymentRequest(paymentRequestRequest *PaymentRequestRequest, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", s.Client.URL, s.ID)
	dataReq, err := json.Marshal(paymentRequestRequest)
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
	var dataRes *PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

// View information about the specified payment request
func (c *Client) GetPaymentRequest(storeID *StoreID, paymentRequestID *PaymentRequestID, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) GetPaymentRequest(paymentRequestID *PaymentRequestID, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (p *PaymentRequest) GetPaymentRequest(ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := p.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

// Archives the specified payment request.
func (c *Client) ArchivePaymentRequest(storeID *StoreID, paymentRequestID *PaymentRequestID, ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
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

func (s *Store) ArchivePaymentRequest(paymentRequestID *PaymentRequestID, ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
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

func (p *PaymentRequest) ArchivePaymentRequest(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := p.Client.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// Update a payment request
func (c *Client) UpdatePaymentRequest(storeID *StoreID, paymentRequestID *PaymentRequestID, paymentRequestUpdate *PaymentRequestRequest, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
	dataReq, err := json.Marshal(paymentRequestUpdate)
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
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UpdatePaymentRequest(paymentRequestID *PaymentRequestID, paymentRequestUpdate *PaymentRequestRequest, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
	dataReq, err := json.Marshal(paymentRequestUpdate)
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
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (p *PaymentRequest) UpdatePaymentRequest(paymentRequestUpdate *PaymentRequestRequest, ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	dataReq, err := json.Marshal(paymentRequestUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := p.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes PaymentRequestResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}
