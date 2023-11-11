package decorator

type MozzarellaTopping struct {
	pizza IPizza
}

func (c *MozzarellaTopping) getDescription() string {
	return c.pizza.getDescription() + ", Mozzarella"
}

func (c *MozzarellaTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func NewMozzarellaTopping(pizza IPizza) *MozzarellaTopping {
	return &MozzarellaTopping{pizza: pizza}
}
