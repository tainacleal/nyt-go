package nyttop

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

// TopStories receives a section and returns the top NYT stories from that section.
func (c *Client) TopStories(ctx context.Context, section Section) ([]Article, error) {
	if _, ok := Sections[section]; !ok {
		return nil, ErrInvalidSection
	}

	url := generateURL(section, c.apiKey)
	resp, err := c.request(ctx, url)
	if err != nil {
		return nil, err
	}

	return resp.Articles, nil
}

// TopNStories receives a section and returns the top N stories from that section.
func (c *Client) TopNStories(ctx context.Context, section Section, topN int) ([]Article, error) {
	if _, ok := Sections[section]; !ok {
		return nil, ErrInvalidSection
	}

	url := generateURL(section, c.apiKey)
	resp, err := c.request(ctx, url)
	if err != nil {
		return nil, err
	}

	if topN > resp.NumResults {
		topN = resp.NumResults
	}

	return resp.Articles[:topN], nil
}

func (c *Client) request(ctx context.Context, url string) (TopStoriesResponse, error) {
	var response TopStoriesResponse
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("request status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func generateURL(section Section, apiKey string) string {
	return fmt.Sprintf("%s/%s.json?api-key=%s", baseURL, section, apiKey)
}
