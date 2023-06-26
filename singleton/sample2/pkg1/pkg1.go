package pkg1

import "github.com/fabriciovz/go_patterns/singleton/sample2/singleton"

func GetInstance1() singleton.Singleton {
	instance1 := singleton.GetInstance()

	return instance1
}
