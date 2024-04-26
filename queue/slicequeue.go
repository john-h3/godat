package queue

type SliceQueue[T any] struct {
	slice []T
}

func (q *SliceQueue[T]) Enqueue(t T) {
	q.slice = append(q.slice, t)
}

func (q *SliceQueue[T]) Dequeue() (v T, ok bool) {
	if len(q.slice) > 0 {
		v, ok = q.slice[0], true
		q.slice = q.slice[1:]
	}
	return
}

func (q *SliceQueue[T]) Peek() (v T, ok bool) {
	if len(q.slice) > 0 {
		v, ok = q.slice[0], true
	}
	return
}

func (q *SliceQueue[T]) Clear() {
	q.slice = q.slice[:0]
}

func (q *SliceQueue[T]) Size() int {
	return len(q.slice)
}
