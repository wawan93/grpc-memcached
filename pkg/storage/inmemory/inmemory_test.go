package inmemory_test

import (
	"context"
	"testing"

	"github.com/wawan93/grpc-memcached/pkg/storage/inmemory"
)

func TestStorage_Get(t *testing.T) {
	cases := map[string]struct {
		key      string
		expected string
	}{
		"success": {},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := inmemory.New()
			err := s.Set(context.Background(), tc.key, tc.expected)
			if err != nil {
				t.Fatalf("failed to set key: %v", err)
			}

			got, err := s.Get(nil, tc.key)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}
