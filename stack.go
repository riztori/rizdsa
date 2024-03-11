package rizdsa

const defaultStackCapacity int = 10

type Stack[T any] interface {
	Push(e T)
	Pop() T
	Peek() T
	Size() int
	IsEmpty() bool
}
