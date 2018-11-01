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
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Driver
}

type Sedan struct {
	color    Color
	wheels   Wheels
	topspeed Speed
	driver   Driver
}

func (s Sedan) Paint(c Color) Builder     { s.color = Color(c); return s }
func (s Sedan) Wheels(w Wheels) Builder   { s.wheels = w; return s }
func (s Sedan) TopSpeed(sp Speed) Builder { s.topspeed = sp; return s }
func (s Sedan) Build() Driver             { return s.driver }
func NewBuilder(d Driver) Builder         { return Sedan{driver: d} }

type Driver interface {
	Drive() error;
	Stop() error
}

type Terminator struct{}

func (Terminator) Drive() error { fmt.Println("driving, I'll be baack"); return nil }
func (Terminator) Stop() error  { fmt.Println("no problemo, stopping"); return nil }

type Replicant struct{}

func (Replicant) Drive() error { fmt.Println("driving, live in fear"); return nil }
func (Replicant) Stop() error  { fmt.Println("that hurt, stopping"); return nil }
```

## Usage

```go
package main

import "./car"

func main() {
	assembly1 := car.NewBuilder(car.Terminator{}).Paint(car.RedColor)
	familyCar := assembly1.Wheels(car.SteelWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()
	familyCar.Stop()

	assembly2 := car.NewBuilder(car.Replicant{}).Paint(car.BlueColor)
	sportsCar := assembly2.Wheels(car.SportsWheels).TopSpeed(2 * 80 * car.KPH).Build()
	sportsCar.Drive()
	sportsCar.Stop()
}
```

## Output

```
driving, I'll be baack
no problemo, stopping
driving, live in fear
that hurt, stopping
```
