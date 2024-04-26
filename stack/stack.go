package stack

type Stack[T any] interface {
	Push(t T)
	Pop() (T, bool)
	Peek() (T, bool)
	Clear()
	Size() int
}
