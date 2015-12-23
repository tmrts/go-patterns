package main

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

func main() {
	mult := Operation{Multiplication{}}

	// Outputs 15
	fmt.Println(mult.Operate(3, 5))

	pow := Operation{Addition{}}

	// Outputs 8
	fmt.Println(pow.Operate(3, 5))
}
