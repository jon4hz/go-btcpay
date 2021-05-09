package btcpay

import (
	"context"
	"fmt"
)

// Get the pull payments of a store
func (c *Client) GetPullPayments(ctx context.Context, storeID *StoreID, includeArchived ...bool) ([]*PullPaymentResponse, int, error) {
	var endpoint string
	if len(includeArchived) > 0 {
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments?includeArchived=%t", c.URL, *storeID, includeArchived[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", c.URL, *storeID)
	}
	var dataRes []*PullPaymentResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
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
	var dataRes []*PullPaymentResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

func (c *Client) CreatePullPayment(ctx context.Context, storeID *StoreID, pullPaymentRequest *PullPaymentRequest) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", c.URL, *storeID)
	var dataRes PullPaymentResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, pullPaymentRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) CreatePullPayment(ctx context.Context, pullPaymentRequest *PullPaymentRequest) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments", s.Client.URL, s.ID)
	var dataRes PullPaymentResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, pullPaymentRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) ArchivePullPayment(ctx context.Context, storeID *StoreID, pullPaymentID *PullPaymentID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", c.URL, *storeID, *pullPaymentID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) ArchivePullPayment(ctx context.Context, pullPaymentID *PullPaymentID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", s.Client.URL, s.ID, *pullPaymentID)
	statusCode, err := s.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *PullPayment) ArchivePullPayment(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/pull-payments/%s", p.Client.URL, p.Store.ID, p.ID)
	statusCode, err := p.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// Approve a payout
func (c *Client) ApprovePayout(ctx context.Context, storeID *StoreID, payoutID *PayoutID, payoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", c.URL, *storeID, *payoutID)
	var dataRes PayoutResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, payoutApproveRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) ApprovePayout(ctx context.Context, payoutID *PayoutID, payoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", s.Client.URL, s.ID, *payoutID)
	var dataRes PayoutResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, payoutApproveRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (p *Payout) ApprovePayout(ctx context.Context, payoutApproveRequest *PayoutApproveRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", p.Client.URL, p.Store.ID, p.ID)
	var dataRes PayoutResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "POST", &dataRes, payoutApproveRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Cancel the payout
func (c *Client) CancelPayout(ctx context.Context, storeID *StoreID, payoutID *PayoutID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", c.URL, *storeID, *payoutID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) CancelPayout(ctx context.Context, payoutID *PayoutID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", s.Client.URL, s.ID, *payoutID)
	statusCode, err := s.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Payout) CancelPayout(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payouts/%s", p.Client.URL, p.Store.ID, p.ID)
	statusCode, err := p.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// Get a pull payment
func (c *Client) GetPullPayment(ctx context.Context, pullPaymentID *PullPaymentID) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s", c.URL, *pullPaymentID)
	var dataRes PullPaymentResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (p *PullPayment) GetPullPayment(ctx context.Context) (*PullPaymentResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s", p.Client.URL, p.ID)
	var dataRes PullPaymentResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
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
	var dataRes []*PayoutResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
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
	var dataRes []*PayoutResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// create a new payout
func (c *Client) CreatePayout(ctx context.Context, pullPaymentID *PullPaymentID, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", c.URL, *pullPaymentID)
	var dataRes PayoutResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, payoutRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (p *PullPayment) CreatePayout(ctx context.Context, payoutRequest *PayoutRequest) (*PayoutResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/pull-payments/%s/payouts", p.Client.URL, p.ID)
	var dataRes PayoutResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "POST", &dataRes, payoutRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
