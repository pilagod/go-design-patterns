package composite

import (
	"fmt"
)

// Athlete implements Trainer interface
type Athlete struct{}

// Train trains Athlete
func (a *Athlete) Train() {
	fmt.Println("Training")
}

// AlthleteCompositeSwimmer composes struct Althlete, and owns self Swim method
type AlthleteCompositeSwimmer struct {
	Althlete Athlete
	Swim     func()
}

// Swim makes SwimmerComposite swim
func Swim() {
	fmt.Println("Swimming")
}

// Animal represents an animal that can eat
type Animal struct{}

// Eat makes Animal eat
func (a *Animal) Eat() {
	fmt.Println("Eating")
}

// Shark composes Animal, and owns self Swim method
type Shark struct {
	Animal
	Swim func()
}

// Swimmer is an interface describing an instance that can swim
type Swimmer interface {
	Swim()
}

// Trainer is an interface describing an instance that can be trained
type Trainer interface {
	Train()
}

// SwimmerImpl implements Swimmer interface
type SwimmerImpl struct{}

// Swim makes SwimmerImpl swim
func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

// InterfaceCompositeSwimmer composes Swimmer and Trainer interface
type InterfaceCompositeSwimmer struct {
	Trainer
	Swimmer
}
