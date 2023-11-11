package decorator

type ToppingDecorator struct {
	pizza IPizza
}

func NewToppingDecorator(pizza IPizza) *ToppingDecorator {
	return &ToppingDecorator{pizza: pizza}
}
