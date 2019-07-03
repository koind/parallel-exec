package parallel_exec

import (
	"runtime"
	"sync"
)

func Execute(funcs []func() error, countParallelExec int, errCount int) {
	runtime.GOMAXPROCS(countParallelExec)

	var wg sync.WaitGroup
	errCh := make(chan error)
	stop := make(chan struct{})

	for _, fun := range funcs {
		wg.Add(1)
		go func(fun func() error, ch chan<- error, stop <-chan struct{}) {
			run := true
			for run {
				select {
				case <-stop:
					run = false
					return
				default:
					err := fun()
					if err != nil {
						ch <- err
					}
					wg.Done()
					run = false
				}
			}
		}(fun, errCh, stop)
	}

	go func() {
		i := 0
		for err := range errCh {
			if err != nil {
				i++
			}

			if errCount >= i {
				close(stop)
				return
			}
		}
	}()
	wg.Wait()
}
