package nova

import (
	"context"
	"fmt"
	"io"

	"server/internal/config"
	"server/internal/util/httputil"
)

type Client struct {
	http    *httputil.Client
	baseURL string
}

func NewClient(cfg config.NovaConfig) *Client {
	return &Client{
		http:    httputil.NewClient(cfg.BaseURL, cfg.Timeout),
		baseURL: cfg.BaseURL,
	}
}

func (c *Client) Login(ctx context.Context, apiKey string) (string, error) {
	var r LoginResponse
	if err := c.http.PostFormDecode(ctx, "/api/login", map[string]string{"apikey": apiKey}, nil, "", &r); err != nil {
		return "", fmt.Errorf("login: %w", err)
	}
	if r.Status != "success" {
		return "", fmt.Errorf("login failed: %s", r.Message)
	}
	return r.Session, nil
}

func (c *Client) Upload(ctx context.Context, session string, file io.Reader, filename string) (int, error) {
	var r UploadResponse
	if err := c.http.PostFormDecode(ctx, "/api/upload", map[string]string{"session": session, "publicly_visible": "n"}, file, filename, &r); err != nil {
		return 0, fmt.Errorf("upload: %w", err)
	}
	if r.Status != "success" {
		return 0, fmt.Errorf("upload failed")
	}
	return r.SubID, nil
}

func (c *Client) GetSubmission(ctx context.Context, subID int) (*Submission, error) {
	var s Submission
	if err := c.http.Get(ctx, fmt.Sprintf("/api/submissions/%d", subID), &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) GetJobStatus(ctx context.Context, jobID int) (string, error) {
	var r JobStatusResponse
	if err := c.http.Get(ctx, fmt.Sprintf("/api/jobs/%d", jobID), &r); err != nil {
		return "", err
	}
	return r.Status, nil
}

func (c *Client) GetJobInfo(ctx context.Context, jobID int) (*JobInfo, error) {
	var info JobInfo
	if err := c.http.Get(ctx, fmt.Sprintf("/api/jobs/%d/info/", jobID), &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *Client) GetAnnotations(ctx context.Context, jobID int) ([]Annotation, error) {
	var r struct {
		Annotations []Annotation `json:"annotations"`
	}
	if err := c.http.Get(ctx, fmt.Sprintf("/api/jobs/%d/annotations/", jobID), &r); err != nil {
		return nil, err
	}
	return r.Annotations, nil
}

func (c *Client) AnnotatedImageURL(jobID int) string {
	return fmt.Sprintf("%s/annotated_display/%d", c.baseURL, jobID)
}
