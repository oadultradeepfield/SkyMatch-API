package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Client struct {
	http    *http.Client
	baseURL string
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		http:    &http.Client{Timeout: timeout},
		baseURL: baseURL,
	}
}

func (c *Client) Get(path string, v any) error {
	resp, err := c.http.Get(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("GET %s: %w", path, err)
	}
	defer func(Body io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("GET %s: status %d", path, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) PostForm(path string, fields map[string]string, file io.Reader, filename string) (*http.Response, error) {
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

	req, err := http.NewRequest(http.MethodPost, c.baseURL+path, body)
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

func (c *Client) PostFormDecode(path string, fields map[string]string, file io.Reader, filename string, v any) error {
	resp, err := c.PostForm(path, fields, file, filename)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("POST %s: status %d", path, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
