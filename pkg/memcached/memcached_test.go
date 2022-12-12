package memcached_test

import (
	"testing"

	"github.com/wawan93/grpc-memcached/pkg/memcached"
)

func TestNewMemcached(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	_, err := memcached.NewMemcached("localhost:11211")
	if err != nil {
		t.Skip("skipping test, memcached is not running")
	}
	t.Logf("memcached is running")
}

func TestMemcached_Get(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	m, err := memcached.NewMemcached("localhost:11211")
	if err != nil {
		t.Skip("skipping test, memcached is not running")
	}

	err = m.Set("foo", "bar")
	if err != nil {
		t.Errorf("unexpected error %v", err)
		t.Fail()
	}

	value, err := m.Get("foo")
	if err != nil {
		t.Errorf("unexpected error %v", err)
		t.Fail()
	}
	if value != "bar" {
		t.Errorf("expected bar, got %s", value)
		t.Fail()
	}

	err = m.Delete("foo")
	if err != nil {
		t.Errorf("unexpected error %v", err)
		t.Fail()
	}

	_, err = m.Get("foo")
	if err == nil {
		t.Errorf("expected error, got nil")
		t.Fail()
	}
}
