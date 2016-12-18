# N-Barrier Pattern
Prevents a process from proceeding until all N processes reach to the barrier.

## Implementation
```go
package barrier

import (
	"errors"
	"reflect"
	"sync"
)

var (
	ErrTypeNotFunction = errors.New("argument type not function")
	ErrInArgsMissMatch = errors.New("input arguments length not match")
)

type Barrier struct {
	n      int
	f      *Handler // registered func executed when all processes come to barrier
	wg     *sync.WaitGroup
	start  *sync.WaitGroup
	finish *sync.WaitGroup
}

func NewBarrier(n int) *Barrier {
	b := new(Barrier)
	b.n = n
	b.wg = new(sync.WaitGroup)
	b.wg.Add(n)
	b.start = new(sync.WaitGroup)
	b.start.Add(n)
	b.finish = new(sync.WaitGroup)
	b.finish.Add(1)
	return b
}

// Register func, to be execute when all processes come to barrier, can be nil
func (b *Barrier) RegisterFunc(f interface{}, args ...interface{}) *Barrier {
	b.f = NewHandler(f, args...)
	return b
}

func (b *Barrier) Wait() {
	b.wg.Done()
	b.wg.Wait()
	if b.f != nil {
		b.start.Done()
		b.finish.Wait()
	}
}

func (b *Barrier) Done() {
	if b.f != nil {
		b.start.Wait()
		b.f.Do()
		b.finish.Done()
	}
}

type Handler struct {
	f    interface{}
	args []interface{}
}

func NewHandler(f interface{}, args ...interface{}) *Handler {
	res := new(Handler)
	res.f = f
	res.args = args
	return res
}

func (h *Handler) Do() {
	f := reflect.ValueOf(h.f)
	typ := f.Type()
	if typ.Kind() != reflect.Func {
		panic(ErrTypeNotFunction)
	}
	// variable parameter, h.args less..
	if typ.NumIn() > len(h.args) {
		panic(ErrInArgsMissMatch)
	}
	inputs := make([]reflect.Value, len(h.args))
	for i := 0; i < len(h.args); i++ {
		if h.args[i] == nil {
			inputs[i] = reflect.Zero(f.Type().In(i))
		} else {
			inputs[i] = reflect.ValueOf(h.args[i])
		}
	}
	f.Call(inputs)
}
```

## Usage
```go
func testBarrier() {
	b := util.NewBarrier(3).RegisterFunc(test1, 1)
	for i := 0; i < 3; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
			log.Printf("%d arrived.", i)
			b.Wait()
			log.Printf("%d finished.", i)
		}(i)
	}
	b.Done() // wait until all processes come to barrier, and run the registered func

	time.Sleep(time.Second)	// wait 1s for other go routine to finish 
}

func test1(i int) {
	time.Sleep(time.Second)
	fmt.Println("exec func: ", i)
}
```
