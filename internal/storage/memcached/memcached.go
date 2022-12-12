package memcached

import (
	"context"
)

type Storage struct {
	c client
}

type client interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

func New(client client) *Storage {
	return &Storage{
		c: client,
	}
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	return s.c.Get(key)
}

func (s *Storage) Set(_ context.Context, key, value string) error {
	return s.c.Set(key, value)
}

func (s *Storage) Delete(_ context.Context, key string) error {
	return s.c.Delete(key)
}
