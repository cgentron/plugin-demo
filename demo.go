package plugindemo

import (
	"context"

	"github.com/cgentron/api/resolvers"
)

// Config ...
type Config struct{}

// CreateConfig ...
func CreateConfig() *Config {
	return &Config{}
}

// Demo ...
type Demo struct {
	name string
}

// New ...
func New(config *Config, name string) (resolvers.ResolverHandler, error) {
	return &Demo{
		name: name,
	}, nil
}

func (a *Demo) Resolve(context context.Context, msg []byte) ([]byte, error) {
	return []byte{}, nil
}
