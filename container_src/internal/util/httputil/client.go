package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	http    *http.Client
	baseURL string
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}
	return &Client{
		http: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		baseURL: baseURL,
	}
}

func (c *Client) Get(ctx context.Context, path string, v any) error {
	return c.GetWithParams(ctx, path, nil, v)
}

func (c *Client) GetWithParams(ctx context.Context, path string, params url.Values, v any) error {
	reqURL := c.baseURL + path
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("GET %s: %w", path, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("GET %s: status %d", path, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) PostForm(ctx context.Context, path string, fields map[string]string, file io.Reader, filename string) (*http.Response, error) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	jsonData, err := json.Marshal(fields)
	if err != nil {
		return nil, fmt.Errorf("marshal fields: %w", err)
	}
	if writeErr := w.WriteField("request-json", string(jsonData)); writeErr != nil {
		return nil, fmt.Errorf("write field: %w", writeErr)
	}

	if file != nil {
		part, createErr := w.CreateFormFile("file", filename)
		if createErr != nil {
			return nil, fmt.Errorf("create form file: %w", createErr)
		}
		if _, copyErr := io.Copy(part, file); copyErr != nil {
			return nil, fmt.Errorf("copy file: %w", copyErr)
		}
	}

	if closeErr := w.Close(); closeErr != nil {
		return nil, fmt.Errorf("close writer: %w", closeErr)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST %s: %w", path, err)
	}
	return resp, nil
}

func (c *Client) PostFormDecode(ctx context.Context, path string, fields map[string]string, file io.Reader, filename string, v any) error {
	resp, err := c.PostForm(ctx, path, fields, file, filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("POST %s: status %d", path, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
