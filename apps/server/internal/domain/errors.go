package domain

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// NotFoundError names the kind of object that is missing, so a caller can say it. It matches ErrNotFound.
type NotFoundError struct {
	Kind string
	ID   string
}

func NewNotFound(kind string, id any) error {
	return &NotFoundError{Kind: kind, ID: fmt.Sprint(id)}
}

func (e *NotFoundError) Error() string { return fmt.Sprintf("%s %s: %s", e.Kind, e.ID, ErrNotFound) }

func (e *NotFoundError) Is(target error) bool { return target == ErrNotFound }

// SourceErrorKind says how a source failed, without the detail that stays in the logs.
type SourceErrorKind string

const (
	SourceTimeout         SourceErrorKind = "timeout"
	SourceUnreachable     SourceErrorKind = "unreachable"
	SourceBadStatus       SourceErrorKind = "bad_status"
	SourceInvalidResponse SourceErrorKind = "invalid_response"
	SourceUnavailable     SourceErrorKind = "unavailable"
)

type SourceError struct {
	Source string
	Kind   SourceErrorKind
	Err    error
}

func (e *SourceError) Error() string { return fmt.Sprintf("%s: %s: %v", e.Source, e.Kind, e.Err) }

func (e *SourceError) Unwrap() error { return e.Err }

// ErrInvalid wraps a refusal whose message is safe to show.
var ErrInvalid = errors.New("invalid request")
