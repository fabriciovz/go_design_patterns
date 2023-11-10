package observer

import (
	"fmt"
)

type Subject interface {
	Register(o Observer)
	Unregister(o Observer)
	NotifyObserver()
}

type Observer interface {
	Update(ibmPrice, aaplPrice, googPrice float64)
}

// ////
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
	fmt.Println(fmt.Sprintf("\nObserver: %d deleted", +(obsIndex + 1)))
	s.observers = append(s.observers[:obsIndex], s.observers[obsIndex+1:]...)
}

func (s *stockGrabber) NotifyObserver() {
	for i := range s.observers {
		s.observers[i].Update(s.ibmPrice, s.aaplPrice, s.googPrice)
	}
}

func (s *stockGrabber) SetIBMPrice(ibmPrice float64) {
	s.ibmPrice = ibmPrice
	s.NotifyObserver()
}

func (s *stockGrabber) SetGOOGPrice(googPrice float64) {
	s.googPrice = googPrice
	s.NotifyObserver()
}

func (s *stockGrabber) SetAAPLPrice(aaplPrice float64) {
	s.aaplPrice = aaplPrice
	s.NotifyObserver()
}

func NewStockGrabber() *stockGrabber {
	return &stockGrabber{}
}

// ////
type stockObserver struct {
	ibmPrice  float64
	aaplPrice float64
	googPrice float64

	observerID   string
	stockGrabber Subject
}

func (s *stockObserver) Update(ibmPrice, aaplPrice, googPrice float64) {
	s.ibmPrice = ibmPrice
	s.aaplPrice = aaplPrice
	s.googPrice = googPrice
	s.PrintPrices()
}

func (s *stockObserver) PrintPrices() {
	fmt.Println(fmt.Sprintf("\nNew pricing changes / Observer: %s", s.observerID))
	fmt.Println(fmt.Sprintf("ibmPrice: %f ", +s.ibmPrice))
	fmt.Println(fmt.Sprintf("aaplPrice: %f ", +s.aaplPrice))
	fmt.Println(fmt.Sprintf("googPrice: %f ", +s.googPrice))
}

func NewStockObserver(observerID string, stockGrabber Subject) *stockObserver {
	stockObserver := &stockObserver{
		stockGrabber: stockGrabber,
		observerID:   observerID,
	}
	stockGrabber.Register(stockObserver)
	return stockObserver
}
