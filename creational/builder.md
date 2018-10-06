# Builder Pattern

Builder pattern separates the construction of a complex object from its
representation so that the same construction process can create different
representations.

In Go, normally a configuration struct is used to achieve the same behavior,
however passing a struct to the builder method fills the code with boilerplate
`if cfg.Field != nil {...}` checks.

## Implementation

```go
package car

import (
	"fmt"
)

type Speed float64
type Color string
type Wheels string

const (
	MPH          Speed  = 1
	KPH          Speed  = 1.60934
	SportsWheels Wheels = "sports"
	SteelWheels  Wheels = "steel"
	BlueColor    Color  = "blue"
	GreenColor   Color  = "green"
	RedColor     Color  = "red"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type carBuilder struct {
	speedOption Speed
	color       Color
	wheels      Wheels
}

func (c *car) Drive() error {
	fmt.Println("VROOM")
	return nil
}

func (c *car) Stop() error {
	return nil
}

func (cb *carBuilder) Color(color Color) Builder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) Builder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) TopSpeed(speed Speed) Builder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Build() Interface {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
	}
}

func NewBuilder() Builder {
	return &carBuilder{}
}

type car struct {
	topSpeed Speed
	color    Color
}
```

## Usage

```go
assembly := car.NewBuilder().Paint(car.RedColor)

familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
familyCar.Drive()

sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
sportsCar.Drive()
```
