package sample2

// Subscriber
type Observer interface {
	update(string)
	getID() string
}
