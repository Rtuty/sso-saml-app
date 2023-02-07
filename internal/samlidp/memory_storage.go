package samlidp

import (
	"encoding/json"
	"sync"
)

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func (s *MemoryStore) Get(key string, value interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.data[key]
	if !ok {
		return ErrNotFound
	}
	return json.Unmarshal([]byte(v), value)
}

func (s *MemoryStore) Put(key string, value interface{}) error { return nil }

func (s *MemoryStore) Delete(key string) error { return nil }

func (s *MemoryStore) List(prefix string) ([]string, error) { return []string{""}, nil }