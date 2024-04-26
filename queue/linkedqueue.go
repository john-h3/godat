package queue

type node[T any] struct {
	v    T
	next *node[T]
}

type LinkedQueue[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func (q *LinkedQueue[T]) Enqueue(t T) {
	if q.size == 0 {
		q.first = &node[T]{v: t}
		q.last = q.first
	} else {
		q.last.next = &node[T]{v: t}
		q.last = q.last.next
	}
	q.size++
}

func (q *LinkedQueue[T]) Dequeue() (v T, ok bool) {
	if q.size == 0 {
		return
	}
	v, ok = q.first.v, true
	q.first = q.first.next
	q.size--
	return v, true
}

func (q *LinkedQueue[T]) Peek() (v T, ok bool) {
	if q.size == 0 {
		return
	}
	return q.first.v, true
}

func (q *LinkedQueue[T]) Clear() {
	q.size = 0
	q.first = nil
	q.last = nil
}

func (q *LinkedQueue[T]) Size() int {
	return q.size
}
