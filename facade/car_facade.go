package car_facade

import (
	"facade_pattern/subclasses"
	"fmt"
)

type CarFacade struct {
	brake   subclasses.BrakeInterface
	cleaner subclasses.CleanerInterface
	engine  subclasses.EngineInterface
	lamp    subclasses.LampInterface
	lock    subclasses.LockInterface
}

func (n CarFacade) PrepareCar() {
	fmt.Println("---PREPARING CAR---")
	n.lock.UnlockCar()
	n.engine.StartEngine()
	n.lamp.TurnOnCarLamp()
	n.cleaner.CleanCarWindows()
	fmt.Println("---PREPARING CAR COMPLETE---")
}

func (n CarFacade) LeaveCar() {
	fmt.Println("---LEAVING CAR---")
	n.brake.BrakeCar()
	n.engine.TurnOffEngine()
	n.lamp.TurnOffCarLamp()
	n.lock.LockCar()
	fmt.Println("---LEAVING CAR COMPLETE---")
	fmt.Println("")
}

func (n CarFacade) TurnLeft() {
	fmt.Println("---TURNING CAR LEFT---")
	n.lamp.TurnOnTurnLamp(true)
	n.brake.BrakeCar()
	n.engine.SlowDownEngine()
	n.brake.UnBrakeCar()
	n.engine.SpeedUpEngine()
	n.lamp.TurnOffTurnLamp()
	fmt.Println("---TURNING CAR LEFT COMPLETE---")
	fmt.Println("")
}
func (n CarFacade) TurnRight() {
	n.lamp.TurnOnTurnLamp(false)
	n.brake.BrakeCar()
	n.engine.SlowDownEngine()
	n.brake.UnBrakeCar()
	n.engine.SpeedUpEngine()
	n.lamp.TurnOffTurnLamp()
	fmt.Println("---TURNING CAR RIGHT COMPLETE---")
	fmt.Println("")
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		brake:   subclasses.Brake{},
		cleaner: subclasses.Cleaner{},
		engine:  &subclasses.Engine{},
		lamp:    subclasses.Lamp{},
		lock:    subclasses.Lock{},
	}
}
