package subclasses

import "fmt"

type LockInterface interface {
	UnlockCar()
	LockCar()
}

type Lock struct{}

func (lock Lock) LockCar() {
	fmt.Println("Locking car...")
}

func (lock Lock) UnlockCar() {
	fmt.Println("Unlocking car...")
}
