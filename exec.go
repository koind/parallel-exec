package parallel_exec

import (
	"sync"
)

func Execute(funcs []func() error, countParallelExec int, errCount int) {
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	chFunc := make(chan func() error, len(funcs))
	counter := 0

	for _, fun := range funcs {
		chFunc <- fun
	}

	for i := 0; i < countParallelExec; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mutex *sync.RWMutex, chFunc <-chan func() error, counter, errCount int) {
			defer wg.Done()
			fun := <-chFunc
			err := fun()
			if err != nil {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}

			mutex.RLock()
			if counter >= errCount {
				mutex.RUnlock()
				return
			}
			mutex.RUnlock()
		}(&wg, &mutex, chFunc, counter, errCount)
	}

	wg.Wait()
}
