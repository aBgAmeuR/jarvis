package pathcheck

import (
	"context"
	"os"

	"jarvis/server/internal/source/failure"
)

type Check struct {
	name string
	path string
}

func New(name, path string) *Check { return &Check{name: name, path: path} }

func (c *Check) Name() string { return c.name }

func (c *Check) Ping(context.Context) error {
	_, err := os.Stat(c.path)
	return failure.Wrap(c.name, err)
}
