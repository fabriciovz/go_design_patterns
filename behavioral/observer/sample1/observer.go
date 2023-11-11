package sample1

// Subscriber
type Observer interface {
	Update(ibmPrice, aaplPrice, googPrice float64)
}
