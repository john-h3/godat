package stack

type SliceStack[T any] struct {
	slice []T
}

func (s *SliceStack[T]) Push(t T) {
	s.slice = append(s.slice, t)
}

func (s *SliceStack[T]) Pop() (v T, ok bool) {
	len := len(s.slice)
	if len > 0 {
		v = s.slice[len-1]
		ok = true
		s.slice = s.slice[:len-1]
	}
	return
}

func (s *SliceStack[T]) Peek() (v T, ok bool) {
	len := len(s.slice)
	if len > 0 {
		v = s.slice[len-1]
		ok = true
	}
	return
}

func (s *SliceStack[T]) Clear() {
	s.slice = s.slice[:0]
}

func (s *SliceStack[T]) Size() int {
	return len(s.slice)
}
