package sample2

import (
	"testing"
)

func TestObserver1(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		shirtItem := newItem("Nike Shirt")

		observerFirst := &Customer{id: "abc@gmail.com"}
		observerSecond := &Customer{id: "xyz@gmail.com"}

		shirtItem.register(observerFirst)
		shirtItem.register(observerSecond)

		shirtItem.updateAvailability()
	})
}
