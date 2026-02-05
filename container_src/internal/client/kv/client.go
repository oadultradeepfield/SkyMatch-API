package kv

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"server/internal/config"
)

type Client interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Put(ctx context.Context, key string, value []byte, ttlSeconds int) error
}

type KVClient struct {
	http        *http.Client
	baseURL     string
	accountID   string
	namespaceID string
	apiToken    string
}

func NewClient(cfg config.KVConfig) *KVClient {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}
	return &KVClient{
		http: &http.Client{
			Timeout:   cfg.Timeout,
			Transport: transport,
		},
		baseURL:     cfg.BaseURL,
		accountID:   cfg.AccountID,
		namespaceID: cfg.NamespaceID,
		apiToken:    cfg.APIToken,
	}
}

func (c *KVClient) Get(ctx context.Context, key string) ([]byte, bool, error) {
	reqURL := fmt.Sprintf("%s/accounts/%s/storage/kv/namespaces/%s/values/%s",
		c.baseURL, c.accountID, c.namespaceID, url.PathEscape(key))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, false, fmt.Errorf("kv get: create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiToken)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("kv get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, false, nil
	}

	if resp.StatusCode >= 400 {
		return nil, false, fmt.Errorf("kv get: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, fmt.Errorf("kv get: read body: %w", err)
	}

	return data, true, nil
}

func (c *KVClient) Put(ctx context.Context, key string, value []byte, ttlSeconds int) error {
	reqURL := fmt.Sprintf("%s/accounts/%s/storage/kv/namespaces/%s/values/%s?expiration_ttl=%d",
		c.baseURL, c.accountID, c.namespaceID, url.PathEscape(key), ttlSeconds)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, reqURL, bytes.NewReader(value))
	if err != nil {
		return fmt.Errorf("kv put: create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("kv put: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("kv put: status %d", resp.StatusCode)
	}

	return nil
}
