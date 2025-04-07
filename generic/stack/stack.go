package stack

type CustomStack[T string | int64] struct {
	stack []T
}

func (s *CustomStack[T]) Push(val T) {
	s.stack = append(s.stack, val)
}

func (s *CustomStack[T]) Pop() T {
	if len(s.stack) == 0 {
		var zero T
		return zero
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

func (s *CustomStack[T]) Peek() T {
	if len(s.stack) == 0 {
		var zero T
		return zero
	}
	return s.stack[len(s.stack)-1]
}

func (s *CustomStack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}
