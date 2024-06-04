package main

import carfacade "facade_pattern/facade"

func main() {
	carFacade := carfacade.NewCarFacade()
	carFacade.PrepareCar()
	carFacade.TurnRight()
	carFacade.LeaveCar()
}
