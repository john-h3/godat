package stack

type node[T any] struct {
	v    T
	next *node[T]
}

type LinkedStack[T any] struct {
	node[T]
}

func (s *LinkedStack[T]) Push(t T) {
	s.next = &node[T]{v: t, next: s.next}
}

func (s *LinkedStack[T]) Pop() (v T, ok bool) {
	if s.next != nil {
		v = s.next.v
		ok = true
		s.next = s.next.next
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

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.next == nil
}

func (s *LinkedStack[T]) Clear() {
	s.next = nil
}
