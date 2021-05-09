package btcpay

import (
	"context"
	"fmt"
)

// View information about the current user
func (c *Client) GetUser(ctx context.Context) (*UserResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me", c.URL)
	var dataRes UserResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) CreateUser(ctx context.Context, userRequest *UserRequest) (*UserResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users", c.URL)
	var dataRes UserResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, &userRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
