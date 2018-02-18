package main

import (
	"fmt"

	"./builder"
)

func main() {

	assembly := car.NewAssembly(&car.Car{}).Color(car.RedColor)

	familyCar := assembly.Wheels(car.SteelWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	fmt.Printf("%#v\n", familyCar)
	fmt.Printf("%#s\n", familyCar)

	sportsCar := assembly.Wheels(car.SportsWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()

	fmt.Printf("%#v\n", sportsCar)
	fmt.Printf("%#s\n", sportsCar)

}
