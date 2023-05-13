package slices

import "golang.org/x/exp/slices"

// Store represents a store.Store backed by a slice.
type Store struct {
	s []item
}

type item struct {
	key, value int
}

func NewStore(cap int) *Store {
	return &Store{
		s: make([]item, cap),
	}
}

func (s *Store) Get(key int) int {
	index := s.find(key)

	if index == -1 {
		return -1
	}

	return s.s[index].value
}

func (s *Store) Put(key, value int) {
	index := s.find(key)

	if index == -1 {
		s.s = append(s.s, item{key, value})
	} else {
		s.s[index].value = value
	}
}

func (s *Store) Delete(key int) {
	index := s.find(key)

	// No-op if key is not present.
	if index == -1 {
		return
	}

	s.s = slices.Delete(s.s, index, index+1)
}

// find returns the index of the element with the given key, or -1 if it is not present.
func (s *Store) find(key int) int {
	return slices.IndexFunc(s.s, func(itm item) bool {
		return itm.key == key
	})
}
