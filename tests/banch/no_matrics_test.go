package banch

import (
	"github.com/ozonmp/omp-bot/tests/caches/cache_v1"
	"github.com/ozonmp/omp-bot/tests/caches/cache_v2"
	"testing"
)

const parallelFactor = 10_000

func Benchmark_NoMutes(b *testing.B) {
	b.Skip("panic in NoMutex")
	c := cache_v1.NewCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_Mutex_BalancedLoad(b *testing.B) {
	b.Skip()
	c := with_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutex_BalancedLoad(b *testing.B) {
	b.Skip()
	//структура c rw mutex
	c := cache_v2.NewCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

func Benchmark_Mutex_ReadIntensiveLoad(b *testing.B) {
	b.Skip()
	c := with_mutex.New()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(c, parallelFactor)
	}
}
