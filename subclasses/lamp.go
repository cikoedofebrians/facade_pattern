package subclasses

import "fmt"

type LampInterface interface {
	TurnOnCarLamp()
	TurnOffCarLamp()
	TurnOnTurnLamp(isLeft bool)
	TurnOffTurnLamp()
}

type Lamp struct{}

func (l Lamp) TurnOnCarLamp() {
	fmt.Println("Turning on lamp")
}

func (l Lamp) TurnOffCarLamp() {
	fmt.Println("Turning off lamp")
}

func (l Lamp) TurnOnTurnLamp(isLeft bool) {
	if isLeft {
		fmt.Println("Turn on left Lamp")
	} else {
		fmt.Println("Turn on right Lamp")
	}
}

func (l Lamp) TurnOffTurnLamp() {
	fmt.Println("Turning off turn lamp")
}
