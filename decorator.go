package main

import (
	"fmt"
	"log"
)

func LogDecorate(fn func(s string)) func(s string) {
	return func(s string) {
		log.Println("Starting the execution with the argument", s)
		fn(s)
		log.Println("Execution is completed.")
	}
}

func Function(s string) {
	fmt.Println(s)
}

func main() {
	f := LogDecorate(Function)

	f("Hello Decorator")
}
