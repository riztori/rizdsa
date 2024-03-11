package rizdsa

var _ Queue[any] = &SliceQueue[any]{}

type SliceQueue[T any] struct {
	data []T
}

func NewSliceQueue[T any](capacity int) *SliceQueue[T] {
	return &SliceQueue[T]{
		data: make([]T, 0, capacity),
	}
}

func (s *SliceQueue[T]) Enqueue(e T) {
	if s == nil {
		s = NewSliceQueue[T](defaultQueueCapacity)
	}
	s.data = append(s.data, e)
}

func (s *SliceQueue[T]) Dequeue() T {
	var e T
	size := s.Size()
	if size <= 0 {
		return e
	}
	e = s.data[0]
	s.data = s.data[1:]
	return e
}

func (s *SliceQueue[T]) Front() T {
	var e T
	size := s.Size()
	if size <= 0 {
		return e
	}
	return s.data[0]
}

func (s *SliceQueue[T]) Tail() T {
	var e T
	size := s.Size()
	if size <= 0 {
		return e
	}
	return s.data[size-1]
}

func (s *SliceQueue[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceQueue[T]) Size() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}
