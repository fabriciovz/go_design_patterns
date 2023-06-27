package main

import "log"

type human struct {
	name     string
	lastName string
	age      int
}

// human builder pattern code
type humanBuilder struct {
	human *human
}

func NewhumanBuilder() *humanBuilder {
	human := &human{}
	b := &humanBuilder{human: human}
	return b
}

func (b *humanBuilder) name(name string) *humanBuilder {
	b.human.name = name
	return b
}

func (b *humanBuilder) lastName(lastName string) *humanBuilder {
	b.human.lastName = lastName
	return b
}

func (b *humanBuilder) age(age int) *humanBuilder {
	b.human.age = age
	return b
}

func (b *humanBuilder) Build() *human {
	return b.human
}

func main() {
	me := NewhumanBuilder().name("Fabricio").lastName("Bravo Guevara").age(37).Build()

	log.Printf("%+V", me)
}
