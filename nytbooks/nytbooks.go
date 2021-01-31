package nytbooks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string, options ...Option) *Client {
	cfg := newConfig(options...)

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: cfg.timeout,
		},
	}
}

// BestSellersListsOptions returns basic information about every best seller list served by the NYT Books API.
func (c *Client) BestSellersListsOptions(ctx context.Context) (BestSellersListOptionsResponse, error) {
	var response BestSellersListOptionsResponse
	url := generateURL(bestSellersListNameURL, c.apiKey, nil)
	err := c.request(ctx, url, &response)
	return response, err
}

// LatestBestSellers returns the latest best sellers of a speciffic list. It expects a list name and offset value for pagination.
func (c *Client) LatestBestSellers(ctx context.Context, listName string, offset int) (BestSellersListResponse, error) {
	var response BestSellersListResponse
	param := url.Values{}
	param.Add("list", listName)

	if offset > 0 {
		param.Add("offset", strconv.Itoa(offset))
	}

	url := generateURL(bestSellersListLatestURL, c.apiKey, param)

	err := c.request(ctx, url, &response)
	return response, err
}

func (c *Client) request(ctx context.Context, url string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&response)
}

func generateURL(baseURL string, apiKey string, params url.Values) string {
	genURL := fmt.Sprintf("%s?api-key=%s", baseURL, apiKey)
	if params != nil {
		return fmt.Sprintf("%s&%s", genURL, params.Encode())
	}
	return genURL
}
