package btcpay

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL    string
	APIKey string
}

type BasicClient struct {
	URL      string
	Username string
	Password string
}

func NewClient(url, apiKey string) *Client {
	return &Client{
		URL:    url,
		APIKey: apiKey,
	}
}

func NewBasicClient(url, username, password string) *BasicClient {
	return &BasicClient{
		URL:      url,
		Username: username,
		Password: password,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.APIKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *BasicClient) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Username, c.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

type BTCPayClient interface {
	doRequest(*http.Request) ([]byte, error)
	GetHealth() (*BTCPayHealthResponse, error)
}
