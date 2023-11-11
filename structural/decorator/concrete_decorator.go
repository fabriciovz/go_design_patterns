package decorator

type TomatoTopping struct {
	*ToppingDecorator
}

func (c *TomatoTopping) getDescription() string {
	return c.pizza.getDescription() + ", Tomato"
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func NewTomatoTopping(pizza IPizza) *TomatoTopping {
	newToppingDecorator := NewToppingDecorator(pizza)
	return &TomatoTopping{ToppingDecorator: newToppingDecorator}
}
