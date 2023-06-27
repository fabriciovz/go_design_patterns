package main

import "log"

type human struct {
	name     string
	lastName string
	age      int
}

// Builders
func (h *human) withName(name string) *human {
	h.name = name
	return h
}

func (h *human) withLastName(lastName string) *human {
	h.lastName = lastName
	return h
}

func (h *human) withAge(age int) *human {
	h.age = age
	return h
}

// Empty constructor
func NewHuman() *human {
	return &human{}
}

func main() {
	me := NewHuman().
		withName("Fabricio").
		withLastName("Bravo Guevara").
		withAge(37)

	log.Printf("%+V", me)
}
