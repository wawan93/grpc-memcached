package inmemory

import (
	"context"
	"fmt"
	"sync"
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var value string

	value, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}

	return value, nil
}

func (s *Storage) Set(_ context.Context, key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	return nil
}

func (s *Storage) Delete(_ context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
	return nil
}
