package main

import (
	"testing"
)

func TestStrategy2(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		laika := NewDog()
		fiu := NewBird()

		laika.animal.tryToFly()
		fiu.animal.tryToFly()

	})
}
