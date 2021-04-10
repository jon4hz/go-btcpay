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
type BTCPayPayoutStatus string

type PayoutStatus struct {
	AwaitingApproval BTCPayPayoutStatus
	AwaitingPayment  BTCPayPayoutStatus
	InProgress       BTCPayPayoutStatus
	Completed        BTCPayPayoutStatus
	Cancelled        BTCPayPayoutStatus
}

func GetPayoutStatus() *PayoutStatus {
	return &PayoutStatus{
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
	ID                  PayoutID           `json:"id"`
	Revision            int64              `json:"revision"`
	PullPaymentID       PullPaymentID      `json:"pullPaymentId"`
	Date                string             `json:"date"`
	Destination         string             `json:"destination"`
	Amount              string             `json:"amount"`
	PaymentMethod       string             `json:"paymentMethod"`
	PaymentMethodAmount string             `json:"paymentMethodAmount"`
	State               BTCPayPayoutStatus `json:"state"`
}

type PayoutApproveRequest struct {
	Revision int64  `json:"revision"`
	RateRule string `json:"rateRule,omitempty"`
}

// Approve a payout
func (c *Client) ApprovePayout(ctx context.Context, storeID *StoreID, payoutID *PayoutID, PayoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", c.URL, *storeID, *payoutID)
	dataReq, err := json.Marshal(PayoutApproveRequest)
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

func (s *Store) ApprovePayout(ctx context.Context, payoutID *PayoutID, PayoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", s.Client.URL, s.ID, *payoutID)
	dataReq, err := json.Marshal(PayoutApproveRequest)
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

func (p *Payout) ApprovePayout(ctx context.Context, PayoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", p.Client.URL, p.Store.ID, p.ID)
	dataReq, err := json.Marshal(PayoutApproveRequest)
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

// Cancel the payout
func (c *Client) CancelPayout(ctx context.Context, storeID *StoreID, payoutID *PayoutID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", c.URL, *storeID, *payoutID)
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

func (s *Store) CancelPayout(ctx context.Context, payoutID *PayoutID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", s.Client.URL, s.ID, *payoutID)
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

func (p *Payout) CancelPayout(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", p.Client.URL, p.Store.ID, p.ID)
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

// Get a pull payment
func (c *Client) GetPullPayment(ctx context.Context, pullPaymentID *PullPaymentID) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s", c.URL, *pullPaymentID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
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

func (p *PullPayment) GetPullPayment(ctx context.Context) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s", p.Client.URL, p.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := p.Client.doRequest(req)
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

// Get payouts
func (c *Client) GetPayouts(ctx context.Context, pullPaymentID *PullPaymentID, includeCancelled ...bool) ([]*PayoutResponse, int, error) {
	var endpoint string
	if len(includeCancelled) > 0 {
		endpoint = fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts?includeCancelled=%t", c.URL, *pullPaymentID, includeCancelled[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", c.URL, *pullPaymentID)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PayoutResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

func (p *PullPayment) GetPayouts(ctx context.Context, includeCancelled ...bool) ([]*PayoutResponse, int, error) {
	var endpoint string
	if len(includeCancelled) > 0 {
		endpoint = fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts?includeCancelled=%t", p.Client.URL, p.ID, includeCancelled[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", p.Client.URL, p.ID)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := p.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*PayoutResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

type PayoutRequest struct {
	Destination   string `json:"destination"`
	Amount        string `json:"amount"`
	PaymentMethod string `json:"paymentMethod"`
}

// create a new payout
func (c *Client) CreatePayout(ctx context.Context, pullPaymentID *PullPaymentID, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", c.URL, *pullPaymentID)
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

func (p *PullPayment) CreatePayout(ctx context.Context, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", p.Client.URL, p.ID)
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
