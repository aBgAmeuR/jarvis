package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type snapshot struct {
	status int
	header http.Header
	body   []byte
}

func readSnapshot(resp *http.Response, maxBody int64) (*snapshot, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBody+1))
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	if int64(len(body)) > maxBody {
		return nil, &DecodeError{Err: ErrBodyTooLarge}
	}
	return &snapshot{status: resp.StatusCode, header: resp.Header, body: body}, nil
}

func (s *snapshot) response(req *http.Request) *http.Response {
	return &http.Response{
		Status:        fmt.Sprintf("%d %s", s.status, http.StatusText(s.status)),
		StatusCode:    s.status,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        s.header.Clone(),
		Body:          io.NopCloser(bytes.NewReader(s.body)),
		ContentLength: int64(len(s.body)),
		Request:       req,
	}
}

func (s *snapshot) storable(limit int) bool {
	return s.status >= 200 && s.status < 300 && len(s.body) <= limit
}
