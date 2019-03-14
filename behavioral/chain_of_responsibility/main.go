// Package main serves as an example application that makes use of the chain of responsibility pattern.
package main

import (
	"fmt"
)

type request struct {
	value float64
}

//NumberProcessor has the processor method which will tell if the value is negative, positive or zero
type NumberProcessor interface {
	process(request)
}

//ZeroHandler handles the value which is zero
type ZeroHandler struct {
	numberProcessor NumberProcessor
}

//PositiveHandler handles the value which is positive
type PositiveHandler struct {
	numberProcessor NumberProcessor
}

//NegativeHandler handles the value which is negative
type NegativeHandler struct {
	numberProcessor NumberProcessor
}

//For returning zero if the value is zero.
func (zero ZeroHandler) process(req request) {

	if req.value == 0.0 {
		fmt.Print("its zero")
	} else {
		zero.numberProcessor.process(req)
	}
}

//For returning its negative if the value is negative.
func (negative NegativeHandler) process(req request) {
	if req.value < 0 {
		fmt.Print("its a negative number")
	} else {
		negative.numberProcessor.process(req)
	}
}

//For returning its positive if the value is positive.
func (positve PositiveHandler) process(req request) {
	if req.value > 0 {
		fmt.Print("its a postitive number")
	}
}

func main() {
	//initialising the chain of actions.
	zeroHandler := ZeroHandler{NegativeHandler{PositiveHandler{}}}

	zeroHandler.process(request{10})
	zeroHandler.process(request{-19.9})
	zeroHandler.process(request{0})
}
