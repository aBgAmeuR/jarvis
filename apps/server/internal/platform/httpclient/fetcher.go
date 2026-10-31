package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Fetcher struct {
	client        *http.Client
	baseURL       string
	authorization string
	timeout       time.Duration
	maxBody       int64
}

func NewFetcher(client *http.Client, baseURL, authorization string, timeout time.Duration, maxBody int64) *Fetcher {
	return &Fetcher{
		client:        client,
		baseURL:       strings.TrimRight(baseURL, "/"),
		authorization: authorization,
		timeout:       timeout,
		maxBody:       maxBody,
	}
}

func (f *Fetcher) GetJSON(ctx context.Context, path string, out any) error {
	body, err := f.get(ctx, path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, out); err != nil {
		return &DecodeError{Err: err}
	}
	return nil
}

func (f *Fetcher) Ping(ctx context.Context, path string) error {
	_, err := f.get(ctx, path)
	return err
}

func (f *Fetcher) get(ctx context.Context, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, f.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	if f.authorization != "" {
		req.Header.Set("Authorization", f.authorization)
	}
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, &StatusError{Code: resp.StatusCode}
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, f.maxBody+1))
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	if int64(len(body)) > f.maxBody {
		return nil, &DecodeError{Err: ErrBodyTooLarge}
	}
	return body, nil
}
