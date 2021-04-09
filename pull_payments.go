package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PullPaymentID string

type PullPayment struct {
	Store  *Store
	Client *Client
	ID     PullPaymentID
}

type PullPaymentResponse struct {
	ID       PullPaymentID `json:"id"`
	Name     string        `json:"name"`
	Currency string        `json:"currency"`
	Amount   string        `json:"amount"`
	Period   int64         `json:"period,omitempty"`
	Archived bool          `json:"archived"`
	ViewLink string        `json:"viewLink"`
}

// Get the pull payments of a store
func (c *Client) GetPullPayments(ctx context.Context, storeID *StoreID, includeArchived ...bool) ([]*PullPaymentResponse, int, error) {
	var endpoint string
	if len(includeArchived) > 0 {
		fmt.Println(includeArchived[0])
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments?includeArchived=%t", c.URL, *storeID, includeArchived[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", c.URL, *storeID)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PullPaymentResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetPullPayments(ctx context.Context, includeArchived ...bool) ([]*PullPaymentResponse, int, error) {
	var endpoint string
	if len(includeArchived) > 0 {
		fmt.Println(includeArchived[0])
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments?includeArchived=%t", s.Client.URL, s.ID, includeArchived[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", s.Client.URL, s.ID)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PullPaymentResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

type PullPaymentRequest struct {
	Name           string   `json:"name,omitempty"`
	Amount         string   `json:"amount"`
	Currency       string   `json:"currency"`
	Period         int64    `json:"period,omitempty"`
	StartsAt       int64    `json:"startsAt,omitempty"`
	ExpiresAt      int64    `json:"expiresAt,omitempty"`
	PaymentMethods []string `json:"paymentMethods"`
}

func (c *Client) CreatePullPayment(ctx context.Context, storeID *StoreID, pullPaymentRequest *PullPaymentRequest) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", c.URL, *storeID)
	dataReq, err := json.Marshal(pullPaymentRequest)
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
	var dataRes PullPaymentResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) CreatePullPayment(ctx context.Context, pullPaymentRequest *PullPaymentRequest) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", s.Client.URL, s.ID)
	dataReq, err := json.Marshal(pullPaymentRequest)
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
	var dataRes PullPaymentResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) ArchivePullPayment(ctx context.Context, storeID *StoreID, pullPaymentID *PullPaymentID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", c.URL, *storeID, *pullPaymentID)
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

func (s *Store) ArchivePullPayment(ctx context.Context, pullPaymentID *PullPaymentID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", s.Client.URL, s.ID, *pullPaymentID)
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

func (p *PullPayment) ArchivePullPayment(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", p.Client.URL, p.Store.ID, p.ID)
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

// Enums NetworkFeeMode
type BTCPayPullPaymentStatus string

type PullPaymentStatus struct {
	AwaitingApproval BTCPayPullPaymentStatus
	AwaitingPayment  BTCPayPullPaymentStatus
	InProgress       BTCPayPullPaymentStatus
	Completed        BTCPayPullPaymentStatus
	Cancelled        BTCPayPullPaymentStatus
}

func GetPullPaymentStatus() *PullPaymentStatus {
	return &PullPaymentStatus{
		AwaitingApproval: "AwaitingApproval",
		AwaitingPayment:  "AwaitingPayment",
		InProgress:       "InProgress",
		Completed:        "Completed",
		Cancelled:        "Cancelled",
	}
}

type PayoutID string

type Payout struct {
	Store  *Store
	Client *Client
	ID     PayoutID
}

type PayoutResponse struct {
	ID                  PayoutID      `json:"id"`
	Revision            int64         `json:"revision"`
	PullPaymentID       PullPaymentID `json:"pullPaymentId"`
	Date                string        `json:"date"`
	Destination         string        `json:"destination"`
	Amount              string        `json:"amount"`
	PaymentMethod       string        `json:"paymentMethod"`
	PaymentMethodAmount string        `json:"paymentMethodAmount"`
}

type PayoutRequest struct {
	Revision int64  `json:"revision"`
	RateRule string `json:"rateRule,omitempty"`
}

// Approve a payout
func (c *Client) ApprovePayout(ctx context.Context, storeID *StoreID, payoutID *PayoutID, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", c.URL, *storeID, *payoutID)
	dataReq, err := json.Marshal(payoutRequest)
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
	var dataRes PayoutResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) ApprovePayout(ctx context.Context, payoutID *PayoutID, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", s.Client.URL, s.ID, *payoutID)
	dataReq, err := json.Marshal(payoutRequest)
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
	var dataRes PayoutResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (p *Payout) ApprovePayout(ctx context.Context, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", p.Client.URL, p.Store.ID, p.ID)
	dataReq, err := json.Marshal(payoutRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := p.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes PayoutResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}
