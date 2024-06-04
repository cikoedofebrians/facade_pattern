package subclasses

import "fmt"

type CleanerInterface interface {
	CleanCarWindows()
}

type Cleaner struct {
}

func (c Cleaner) CleanCarWindows() {
	fmt.Println("Cleaning car windows...")
}
