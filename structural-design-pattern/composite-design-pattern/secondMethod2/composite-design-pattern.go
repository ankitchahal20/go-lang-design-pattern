package main

import "fmt"

/*

We must have an Athlete struct with a Train method
We must have a Swimmer with a Swim method
We must have an Animal struct with an Eat method
We must have a Fish struct with a Swim method that is shared with the Swimmer, and not have inheritance or hierarchy issues

*/

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

// We'll create a composite swimmer that has an Athlete struct inside it.

type CompositeSwimmerA struct {
	MyAthlete Athlete //the direct composition
	MySwim    func()  // in Go, functions are first-class citizens and they can be used as parameters, fields, or arguments just like any variable
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal //the embedding composition.
	Swim   func()
}

func main() {
	// With this method, you have no control over object creation.
	fish := Shark{
		Swim: Swim,
	}

	// which is zero-initialized and embedded
	fish.Eat()
	fish.Swim()
}
