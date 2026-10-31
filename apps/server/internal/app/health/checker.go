package health

import "context"

type Checker interface {
	Name() string
	Ping(ctx context.Context) error
}
