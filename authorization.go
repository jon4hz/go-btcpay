package btcpay

import (
	"context"
	"fmt"
)

// !Needs testing

func (c *Client) Authorize(ctx context.Context, authRequest *AuthorizationRequest) (int, error) {
	endpoint := fmt.Sprintf("%s/api-keys/authorize", c.URL)
	statusCode, err := c.doRequest(ctx, endpoint, "GET", nil, authRequest)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
