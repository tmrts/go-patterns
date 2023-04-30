# Future & Promise Pattern

    Future pattern (also called Promise pattern) is a design pattern used in concurrent and asynchronous programming, which allows clients to start a potentially long-running operation and obtain its result in a non-blocking way by returning a placeholder object (a future or a promise).

    When the operation completes, the result is set on the future/promise object, which can then be obtained by the client code through a blocking or non-blocking method call.

### Types

```go
package main

import (
	"fmt"
	"time"
)

type FutureInt struct {
	ch chan int
}
```

### Implementation

```go
func (f *FutureInt) Get() int {
	return <-f.ch
}

func NewFutureInt() *FutureInt {
	return &FutureInt{
		ch: make(chan int),
	}
}

func add(a, b int) *FutureInt {
	f := NewFutureInt()
	go func() {
		time.Sleep(3 * time.Second) // Simulate long running task
		f.ch <- a + b
	}()
	return f
}

func main() {
	fmt.Println("Starting task...")
	f := add(2, 3)
	fmt.Println("Task started...")
	fmt.Printf("Result: %d\n", f.Get())
}
```
