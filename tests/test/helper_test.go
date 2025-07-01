package test

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/tests/storage"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func emulateLoad(t *testing.T, c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		// С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("#{i}-key")
		value := fmt.Sprintf("#{i}-value")

		wg.Add(1)
		//Запись в кеш
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		wg.Add(1)
		//чтение из кеша
		go func(k, v string) {
			storedValue, err := c.Get(k)
			//Если другая горутина не успела удалить значение из кеша
			//Проверим, что оно совпадает с тем что мы хотели добавить в кеш
			if !errors.Is(err, storage.ErrNotFound) {
				assert.Equal(t, v, storedValue)

			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		//удаление из кеша
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
		}(key)
	}
}

// вспомогательная функция, создает нагрузку на кеш через горутины ипроверяет количество записей в кеше

func emulateLoadWithMetrics(t *testing.T, cm storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		emulateLoad(t, cm, parallelFactor)
	}()

	//добавим набор метрик с кеша
	var min, max int64
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			total := cm.TotalAmount()
			if total > max {
				max = total
			}
			if total < min {
				min = total
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(max, min)
}
