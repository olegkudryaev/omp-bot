package with_mutex

import (
	"github.com/ozonmp/omp-bot/tests/storage"
	"sync"
)

type impl struct {
	st map[string]string
	mu sync.Mutex
}

func New() storage.Cache {
	return &impl{
		st: map[string]string{},
		mu: sync.Mutex{},
	}
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

func (i impl) Set(key, value string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.st[key] = value
	return nil
}

func (i impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.st, key)
	return nil
}
