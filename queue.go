package rizdsa

const defaultQueueCapacity int = 10

type Queue[T any] interface {
	Enqueue(e T)
	Dequeue() T
	Front() T
	Tail() T
	IsEmpty() bool
	Size() int
}
