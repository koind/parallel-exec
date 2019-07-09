# parallel-exec

Parallel execution of functions on Go.

## Installation

Run the following command from you terminal:

```bash
go get github.com/koind/parallel-exec
```


## Usage

Package usage example.

```go
package main

import (
	"fmt"
	"errors"
	
	parallel "github.com/koind/parallel-exec"
)

func main() {
    fns := make([]func() error, 0, 3)
    fns = append(fns, func() error {
        fmt.Println("func 1")
        return nil
    }, func() error {
        fmt.Println("func 2")
        return errors.New("Hi this error1")
    }, func() error {
        fmt.Println("func 3")
        return nil
    })
    
    parallel.Execute(fns, 1, 1)
}
```

## Available Methods

The following methods are available:

##### koind/parallel-exec

```go
Execute(funcs []func() error, countParallelExec int, errCount int)
```

## Tests

Run the following command from you terminal:

```
go test -v .
```

