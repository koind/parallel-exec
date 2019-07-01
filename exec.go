package parallel_exec

import (
	"runtime"
	"sync"
)

func Execute(funcs []func() error, countParallelExec int, errCount int) []interface{} {
	runtime.GOMAXPROCS(countParallelExec)
	var wg sync.WaitGroup
	funcErrors := make([]interface{}, 0, errCount)
	errCounter := 0

	for _, fun := range funcs {
		wg.Add(1)
		go func(func() error) {
			defer func() {
				if r := recover(); r != nil {
					funcErrors = append(funcErrors, r)
					errCounter++
				}
			}()
			defer wg.Done()

			fun()
		}(fun)
	}

	wg.Wait()

	return funcErrors
}
