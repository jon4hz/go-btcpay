package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PullPaymentID string

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
