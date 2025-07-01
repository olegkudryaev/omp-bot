package banch

import (
	"github.com/ozonmp/omp-bot/tests/impl/with_atomic"
	"github.com/ozonmp/omp-bot/tests/impl/with_int"
	"testing"
)

func Benchmark_MutexWithMetricsInt(b *testing.B) {
	c := with_int.New()
	for i := 0; i < b.N; i++ {
		emulateLoadWithMetrics(c, parallelFactor)
	}
}

func Benchmark_MutexWithMetricsAtomic(b *testing.B) {
	c := with_atomic.New()
	for i := 0; i < b.N; i++ {
		emulateLoadWithMetrics(c, parallelFactor)
	}
}
