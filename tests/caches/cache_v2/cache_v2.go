package cache_v2

import (
	"github.com/ozonmp/omp-bot/tests/storage"
	"sync"
)

type simpleCache struct {
	m  map[string]string
	mu sync.RWMutex
}

func NewCache() storage.Cache {
	return &simpleCache{m: make(map[string]string)}
}

func (s *simpleCache) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.m[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return value, nil
}

func (s *simpleCache) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
	return nil
}

func (s *simpleCache) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
	return nil
}
