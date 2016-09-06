package main

import (
	"fmt"
)

//FibonacciClosure implements fibonacci number generatoin using closure
func FibonacciClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//FibonacciChan implements fibonacci number generatoin using channel
func FibonacciChan(n int) chan int {
	c := make(chan int)

	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			c <- b
			a, b = b, a+b
		}
		close(c)
	}()

	return c
}

func main() {
	//closure
	nextFib := FibonacciClosure()
	for i := 0; i < 20; i++ {
		fmt.Println(nextFib())
	}

	//channel
	for i := range FibonacciChan(20) {
		fmt.Println(i)
	}
}
