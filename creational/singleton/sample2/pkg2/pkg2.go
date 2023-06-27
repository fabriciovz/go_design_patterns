package pkg2

import "github.com/fabriciovz/go_patterns/creational/singleton/sample2/singleton"

func GetInstance2() singleton.Singleton {
	instance2 := singleton.GetInstance()

	return instance2
}
