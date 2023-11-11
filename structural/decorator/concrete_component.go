package decorator

type PlainPizza struct {
}

func (p *PlainPizza) getDescription() string {
	return "Thin Dough"
}

func (p *PlainPizza) getPrice() int {
	return 15
}

func NewPlainPizza() *PlainPizza {
	return &PlainPizza{}
}
