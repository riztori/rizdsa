package rizdsa

var _ Stack[any] = &SliceStack[any]{}

type SliceStack[T any] struct {
	data []T
}

func NewSliceStack[T any](capacity int) *SliceStack[T] {
	return &SliceStack[T]{
		data: make([]T, 0, capacity),
	}
}

func (s *SliceStack[T]) Push(e T) {
	if s == nil {
		s = NewSliceStack[T](defaultStackCapacity)
	}
	s.data = append(s.data, e)
}

func (s *SliceStack[T]) Pop() T {
	var e T
	size := s.Size()
	if size <= 0 {
		return e
	}
	e = s.data[size-1]
	s.data = s.data[:size-1]
	return e
}

func (s *SliceStack[T]) Peek() T {
	var e T
	size := s.Size()
	if size <= 0 {
		return e
	}
	e = s.data[size-1]
	return e
}

func (s *SliceStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceStack[T]) Size() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}
