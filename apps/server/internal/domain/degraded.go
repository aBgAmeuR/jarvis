package domain

import "errors"

type Degraded struct {
	Source string
	Kind   SourceErrorKind
}

func NewDegraded(source string, err error) Degraded {
	var sourceErr *SourceError
	if errors.As(err, &sourceErr) {
		return Degraded{Source: sourceErr.Source, Kind: sourceErr.Kind}
	}
	return Degraded{Source: source, Kind: SourceUnavailable}
}
