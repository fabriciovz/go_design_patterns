package main

import "fmt"

//https://upload.wikimedia.org/wikipedia/commons/4/45/W3sDesign_Strategy_Design_Pattern_UML.jpg

// Strategy interface
type SuperHero interface {
	SayHi()
}

// Concrete Strategy 1
type antMan struct{}

func (a *antMan) SayHi() {
	fmt.Println("Hi, I am Ant-man and you??")
}

// Concrete Strategy 2
type theWasp struct{}

func (a *theWasp) SayHi() {
	fmt.Println("Hi, I am The Wasp and you??")
}

// Context
type hero struct {
	hero SuperHero
}

func (h *hero) Talk() {
	h.hero.SayHi()
}

func NewHero(sh SuperHero) *hero {
	return &hero{
		hero: sh,
	}
}
