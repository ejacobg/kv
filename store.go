package kv

// Store stores positive integers.
type Store interface {
	Get(key int) int    // Returns -1 if the key is not present.
	Put(key, value int) // Inserts a new item, updating it if present.
	Delete(key int)     // No-op if key is not present.
}
