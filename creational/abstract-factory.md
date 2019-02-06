# Abstract Factory Pattern

The abstract factory pattern provides a way to encapsulate a group of individual factories 
that have a common theme without specifying their concrete classes.

## Implementation


```go
package main

import (
	"fmt"
)

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type ConcreteFactory1 struct {
}

func NewConcreteFactory1() *ConcreteFactory1 {
	return &ConcreteFactory1{}
}

func (ConcreteFactory1) CreateProductA() AbstractProductA {
	return NewConcreteProductA1()
}

func (ConcreteFactory1) CreateProductB() AbstractProductB {
	return NewConcreteProductB1()
}

type ConcreteFactory2 struct {
}

func NewConcreteFactory2() *ConcreteFactory2 {
	return &ConcreteFactory2{}
}

func (ConcreteFactory2) CreateProductA() AbstractProductA {
	return NewConcreteProductA2()
}

func (ConcreteFactory2) CreateProductB() AbstractProductB {
	return NewConcreteProductB2()
}

type AbstractProductA interface {
	UsefulFunctionA() string
}

type AbstractProductB interface {
	UsefulFunctionB() string
	AnotherUsefulFunctionB(collaborator AbstractProductA) string
}

type ConcreteProductA1 struct {
}

func NewConcreteProductA1() *ConcreteProductA1 {
	return &ConcreteProductA1{}
}

func (ConcreteProductA1) UsefulFunctionA() string {
	return "The result of the product A2."
}

type ConcreteProductA2 struct {
}

func NewConcreteProductA2() *ConcreteProductA2 {
	return &ConcreteProductA2{}
}

func (ConcreteProductA2) UsefulFunctionA() string {
	return "The result of the product A2."
}

type ConcreteProductB1 struct {
}

func NewConcreteProductB1() *ConcreteProductB1 {
	return &ConcreteProductB1{}
}

func (ConcreteProductB1) UsefulFunctionB() string {
	return "The result of the product B1."
}

func (ConcreteProductB1) AnotherUsefulFunctionB(collaborator AbstractProductA) string {
	return "The result of the product B2."
}

type ConcreteProductB2 struct {
}

func NewConcreteProductB2() *ConcreteProductB2 {
	return &ConcreteProductB2{}
}

func (ConcreteProductB2) UsefulFunctionB() string {
	return "The result of the product B2."
}

func (ConcreteProductB2) AnotherUsefulFunctionB(collaborator AbstractProductA) string {
	result := collaborator.UsefulFunctionA()
	return fmt.Sprintf("The result of the B2 collaborating with the (%s)", result)
}
```

## Usage

```go

func clientCode(factory AbstractFactory) {
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()
	fmt.Println(productB.UsefulFunctionB())
	fmt.Println(productB.AnotherUsefulFunctionB(productA))
}

func main() {
	fmt.Println("Client: Testing client code with the first factory type:")
	clientCode(NewConcreteFactory1())

	fmt.Println()

	fmt.Println("Client: Testing the same client code with the second factory type:")
	clientCode(NewConcreteFactory2())
}
```
