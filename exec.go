package parallel_exec

import (
	"sync"
)

func Execute(funcs []func() error, countParallelExec int, errCount int) {
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	chFuncs := make(chan []func() error, len(funcs))
	couter := 0

	for i := 0; i < countParallelExec; i++ {
		chFuncs <- funcs
	}

	for i := 0; i < countParallelExec; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mutex *sync.RWMutex, chFuncs <-chan []func() error, counter, errCount int) {
			defer wg.Done()
			funcs := <-chFuncs
			for _, fun := range funcs {
				err := fun()
				if err != nil {
					mutex.Lock()
					couter++
					mutex.Unlock()
				}

				mutex.RLock()
				if couter >= errCount {
					mutex.RUnlock()
					return
				}
				mutex.RUnlock()
			}
		}(&wg, &mutex, chFuncs, couter, errCount)
	}

	wg.Wait()
}
