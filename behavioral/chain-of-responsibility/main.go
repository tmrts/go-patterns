package main

import (
	"fmt"
)

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

func main() {
	data := &Data{Data: "Example Data"}
	var concreteBHandler Handler = ConcreteBHandler{}
	var concreteAHandler Handler = ConcreteAHandler{Successor: concreteBHandler}
	concreteAHandler.Handle(data)
}
