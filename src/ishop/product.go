package ishop

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  	string
	ProductId	string
	httpClient *http.Client
}

type Product struct {
	Id 	int
}

// Get the one product
func (c *Client) GetProduct([]Product, error)  {
	rel := &url.URL{Path: "/admin/products/" + c.ProductId +".json"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	return products, err
}

// Get the list of products
func (c *Client) GetProducts() ([]Product, error) {
	rel := &url.URL{Path: "/admin/products.json"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	return products, err
}

// Get the total of products
func (c *Client) GetCount() ([]Total, error) {
	rel := &url.URL{Path: "/admin/products/count.json"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var totals []Total
	err = json.NewDecoder(resp.Body).Decode(&totals)
	return totals, err
}
