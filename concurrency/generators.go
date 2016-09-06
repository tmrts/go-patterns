package main

import (
	"fmt"
)

//implement generator by closure
func FibnacciClosure() func() (ret int) {
	a, b := 0, 1
	return func() (ret int) {
		ret = b
		a, b = b, a+b
		return
	}
}

//implement generator by channel
func FibnacciChan(n int) chan int {
	ret := make(chan int)

	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			ret <- b
			a, b = b, a+b
		}
		close(ret)
	}()

	return ret
}

func main() {
	//closure
	nextFib := FibnacciClosure()
	for i := 0; i < 20; i++ {
		fmt.Println(nextFib())
	}

	//channel
	for i := range FibnacciChan(20) {
		fmt.Println(i)
	}
}
