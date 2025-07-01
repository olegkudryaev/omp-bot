package with_int

import (
	"github.com/ozonmp/omp-bot/tests/storage"
	"sync"
)

type impl struct {
	st    map[string]string
	mu    sync.Mutex
	total int64
}

func New() storage.CacheWithMetrics {
	return &impl{
		st: make(map[string]string),
		mu: sync.Mutex{},
	}
}

func (i impl) Set(key, value string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.st[key] = value
	i.total++
	return nil
}

func (i impl) Get(key string) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, ok := i.st[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return v, nil
}

func (i impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.st, key)
	i.total--
	return nil
}

func (i impl) TotalAmount() int64 {
	return i.total
}
