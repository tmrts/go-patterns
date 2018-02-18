package car

import (
	"fmt"
	"log"
	"strings"
)

const Version = 1

type (
	// Color is string representation of Vihicle color
	Color string

	// Wheels is int representation of Vihicle wheels number
	Wheels string

	// Speed is string representation of Vihicle top speed
	// representation in miles per hour
	Speed float64
)

// Used constants definition

// Type WheelsType enum
const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

// Type Color Enum
const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

// Speed constants for easy spead conversation
const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

// Assembly is Assembly line struct that
// implement Builder interface
type Assembly struct {
	v *Car
}

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Vihicle
}

// Stoper for Stop implementation and intefae segregation
type Stoper interface {
	Stop()
}

// Driver for Drive implementation and interface segregation
type Driver interface {
	Drive()
}

// Vihicle interface to clarify what is vihicle:
// - is it drives and if its stops its vihicle
type Vihicle interface {
	Stoper
	Driver
}

// Vihicle interface implementation

// Drive is driving the car
func (c Car) Drive() {
	fmt.Println("Car drives")
}

// Stop stops the car.
func (c Car) Stop() {
	fmt.Println("Car stops")
}

func (c Car) String() string {

	var color string

	if c.color != "" {
		color = fmt.Sprintf("%s Car", strings.Title(string(c.color)))
	} else {
		color = "(Soon to be rusty) Car"
	}

	return fmt.Sprintf("%s on %s wheels with top speed %.2f mph", color, c.wheels, c.speed)
}

// Builder interface implementation

// NewAssembly to create new Assembly line instance
func NewAssembly(vihicle *Car) *Assembly {
	return &Assembly{v: vihicle}
}

// Color building color option of car in assembly line.
func (a *Assembly) Color(color Color) *Assembly {
	a.v.color = color
	return a
}

// Speed building speed option of car in assembly line.
func (a *Assembly) TopSpeed(speed Speed) *Assembly {
	if speed <= 0.0 {
		log.Fatal("Final Product wouldn't be able to implement Drive interface")
	}

	if speed > 276.0 {
		log.Fatal("Mister Musk, stop doing that, Tesla 3 fastest in the space, but not in on the earth")
	}

	a.v.speed = speed
	return a
}

// Wheels building wheels option of car in assembly line.
func (a *Assembly) Wheels(wheels Wheels) *Assembly {
	a.v.wheels = wheels
	return a
}

// Build interface Implementation
func (a Assembly) Build() Vihicle {
	return a.v
}
