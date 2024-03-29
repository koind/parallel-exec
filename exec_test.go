package parallel_exec

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestExecute(t *testing.T) {
	fns := make([]func() error, 0, 5)
	fns = append(fns, func() error {
		fmt.Println("func 1")
		return nil
	}, func() error {
		fmt.Println("func 2")
		return errors.New("Hi this error1")
	}, func() error {
		fmt.Println("func 3")
		return nil
	}, func() error {
		fmt.Println("func 4")
		return nil
	}, func() error {
		fmt.Println("func 5")
		return errors.New("Hi this error")
	})

	Execute(fns, 3, 1)
}
