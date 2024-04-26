package stack

type node[T any] struct {
	v    T
	next *node[T]
}

type LinkedStack[T any] struct {
	node[T]
	size int
}

func (s *LinkedStack[T]) Push(t T) {
	s.next = &node[T]{v: t, next: s.next}
	s.size++
}

func (s *LinkedStack[T]) Pop() (v T, ok bool) {
	if s.next != nil {
		v = s.next.v
		ok = true
		s.next = s.next.next
		s.size--
	}
	return
}

func (s *LinkedStack[T]) Peek() (v T, ok bool) {
	if s.next != nil {
		v = s.next.v
		ok = true
	}
	return
}

func (s *LinkedStack[T]) Clear() {
	s.next = nil
	s.size = 0
}

func (s *LinkedStack[T]) Size() int {
	return s.size
}
