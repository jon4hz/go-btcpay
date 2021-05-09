package btcpay

import (
	"context"
	"fmt"
)

// Get an array of all Invoices from a single store.
func (c *Client) GetInvoices(ctx context.Context, storeID *StoreID) ([]*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, *storeID)
	var dataRes []*InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetInvoices(ctx context.Context) ([]*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", s.Client.URL, s.ID)
	var dataRes []*InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// Create an invoice for a certain store
func (c *Client) CreateInvoice(ctx context.Context, storeID *StoreID, invoiceRequest *InvoiceRequest) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", c.URL, *storeID)
	var dataRes InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, invoiceRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) CreateInvoice(ctx context.Context, invoiceRequest *InvoiceRequest) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices", s.Client.URL, s.ID)
	var dataRes InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, invoiceRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Get a signle invoice from a single store.
func (c *Client) GetInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) GetInvoice(ctx context.Context, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (i *Invoice) GetInvoice(ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	var dataRes InvoiceResponse
	statusCode, err := i.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Archive a single invoice of a store
func (c *Client) ArchiveInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) ArchiveInvoice(ctx context.Context, invoiceID *InvoiceID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	statusCode, err := s.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (i *Invoice) ArchiveInvoice(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	statusCode, err := i.Client.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

// Update the metadata from an existing invoice
func (c *Client) UpdateInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", c.URL, *storeID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "PUT", &dataRes, invoiceUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UpdateInvoice(ctx context.Context, invoiceID *InvoiceID, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", s.Client.URL, s.ID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "PUT", &dataRes, invoiceUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (i *Invoice) UpdateInvoice(ctx context.Context, invoiceUpdate *InvoiceUpdate) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s", i.Client.URL, i.Store.ID, i.ID)
	var dataRes InvoiceResponse
	statusCode, err := i.Client.doRequest(ctx, endpoint, "PUT", &dataRes, invoiceUpdate)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// View information about the specified invoice's payment methods
func (c *Client) GetInvoicePaymentMethod(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", c.URL, *storeID, *invoiceID)
	var dataRes []*InvoicePaymentMethodResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

func (s *Store) GetInvoicePaymentMethod(ctx context.Context, invoiceID *InvoiceID) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", s.Client.URL, s.ID, *invoiceID)
	var dataRes []*InvoicePaymentMethodResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

func (i *Invoice) GetInvoicePaymentMethod(ctx context.Context) ([]*InvoicePaymentMethodResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/payment-methods", i.Client.URL, i.Store.ID, i.ID)
	var dataRes []*InvoicePaymentMethodResponse
	statusCode, err := i.Client.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}

// Mark an invoice as invalid or settled.
func (c *Client) MarkInvoiceStatus(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", c.URL, *storeID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, markInvoiceStatusRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) MarkInvoiceStatus(ctx context.Context, invoiceID *InvoiceID, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", s.Client.URL, s.ID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, markInvoiceStatusRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (i *Invoice) MarkInvoiceStatus(ctx context.Context, markInvoiceStatusRequest *MarkInvoiceStatusRequest) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/status", i.Client.URL, i.Store.ID, i.ID)
	var dataRes InvoiceResponse
	statusCode, err := i.Client.doRequest(ctx, endpoint, "POST", &dataRes, markInvoiceStatusRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// unarchive an invoice
func (c *Client) UnarchiveInvoice(ctx context.Context, storeID *StoreID, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", c.URL, *storeID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UnarchiveInvoice(ctx context.Context, invoiceID *InvoiceID) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", s.Client.URL, s.ID, *invoiceID)
	var dataRes InvoiceResponse
	statusCode, err := s.Client.doRequest(ctx, endpoint, "POST", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (i *Invoice) UnarchiveInvoice(ctx context.Context) (*InvoiceResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s/invoices/%s/unarchive", i.Client.URL, i.Store.ID, i.ID)
	var dataRes InvoiceResponse
	statusCode, err := i.Client.doRequest(ctx, endpoint, "POST", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
