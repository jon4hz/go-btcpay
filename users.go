package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserID string

type UserResponse struct {
	ID                        UserID   `json:"id"`
	Email                     string   `json:"email"`
	EmailConfirmed            bool     `json:"emailConfirmed"`
	RequiresEmailConfirmation bool     `json:"requiresEmailConfirmation"`
	Created                   int64    `json:"created,omitempty"`
	Roles                     []string `json:"roles"`
}

// View information about the current user
func (c *Client) GetUser(ctx context.Context) (*UserResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes UserResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

type CreateUserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	IsAdministrator bool   `json:"isAdministrator"`
}

func (c *Client) CreateUser(userRequest *CreateUserRequest, ctx context.Context) (*UserResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users", c.URL)
	dataReq, err := json.Marshal(userRequest)
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
	var dataRes UserResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}
