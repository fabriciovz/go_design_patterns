package decorator

type MozzarellaTopping struct {
	*ToppingDecorator
}

func (c *MozzarellaTopping) getDescription() string {
	return c.pizza.getDescription() + ", Mozzarella"
}

func (c *MozzarellaTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func NewMozzarellaTopping(pizza IPizza) *MozzarellaTopping {
	newToppingDecorator := NewToppingDecorator(pizza)
	return &MozzarellaTopping{ToppingDecorator: newToppingDecorator}
}
