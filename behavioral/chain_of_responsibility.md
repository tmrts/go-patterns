# Chain Of Responsibility Pattern


Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it. [Source](http://www.dofactory.com/net/chain-of-responsibility-design-pattern)


## Implementation

```go
type (
	//Data will be processed by chain
	Data struct {
		Data string
	}

	//Handler defines an interface for handling requests
	Handler interface {
		Handle(data *Data)
	}
)


//ConcreteAHandler implements Handler interface
type ConcreteAHandler struct {
	//Next element in the chain
	Successor Handler
}

func (ca ConcreteAHandler) Handle(data *Data) {
	fmt.Printf("%s handled by Concrete A\n", data.Data)
	if ca.Successor != nil {
		ca.Successor.Handle(data)
	}
}

//ConcreteBHandler implements Handler interface
type ConcreteBHandler struct {
	Successor Handler
}

func (cb ConcreteBHandler) Handle(data *Data) {
	fmt.Printf("%s handled by Concrete B\n", data.Data)
	if cb.Successor != nil {
		cb.Successor.Handle(data)
	}
}

```
We have created an interface for handling request. Then we have created  2 concretes struct implementing Handler. Each handler writes own message using by data.


```go
func main() {
	data := &Data{Data: "Example Data"}
	var concreteBHandler Handler = ConcreteBHandler{}
	var concreteAHandler Handler = ConcreteAHandler{Successor: concreteBHandler}
	concreteAHandler.Handle(data)
}
```
Client initializes chain. Objects in the chain process data in order.

## Usage

For usage, see [chain-of-responsibility/main.go](chain-of-responsibility/main.go) or [view in the Playground](https://play.golang.org/p/tN8LZvjNpP).
