package sample1

import "fmt"

// //// Concrete Subject/Publisher
type stockGrabber struct {
	observers []Observer
	ibmPrice  float64
	aaplPrice float64
	googPrice float64
}

func (s *stockGrabber) Register(newObserver Observer) {
	s.observers = append(s.observers, newObserver)
}

func (s *stockGrabber) Unregister(deleteObserver Observer) {
	obsIndex := 0
	for i := range s.observers {
		if deleteObserver == s.observers[i] {
			obsIndex = i
		}
	}
	fmt.Println(fmt.Sprintf("\nObserver: %d deleted", obsIndex+1))
	s.observers = append(s.observers[:obsIndex], s.observers[obsIndex+1:]...)
}

func (s *stockGrabber) NotifyAll() {
	for i := range s.observers {
		s.observers[i].Update(s.ibmPrice, s.aaplPrice, s.googPrice)
	}
}

func (s *stockGrabber) SetIBMPrice(ibmPrice float64) {
	s.ibmPrice = ibmPrice
	s.NotifyAll()
}

func (s *stockGrabber) SetGOOGPrice(googPrice float64) {
	s.googPrice = googPrice
	s.NotifyAll()
}

func (s *stockGrabber) SetAAPLPrice(aaplPrice float64) {
	s.aaplPrice = aaplPrice
	s.NotifyAll()
}

func NewStockGrabber() *stockGrabber {
	return &stockGrabber{}
}
