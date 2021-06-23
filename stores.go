package btcpay

import (
	"context"
	"fmt"
)

// View information about the available stores
func (c *Client) GetStores(ctx context.Context) ([]*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores", c.URL)
	var dataRes []*StoreResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// create a new store
func (c *Client) CreateStore(ctx context.Context, storeRequest *StoreRequest) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores", c.URL)
	var dataRes StoreResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, storeRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// View information about the specified store
func (c *Client) GetStore(ctx context.Context, storeID *StoreID) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	var dataRes StoreResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) GetStore(ctx context.Context) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	var dataRes StoreResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) UpdateStore(ctx context.Context, storeID *StoreID, storeUpdate *StoreUpdate) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	var dataRes StoreResponse
	statusCode, err := c.doRequest(ctx, endpoint, "PUT", &dataRes, storeUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UpdateStore(ctx context.Context, storeUpdate *StoreUpdate) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	var dataRes StoreResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "PUT", &dataRes, storeUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Removes the specified store. If there is another user with access, only your access will be removed.
func (c *Client) RemoveStore(ctx context.Context, storeID *StoreID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) RemoveStore(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	statusCode, err := s.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
