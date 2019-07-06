package parallel_exec

import (
	"fmt"
	"sync"
)

func Execute(funcs []func() error, countParallelExec int, errCount int) {
	var wg sync.WaitGroup
	chFuncs := make(chan []func() error, len(funcs))
	chErrors := make(chan []error, errCount)

	for i := 0; i < countParallelExec; i++ {
		chFuncs <- funcs
	}

	for i := 0; i < countParallelExec; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, chFuncs <-chan []func() error, chErrors chan<- []error) {
			defer wg.Done()
			funcs := <-chFuncs
			errors := make([]error, 0)
			for _, fun := range funcs {
				err := fun()
				if err != nil {
					errors = append(errors, err)
				}
			}
			chErrors <- errors
		}(&wg, chFuncs, chErrors)
	}

	i := 0
	for _, err := range <-chErrors {
		fmt.Println(err)
		if err != nil {
			i++
		}

		if errCount >= i {
			close(chFuncs)
			return
		}
	}

	wg.Wait()
}
