package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) GetNotifications(ctx context.Context, seen ...bool) ([]*NotificationResponse, int, error) {
	var endpoint string
	var dataRes []*NotificationResponse
	if len(seen) > 0 {
		endpoint = fmt.Sprintf("%s/api/v1/users/me/notifications?%t", c.URL, seen[0])
	} else {
		endpoint = fmt.Sprintf("%s/api/v1/users/me/notifications", c.URL)
	}
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// View information about the specified notification
func (c *Client) GetNotification(ctx context.Context, notificationID *NotificationID) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", c.URL, *notificationID)
	var dataRes NotificationResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (n *Notification) GetNotification(ctx context.Context) (*NotificationResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", n.Client.URL, n.ID)
	var dataRes NotificationResponse
	statusCode, err := n.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
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
	var dataRes NotificationResponse
	statusCode, err := c.doRequest(ctx, endpoint, "PUT", &dataRes, dataReq)
	if err != nil {
		return nil, statusCode, err
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
	if err != nil {
		return nil, 0, err
	}
	var dataRes NotificationResponse
	statusCode, err := n.Client.doRequest(ctx, endpoint, "PUT", &dataRes, dataReq)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Removes the specified notification.
func (c *Client) RemoveNotification(ctx context.Context, notificationID *NotificationID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", c.URL, *notificationID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (n *Notification) RemoveNotification(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/me/notifications/%s", n.Client.URL, n.ID)
	statusCode, err := n.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
