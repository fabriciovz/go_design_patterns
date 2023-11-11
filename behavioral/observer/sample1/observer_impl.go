package sample1

import "fmt"

// REF: https://refactoring.guru/design-patterns/observer
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
	fmt.Println(fmt.Sprintf("\nNew Observer has been created: %s ", stockObserver.observerID))
	return stockObserver
}
