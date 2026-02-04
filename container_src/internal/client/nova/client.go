package nova

import (
	"fmt"
	"io"
	"time"

	"server/internal/util/httputil"
)

const baseURL = "https://nova.astrometry.net"

type Client struct {
	http *httputil.Client
}

func NewClient() *Client {
	return &Client{http: httputil.NewClient(baseURL, 30*time.Second)}
}

func (c *Client) Login(apiKey string) (string, error) {
	var r LoginResponse
	if err := c.http.PostFormDecode("/api/login", map[string]string{"apikey": apiKey}, nil, "", &r); err != nil {
		return "", fmt.Errorf("login: %w", err)
	}
	if r.Status != "success" {
		return "", fmt.Errorf("login failed: %s", r.Message)
	}
	return r.Session, nil
}

func (c *Client) Upload(session string, file io.Reader, filename string) (int, error) {
	var r UploadResponse
	if err := c.http.PostFormDecode("/api/upload", map[string]string{"session": session, "publicly_visible": "n"}, file, filename, &r); err != nil {
		return 0, fmt.Errorf("upload: %w", err)
	}
	if r.Status != "success" {
		return 0, fmt.Errorf("upload failed")
	}
	return r.SubID, nil
}

func (c *Client) GetSubmission(subID int) (*Submission, error) {
	var s Submission
	if err := c.http.Get(fmt.Sprintf("/api/submissions/%d", subID), &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) GetJobStatus(jobID int) (string, error) {
	var r JobStatusResponse
	if err := c.http.Get(fmt.Sprintf("/api/jobs/%d", jobID), &r); err != nil {
		return "", err
	}
	return r.Status, nil
}

func (c *Client) GetJobInfo(jobID int) (*JobInfo, error) {
	var info JobInfo
	if err := c.http.Get(fmt.Sprintf("/api/jobs/%d/info/", jobID), &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *Client) GetAnnotations(jobID int) ([]Annotation, error) {
	var r struct {
		Annotations []Annotation `json:"annotations"`
	}
	if err := c.http.Get(fmt.Sprintf("/api/jobs/%d/annotations/", jobID), &r); err != nil {
		return nil, err
	}
	return r.Annotations, nil
}

func (c *Client) AnnotatedImageURL(jobID int) string {
	return fmt.Sprintf("%s/annotated_display/%d", baseURL, jobID)
}
