package bad_singleton_impl

// a bad singleton implementation for single threaded context
// creational design pattern

type HumanService struct {
	human *human
}

func (s *HumanService) GetSingleton() *human {
	if s.human == nil {
		s.human = newHuman()
	}
	return s.human
}

type human struct {
	age int
}

func newHuman() *human {
	return &human{age: 0}
}

func (h *human) Birthday() {
	h.age++
}

func (h *human) GetAge() int {
	return h.age
}
