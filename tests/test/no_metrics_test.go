package test

import (
	"github.com/ozonmp/omp-bot/tests/caches/cache_v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cache(t *testing.T) {
	t.Parallel()
	//Разные элементы кешей
	testCache := cache_v2.NewCache()
	//testCache := no_mutes.New

	t.Run("corrently stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)

		newValue := "someValue"
		err = testCache.Set(key, newValue)
		assert.NoError(t, err)
		newStoredValue, err := testCache.Get(key)
		assert.Equal(t, value, newStoredValue)

	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoad(t, testCache, parallelFactor)
	})

}
