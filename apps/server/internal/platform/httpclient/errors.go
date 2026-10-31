package httpclient

import (
	"errors"
	"fmt"
)

var ErrBodyTooLarge = errors.New("response body too large")

type StatusError struct {
	Code int
}

func (e *StatusError) Error() string { return fmt.Sprintf("unexpected status %d", e.Code) }

type DecodeError struct {
	Err error
}

func (e *DecodeError) Error() string { return fmt.Sprintf("decode response: %v", e.Err) }

func (e *DecodeError) Unwrap() error { return e.Err }
