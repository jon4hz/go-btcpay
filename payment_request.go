package btcpay

import (
	"context"
	"fmt"
)

// View information about the existing payment requests
func (c *Client) GetPaymentRequests(ctx context.Context, storeID *StoreID) ([]*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", c.URL, *storeID)
	var dataRes []*PaymentRequestResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetPaymentRequests(ctx context.Context) ([]*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", s.Client.URL, s.ID)
	var dataRes []*PaymentRequestResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// Create a new payment request
func (c *Client) CreatePaymentRequest(ctx context.Context, storeID *StoreID, paymentRequestRequest *PaymentRequestRequest) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", c.URL, *storeID)
	var dataRes PaymentRequestResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, paymentRequestRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) CreatePaymentRequest(ctx context.Context, paymentRequestRequest *PaymentRequestRequest) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests", s.Client.URL, s.ID)
	var dataRes PaymentRequestResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, paymentRequestRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// View information about the specified payment request
func (c *Client) GetPaymentRequest(ctx context.Context, storeID *StoreID, paymentRequestID *PaymentRequestID) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
	var dataRes PaymentRequestResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) GetPaymentRequest(ctx context.Context, paymentRequestID *PaymentRequestID) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
	var dataRes PaymentRequestResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (p *PaymentRequest) GetPaymentRequest(ctx context.Context) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	var dataRes PaymentRequestResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Archives the specified payment request.
func (c *Client) ArchivePaymentRequest(ctx context.Context, storeID *StoreID, paymentRequestID *PaymentRequestID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) ArchivePaymentRequest(ctx context.Context, paymentRequestID *PaymentRequestID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
	statusCode, err := s.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *PaymentRequest) ArchivePaymentRequest(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	statusCode, err := p.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// Update a payment request
func (c *Client) UpdatePaymentRequest(ctx context.Context, storeID *StoreID, paymentRequestID *PaymentRequestID, paymentRequestUpdate *PaymentRequestRequest) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", c.URL, *storeID, *paymentRequestID)
	var dataRes PaymentRequestResponse
	statusCode, err := c.doRequest(ctx, endpoint, "PUT", &dataRes, paymentRequestUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UpdatePaymentRequest(ctx context.Context, paymentRequestID *PaymentRequestID, paymentRequestUpdate *PaymentRequestRequest) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", s.Client.URL, s.ID, *paymentRequestID)
	var dataRes PaymentRequestResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "PUT", &dataRes, paymentRequestUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (p *PaymentRequest) UpdatePaymentRequest(ctx context.Context, paymentRequestUpdate *PaymentRequestRequest) (*PaymentRequestResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/payment-requests/%s", p.Client.URL, p.Store.ID, p.ID)
	var dataRes PaymentRequestResponse
	statusCode, err := p.Client.doRequest(ctx, endpoint, "PUT", &dataRes, paymentRequestUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
