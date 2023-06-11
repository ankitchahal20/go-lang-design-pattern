package main

import "fmt"

// Finally, there is a third method to use the Composite pattern.
// We could create a Swimmer interface with a Swim method and a SwimmerImpl type to embed it in the athlete swimmer:

type Trainer interface {
	Train()
}

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training1")
}

//

type Swimmer interface {
	Swim()
}
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming1!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	/*
		With this method, you have more explicit control over object creation.
		The Swimmer field is embedded, but won't be zero-initialized as it is a pointer to an interface.
		The correct use of this approach will be the following:
	*/
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}
