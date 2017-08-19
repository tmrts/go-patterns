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

import "strconv"

type Speed float32

const (
    MPH Speed = 1
    KPH       = 1.60934
)

type Color string

const (
    BlueColor  Color = "blue"
    GreenColor       = "green"
    RedColor         = "red"
)

type Wheels string

const (
    SportsWheels Wheels = "sports"
    SteelWheels         = "steel"
)

type Builder interface {
    Color(Color) Builder
    Wheels(Wheels) Builder
    TopSpeed(Speed) Builder
    Build() Interface
}

type Interface interface {
    Drive() string
    Stop() string
}

type carBuilder struct {
	speed  Speed
	color  Color
	wheels Wheels
}

type carObject struct {
	topSpeed Speed
	color    Color
	wheels   Wheels
}

func NewBulder() Builder {
	return &carBuilder{}
}

func (cb *carBuilder) TopSpeed(speed Speed) Builder {
	cb.speed = speed
	return cb
}

func (cb *carBuilder) Paint(color Color) Builder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) Builder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) Build() Interface {
	return &carObject{
		topSpeed: cb.speed,
		color:    cb.color,
		wheels:   cb.wheels,
	}
}

func (c *carObject) Drive() string {
	return "Driving at speed: " + strconv.FormatFloat(float64(c.topSpeed), 'f', 2, 32)
}

func (c *carObject) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}
```

## Usage

```go
assembly := car.NewBuilder().Paint(car.RedColor)

familyCar := assembly.Wheels(car.SteelWheels).TopSpeed(50 * car.MPH).Build()
fmt.Println(familyCar.Drive())

sportsCar := assembly.Paint(car.BlueColor).Wheels(car.SportsWheels).TopSpeed(150 * car.MPH).Build()
fmt.Println(sportsCar.Drive())

fmt.Println(familyCar.Stop())
fmt.Println(sportsCar.Stop()
```
