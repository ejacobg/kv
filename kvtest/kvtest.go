package kvtest

import (
	"github.com/ejacobg/kv"
	"testing"
)

// Based off of https://google.github.io/styleguide/go/best-practices#writing-an-acceptance-test

func TestStore(t *testing.T, s kv.Store) {
	var key, value, got, want int

	// Can read what was inserted.
	key, value = 1, 1
	s.Put(key, value)

	got, want = s.Get(key), 1
	if got != want {
		t.Errorf("inserted key %d returns %d, want %d", key, got, want)
	}

	// Can read what was updated.
	key, value = 1, 2
	s.Put(key, value)

	got, want = s.Get(key), 2
	if got != want {
		t.Errorf("updated key %d returns %d, want %d", key, got, want)
	}

	// Cannot read after a delete operation.
	key = 1
	s.Delete(key)

	got, want = s.Get(key), -1
	if got != want {
		t.Errorf("deleted key %d returns %d, want %d", key, got, want)
	}
}
