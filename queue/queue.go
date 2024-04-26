package queue

type Queue[T any] interface {
	Enqueue(t T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Clear()
	Size() int
}
