package banch

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/tests/storage"
	"sync"
)

func emulateLoad(c storage.Cache, paralleFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < paralleFactor; i++ {
		//С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("#{i}-key")
		value := fmt.Sprintf("#{i}-value")

		wg.Add(1)
		//Запись в Кеш
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)
		//Чтение из кеша
		go func(k, v string) {
			_, err := c.Get(k)
			//Проверим, что ошбика не связана с тем что записи нет в кеше
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		//Удаление из уша
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
		}(key)
	}
	//Ждем пока все горутины отработают
	wg.Wait()
}

// emulateLoad вспомогательная функция, создает нагрузку на кеш через горутины
func emulateReadIntensiveLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	//Понижаем в 10 раз нагрузку на запись и удаление
	for i := 0; i < parallelFactor/10; i++ {
		//С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("#{i}-key")
		value := fmt.Sprintf("#{i}-value")

		wg.Add(1)
		//Запись в Кеш
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)
		//Удаление из уша
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
		}(key)
	}

	for i := 0; i < parallelFactor/10; i++ {
		//С этими key/value будем работать на этой итерации цикла
		key := fmt.Sprintf("#{i}-key")
		value := fmt.Sprintf("#{i}-value")

		wg.Add(1)
		//Чтение из кеша
		go func(k, v string) {
			_, err := c.Get(k)
			//Проверим, что ошбика не связана с тем что записи нет в кеше
			if err != nil && errors.Is(err, storage.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)
	}
	//Ждем пока все горутины отработают
	wg.Wait()
}

// вспомогательная функция, создает нагрузку на кеш через горутины и проверяет количество записей в кеше
func emulateLoadWithMetrics(cm storage.CacheWithMetrics, parallel int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		emulateLoad(cm, parallelFactor)
		wg.Done()
	}()

	//Добавим забор метрик с кеша
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			_ = cm.TotalAmount()
			wg.Done()
		}()
	}

	wg.Wait()
}
