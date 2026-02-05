package simbad

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"server/internal/config"
	"server/internal/util/httputil"
	"server/internal/util/ratelimit"
)

type Client struct {
	http    *httputil.Client
	limiter *ratelimit.Limiter
}

func NewClient(cfg config.SimbadConfig) *Client {
	return &Client{
		http:    httputil.NewClient(cfg.BaseURL, cfg.Timeout),
		limiter: ratelimit.New(5, 10),
	}
}

func (c *Client) QueryObject(ctx context.Context, identifier string) (*ObjectInfo, error) {
	if err := c.limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limit: %w", err)
	}

	query := fmt.Sprintf(`
		SELECT TOP 1 main_id, otype_txt, sp_type, flux AS vmag, plx_value, ra, dec
		FROM basic
		LEFT JOIN flux ON basic.oid = flux.oidref AND flux.filter = 'V'
		WHERE main_id = '%s' OR oid IN (SELECT oidref FROM ident WHERE id = '%s')
	`, escape(identifier), escape(identifier))

	params := url.Values{
		"request": {"doQuery"},
		"lang":    {"adql"},
		"format":  {"json"},
		"query":   {query},
	}

	var r tapResponse
	if err := c.http.GetWithParams(ctx, "", params, &r); err != nil {
		return nil, fmt.Errorf("simbad request: %w", err)
	}
	if len(r.Data) == 0 {
		return nil, fmt.Errorf("not found: %s", identifier)
	}

	return parseRow(r.Data[0], identifier), nil
}

type tapResponse struct {
	Data [][]any `json:"data"`
}

func parseRow(row []any, identifier string) *ObjectInfo {
	info := &ObjectInfo{Identifier: identifier}

	getString := func(idx int) string {
		if idx < len(row) {
			if v, ok := row[idx].(string); ok {
				return v
			}
		}
		return ""
	}

	getFloat := func(idx int) *float64 {
		if idx < len(row) {
			if v, ok := row[idx].(float64); ok {
				return &v
			}
		}
		return nil
	}

	if s := getString(0); s != "" {
		info.Identifier = s
	}
	info.ObjectType = getString(1)
	info.SpectralType = getString(2)
	info.VMagnitude = getFloat(3)
	info.Parallax = getFloat(4)
	info.RA = getFloat(5)
	info.Dec = getFloat(6)

	return info
}

func escape(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
