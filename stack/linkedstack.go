package stack

type node[T any] struct {
	v    T
	next *node[T]
}

type LinkedStack[T any] struct {
	sentinel node[T]
}

func (s *LinkedStack[T]) Push(t T) {
	head := s.sentinel.next
	s.sentinel.next = &node[T]{v: t, next: head}
}

func (s *LinkedStack[T]) Pop() (v T, ok bool) {
	head := s.sentinel.next
	if head != nil {
		v = head.v
		ok = true
		s.sentinel.next = head.next
	}
	return
}

func (s *LinkedStack[T]) Peek() (v T, ok bool) {
	if s.sentinel.next != nil {
		v = s.sentinel.next.v
		ok = true
	}
	return
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.sentinel.next == nil
}

func (s *LinkedStack[T]) Clear() {
	s.sentinel.next = nil
}
