# parallel-exec

Parallel execution of functions on Go.

## Installation

Run the following command from you terminal:

```bash
go get github.com/koind/parallel-exec
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

