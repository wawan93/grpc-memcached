package memcached

import (
	"context"

	"github.com/wawan93/grpc-memcached/pkg/memcached"
)

type Storage struct {
	connection *memcached.Memcached
}

func New(addr string) (*Storage, error) {
	connection, err := memcached.NewMemcached(addr)
	if err != nil {
		return nil, err
	}

	return &Storage{
		connection: connection,
	}, nil
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	return s.connection.Get(key)
}

func (s *Storage) Set(_ context.Context, key, value string) error {
	return s.connection.Set(key, value)
}

func (s *Storage) Delete(_ context.Context, key string) error {
	return s.connection.Delete(key)
}
