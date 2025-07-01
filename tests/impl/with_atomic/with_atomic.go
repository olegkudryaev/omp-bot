package with_atomic

import (
	"github.com/ozonmp/omp-bot/tests/storage"
	"sync"
	"sync/atomic"
)

// имплементация кеша с mutex и atomic метриками
type impl struct {
	st    map[string]string
	mu    sync.Mutex
	total int64
}

func New() storage.CacheWithMetrics {
	return &impl{st: make(map[string]string),
		mu: sync.Mutex{}}
}

func (i *impl) Set(key, value string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.st[key] = value
	atomic.StoreInt64(&i.total, 1)
	return nil
}

func (i *impl) Get(key string) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, ok := i.st[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return v, nil
}

func (i *impl) Delete(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.st, key)
	atomic.StoreInt64(&i.total, -1)
	return nil
}

func (i *impl) TotalAmount() int64 {
	return atomic.LoadInt64(&i.total)
}
