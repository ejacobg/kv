package maps

// Store represents a kv.Store backed by a map.
type Store struct {
	s map[int]int
}

func NewStore() *Store {
	return &Store{
		s: make(map[int]int),
	}
}

func (s *Store) Get(key int) int {
	if value, ok := s.s[key]; ok {
		return value
	}

	return -1
}

func (s *Store) Put(key, value int) {
	s.s[key] = value
}

func (s *Store) Delete(key int) {
	delete(s.s, key)
}
