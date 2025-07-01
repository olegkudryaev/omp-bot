package cache_v1

import (
	"github.com/ozonmp/omp-bot/tests/storage"
)

type simpleCache struct {
	storage map[string]string
}

func NewCache() storage.Cache {
	return &simpleCache{make(map[string]string)}
}

func (s simpleCache) Set(key, value string) error {
	s.storage[key] = value
	return nil
}

func (s simpleCache) Get(key string) (string, error) {
	vakue, ok := s.storage[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return vakue, nil
}

func (s simpleCache) Delete(key string) error {
	delete(s.storage, key)
	return nil
}
