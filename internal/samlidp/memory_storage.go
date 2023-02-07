package samlidp

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]string
}

// Get извлекает данные из key и преобразует их в value
func (s *MemoryStore) Get(key string, value interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.data[key]
	if !ok {
		return ErrNotFound
	}
	return json.Unmarshal([]byte(v), value)
}

// Put сохраняет value в key
func (s *MemoryStore) Put(key string, value interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data == nil {
		s.data = map[string]string{}
	}

	buf, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("Put method error: %s", err)
		return err
	}

	s.data[key] = string(buf)
	return nil
}

// Delete удаляет key
func (s *MemoryStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}

/*
Список возвращает все ключи, которые начинаются с `prefix`.
Префикс удаляется из каждого возвращаемого значения. Например, если ключи ["aa", "ab", "cd"], то List("a") выдаст []строку {"a", "b"} */
func (s *MemoryStore) List(prefix string) ([]string, error) {
	rv := []string{}
	for k := range s.data {
		if strings.HasPrefix(k, prefix) {
			rv = append(rv, strings.TrimPrefix(k, prefix))
		}
	}
	return rv, nil
}
