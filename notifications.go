package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Notification struct {
	ID     NotificationID
	Client *Client
}

type NotificationID string

type NotificationResponse struct {
	ID          NotificationID `json:"id"`
	Body        string         `json:"body"`
	Link        string         `json:"link"`
	CreatedTime int64          `json:"createdTime"`
	Seen        bool           `json:"seen"`
}

func (c *Client) GetNotifications(ctx context.Context, seen ...bool) ([]*NotificationResponse, int, error) {
	var endpoint string
	if len(seen) > 0 {
		endpoint = fmt.Sprintf("%s/api/v1/users/me/notifications?%t", c.URL, seen[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/users/me/notifications", c.URL)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*NotificationResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return dataRes, statusCode, nil
}

// View information about the specified notification
func (c *Client) GetNotification(ctx context.Context, notificationID *NotificationID) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", c.URL, *notificationID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes NotificationResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (n *Notification) GetNotification(ctx context.Context) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", n.Client.URL, n.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := n.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes NotificationResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

type UpdateNotification struct {
	Seen bool `json:"seen"`
}

// Updates the notification
func (c *Client) UpdateNotification(ctx context.Context, notificationID *NotificationID, updateNotification ...*UpdateNotification) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", c.URL, *notificationID)
	var dataReq []byte
	var err error
	if len(updateNotification) > 0 {
		dataReq, err = json.Marshal(updateNotification[0])
	} else {
		dataReq = []byte(`{"seen":null}`)
	}
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
	var dataRes NotificationResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (n *Notification) UpdateNotification(ctx context.Context, updateNotification ...*UpdateNotification) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", n.Client.URL, n.ID)
	var dataReq []byte
	var err error
	if len(updateNotification) > 0 {
		dataReq, err = json.Marshal(updateNotification[0])
	} else {
		dataReq = []byte(`{"seen":null}`)
	}
	fmt.Println(string(dataReq))
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := n.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes NotificationResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

// Removes the specified notification.
func (c *Client) RemoveNotification(ctx context.Context, notificationID *NotificationID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", c.URL, *notificationID)
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

func (n *Notification) RemoveNotification(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", n.Client.URL, n.ID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := n.Client.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
