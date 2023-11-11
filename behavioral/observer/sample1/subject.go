package sample1

// Subject/Publisher
type Subject interface {
	Register(o Observer)
	Unregister(o Observer)
	NotifyAll()
}
