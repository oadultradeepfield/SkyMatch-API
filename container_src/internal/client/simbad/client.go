package simbad

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseURL = "https://simbad.u-strasbg.fr/simbad/sim-tap/sync"

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{http: &http.Client{Timeout: 10 * time.Second}}
}

func (c *Client) QueryObject(identifier string) (*ObjectInfo, error) {
	query := fmt.Sprintf(`
		SELECT TOP 1 main_id, otype_txt, sp_type, flux AS vmag, plx_value, ra, dec
		FROM basic
		LEFT JOIN flux ON basic.oid = flux.oidref AND flux.filter = 'V'
		WHERE main_id = '%s' OR oid IN (SELECT oidref FROM ident WHERE id = '%s')
	`, escape(identifier), escape(identifier))

	reqURL := baseURL + "?" + url.Values{
		"request": {"doQuery"},
		"lang":    {"adql"},
		"format":  {"json"},
		"query":   {query},
	}.Encode()

	resp, err := c.http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("simbad request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("simbad status: %d", resp.StatusCode)
	}

	var r struct {
		Data [][]any `json:"data"`
	}
	if decodeErr := json.NewDecoder(resp.Body).Decode(&r); decodeErr != nil {
		return nil, fmt.Errorf("decode response: %w", decodeErr)
	}
	if len(r.Data) == 0 {
		return nil, fmt.Errorf("not found: %s", identifier)
	}

	return parseRow(r.Data[0], identifier), nil
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
