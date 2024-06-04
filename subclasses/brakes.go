package subclasses

import "fmt"

type BrakeInterface interface {
	BrakeCar()
	UnBrakeCar()
}

type Brake struct{}

func (b Brake) BrakeCar() {
	fmt.Println("Braking car")
}

func (b Brake) UnBrakeCar() {
	fmt.Println("Un braking car")
}
