package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		pizza := NewPlainPizza()

		fmt.Printf("\nPizza description: %s\n", pizza.getDescription())
		fmt.Printf("Price: %d\n", pizza.getPrice())

		//Add cheese topping
		pizzaWithCheese := NewMozzarellaTopping(pizza)

		fmt.Printf("\nPizza description: %s\n", pizzaWithCheese.getDescription())
		fmt.Printf("Price: %d\n", pizzaWithCheese.getPrice())

		//Add tomato topping
		pizzaWithCheeseAndTomato := NewTomatoTopping(pizzaWithCheese)

		fmt.Printf("\nPizza description: %s\n", pizzaWithCheeseAndTomato.getDescription())
		fmt.Printf("Price: %d\n", pizzaWithCheeseAndTomato.getPrice())
	})
}
