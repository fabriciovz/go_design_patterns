package iterator

type Iterator interface {
	hasNext() bool
	next() *User
}
