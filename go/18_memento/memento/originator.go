package memento

type originator interface {
	setState(state [3]([3]string))
	getState() [3]([3]string)
	backup() *memento
}