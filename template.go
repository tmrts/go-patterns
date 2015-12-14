// In Template pattern, an abstract struct exposes defined way(s)/template(s)
// to execute its methods. This pattern comes under behavior pattern category.
package main

import (
	"fmt"
	"math"
)

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

func (_ Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

type Pow struct{}

func (_ Pow) Apply(lval, rval int) int {
	return int(math.Pow(float64(lval), float64(rval)))
}

func main() {
	mult := Operation{
		Operator: Multiplication{},
	}

	// Outputs 15
	fmt.Println(mult.Operate(3, 5))

	pow := Operation{
		Operator: Pow{},
	}

	// Outputs 243
	fmt.Println(pow.Operate(3, 5))
}
