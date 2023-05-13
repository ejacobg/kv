# Acceptance Tests

This project defines a [key-value store](/store.go) with two different implementations, one using a map, and another using a slice. A [simple test](/kvtest/kvtest.go) was written to confirm the functionality of the store.

```go
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
```

If you want to test a new concrete implementation of a store, then you should add this test to your package:

```go
package mystore

import (
	"github.com/ejacobg/kv"
	"github.com/ejacobg/kv/kvtest"
	"testing"
)

func TestAcceptance(t *testing.T) {
	var s kv.Store

	// Any setup code...

	kvtest.TestStore(t, s)
}
```

This way, you can confirm that your implementation satisfies the base functionality for a store, without having to rewrite the same test every time.
