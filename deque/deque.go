package deque

type node[T any] struct {
	v    T
	next *node[T]
	prev *node[T]
}

type Deque[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func (q *Deque[T]) EnqueueFront(t T) {
	n := &node[T]{v: t, next: q.first}
	if q.first == nil && q.last == nil {
		q.first = n
		q.last = n
	} else {
		q.first.prev = n
		q.first = n
	}
	q.size++
}

func (q *Deque[T]) DequeueFront() (v T, ok bool) {
	if q.size > 0 {
		v, ok = q.first.v, true
		q.first = q.first.next
		if q.first != nil {
			q.first.prev = nil
		}
		q.size--
	}
	return
}

func (q *Deque[T]) EnqueueBack(t T) {
	n := &node[T]{v: t, prev: q.last}
	if q.first == nil && q.last == nil {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		q.last = n
	}
	q.size++
}

func (q *Deque[T]) DequeueBack() (v T, ok bool) {
	if q.size > 0 {
		v, ok = q.last.v, true
		q.last = q.last.prev
		if q.last != nil {
			q.last.next = nil
		}
		q.size--
	}
	return
}

func (q *Deque[T]) PeekFront() (v T, ok bool) {
	if q.size > 0 {
		v, ok = q.first.v, true
	}
	return
}

func (q *Deque[T]) PeekBack() (v T, ok bool) {
	if q.size > 0 {
		v, ok = q.last.v, true
	}
	return
}

func (q *Deque[T]) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}

func (q *Deque[T]) Size() int {
	return q.size
}
