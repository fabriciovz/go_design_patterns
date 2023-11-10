package main

import "fmt"

//https://upload.wikimedia.org/wikipedia/commons/4/45/W3sDesign_Strategy_Design_Pattern_UML.jpg

// Strategy interface
type Flys interface {
	Fly() string
}

// Concrete Strategy 1
type itFlys struct{}

func (a *itFlys) Fly() string {
	return "I can fly"
}

// Concrete Strategy 2
type cantFly struct{}

func (a *cantFly) Fly() string {
	return "I can't fly"
}

// Context
type animal struct {
	flyingType Flys
}

func (h *animal) tryToFly() {
	fmt.Println(h.flyingType.Fly())
}
func (h *animal) setFlyingAbility(flyingType Flys) {
	h.flyingType = flyingType
}

func newAnimal(flyingType Flys) *animal {
	return &animal{
		flyingType: flyingType,
	}
}

type dog struct {
	animal *animal
}

func NewDog() *dog {
	flyingType := &cantFly{}
	animal := newAnimal(flyingType)
	return &dog{
		animal: animal,
	}
}

type bird struct {
	animal *animal
}

func NewBird() *bird {
	flyingType := &itFlys{}
	animal := newAnimal(flyingType)
	return &bird{
		animal: animal,
	}
}
