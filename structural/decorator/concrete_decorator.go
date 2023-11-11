package decorator

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getDescription() string {
	return c.pizza.getDescription() + ", Tomato"
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func NewTomatoTopping(pizza IPizza) *TomatoTopping {
	return &TomatoTopping{pizza: pizza}
}
