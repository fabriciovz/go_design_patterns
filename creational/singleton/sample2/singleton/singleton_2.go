package singleton

// singleton implementation for single threaded context

type Singleton interface {
	AddOne()
	GetCount() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) AddOne() {
	s.count++
}

func (s *singleton) GetCount() int {
	return s.count
}
