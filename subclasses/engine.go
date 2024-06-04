package subclasses

import "fmt"

type EngineInterface interface {
	StartEngine()
	StopEngine()
	SpeedUpEngine()
	SlowDownEngine()
	TurnOffEngine()
}

type Engine struct {
	isOn         bool
	currentSpeed int
}

func (e *Engine) StartEngine() {
	e.isOn = true
	fmt.Println("Engine is on")
}
func (e *Engine) StopEngine() {
	e.isOn = false
	fmt.Println("Engine is off")
}

func (e *Engine) SpeedUpEngine() {
	if e.isOn {
		e.currentSpeed += 10
		fmt.Println("Car speed is increased by 10")
	} else {
		fmt.Println("Car has not been turned on")
	}
}
func (e *Engine) SlowDownEngine() {
	if !e.isOn {
		e.currentSpeed -= 10
		fmt.Println("Car speed is decreased by 10")
	} else {
		fmt.Println("Car has not been turned on")
	}
}
func (e *Engine) TurnOffEngine() {
	if e.isOn {
		fmt.Println("Car is already turned off")
	} else {
		e.currentSpeed = 0
		e.isOn = false
		e.StopEngine()
	}
}
