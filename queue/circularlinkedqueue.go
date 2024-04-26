package queue

type CircularLinkedQueue[T any] struct {
	free  *node[T]
	first *node[T]
	size  int
}

func (q *CircularLinkedQueue[T]) Enqueue(t T) {
	if q.free == nil && q.first == nil {
		newNode := &node[T]{}
		newNode.next = newNode
		q.free = newNode
		q.first = newNode
	}
	if q.free.next == q.first {
		q.free.next = &node[T]{next: q.first}
	}
	q.free.v = t
	q.free = q.free.next
	q.size++
}

func (q *CircularLinkedQueue[T]) Dequeue() (v T, ok bool) {
	if q.size > 0 {
		v, ok = q.first.v, true
		q.first = q.first.next
		q.size--
	}
	return
}

func (q *CircularLinkedQueue[T]) Peek() (v T, ok bool) {
	if q.size == 0 {
		return
	}
	return q.first.v, true
}

func (q *CircularLinkedQueue[T]) Clear() {
	q.size = 0
	q.free = nil
	q.first = nil
}

func (q *CircularLinkedQueue[T]) Size() int {
	return q.size
}
