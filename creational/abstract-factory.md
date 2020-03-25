# Abstract Factory Method Pattern

Abstract Factory is a creational design pattern that allows you to create a _collection of related objects_. 
It provides an extra level of indirection over the [factory design pattern](factory.md).
Using this pattern you can decouple the factory from your client code.

## Implementation

The example implementation shows how to create a collection (top & bottom) of summer clothes or winter clothes.

### Types

We begin with the simplest interfaces to represent a top or bottom article of clothing. Our items of clothing are [Stringer](https://tour.golang.org/methods/17) types.

```go
package main

import "fmt"

type Top interface {
	String() string
}

type Bottom interface {
	String() string
}
```

Next we introduce a couple of tops and bottoms into our system.

```go
type TankTop struct {
}

func (top TankTop) String() string {
	return "Tank Top"
}

type Shorts struct {
}

func (bottom Shorts) String() string {
	return "Shorts"
}

type Sweater struct {
}

func (top Sweater) String() string {
	return "Sweater"
}

type WoolenPant struct {
}

func (bottom WoolenPant) String() string {
	return "Woolen Pant"
}

```

Next we create an interface to manage the _relatedness_ of clothes (summery or wintery)

```go
type ClothesFactory interface {
	GetTop() Top
	GetBottom() Bottom
}
```

We give concrete implementations to this interface and bring it all together.

```go
type SummerClothesFactory struct {
}

func (summer SummerClothesFactory) GetTop() Top {
	return TankTop{}
}

func (summer SummerClothesFactory) GetBottom() Bottom {
	return Shorts{}
}

type WinterClothesFactory struct {
}

func (winter WinterClothesFactory) GetTop() Top {
	return Sweater{}
}

func (winter WinterClothesFactory) GetBottom() Bottom {
	return WoolenPant{}
}
```

## Usage

```go
type Wardrobe struct {
	clothesFactory ClothesFactory
}

func NewWardrobe(factory ClothesFactory) *Wardrobe {
	return &Wardrobe{
		clothesFactory: factory,
	}
}

func (wardrobe Wardrobe) GetTop() Top {
	return wardrobe.clothesFactory.GetTop()
}

func (wardrobe Wardrobe) GetBottom() Bottom {
	return wardrobe.clothesFactory.GetBottom()
}


func main() {
	summerWardrobe := NewWardrobe(SummerClothesFactory{})
	fmt.Printf("Summer clothes are %s & %s", summerWardrobe.GetTop(), summerWardrobe.GetBottom())

	fmt.Println()
	winterWardrobe := NewWardrobe(WinterClothesFactory{})
	fmt.Printf("Winter clothes are %s & %s", winterWardrobe.GetTop(), winterWardrobe.GetBottom())

}
```
