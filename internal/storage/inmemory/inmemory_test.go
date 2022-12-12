package inmemory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/wawan93/grpc-memcached/internal/storage/inmemory"
	myErrors "github.com/wawan93/grpc-memcached/pkg/errors"
)

func TestStorage_Get(t *testing.T) {
	cases := map[string]struct {
		key           string
		expected      string
		expectedError error
		setup         func(*inmemory.Storage)
	}{
		"success": {
			key:      "foo",
			expected: "bar",
			setup: func(s *inmemory.Storage) {
				_ = s.Set(context.Background(), "foo", "bar")
			},
		},
		"not found": {
			key:           "not exist",
			expectedError: myErrors.ErrNotFound,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := inmemory.New()
			if tc.setup != nil {
				tc.setup(s)
			}

			got, err := s.Get(context.Background(), tc.key)

			if tc.expectedError == nil && err != nil {
				t.Errorf("unexpected error %v", err)
			}

			if tc.expectedError != nil && !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}

			if got != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}
