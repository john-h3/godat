package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSliceStackPush(b *testing.B) {
	s := SliceStack[int]{}
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkLinkedStackPush(b *testing.B) {
	s := LinkedStack[int]{}
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func TestStack(t *testing.T) {
	assert.Implements(t, (*Stack[int])(nil), new(LinkedStack[int]))
	assert.Implements(t, (*Stack[int])(nil), new(SliceStack[int]))
}

func TestLinkedStackPush(t *testing.T) {
	s := LinkedStack[int]{}
	s.Push(1)
	assert.NotNil(t, s.next)
	assert.NotNil(t, s.next.v)
	assert.Equal(t, 1, s.next.v)
}

func TestSliceStackPush(t *testing.T) {
	s := SliceStack[int]{}
	s.Push(1)
	assert.NotEmpty(t, s.slice)
	assert.Equal(t, 1, s.slice[0])
}

func TestLinkedStackPop(t *testing.T) {
	s := LinkedStack[int]{}
	_, ok := s.Pop()
	assert.False(t, ok)
	s.next = &node[int]{v: 1}
	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestSliceStackPop(t *testing.T) {
	s := SliceStack[int]{}
	_, ok := s.Pop()
	assert.False(t, ok)
	s.slice = []int{1}
	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestLinkedStackPeek(t *testing.T) {
	s := LinkedStack[int]{}
	_, ok := s.Peek()
	assert.False(t, ok)
	s.next = &node[int]{v: 1}
	v, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestSliceStackPeek(t *testing.T) {
	s := SliceStack[int]{}
	_, ok := s.Peek()
	assert.False(t, ok)
	s.slice = []int{1}
	v, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestLinkedStackSize(t *testing.T) {
	s := LinkedStack[int]{}
	assert.Equal(t, 0, s.Size())
	s.next = &node[int]{v: 1}
	s.size = 1
	assert.Equal(t, 1, s.Size())
}

func TestSliceStackSize(t *testing.T) {
	s := SliceStack[int]{}
	assert.Equal(t, 0, s.Size())
	s.slice = []int{1}
	assert.Equal(t, 1, s.Size())
}

func TestLinkedStackClear(t *testing.T) {
	s := LinkedStack[int]{}
	s.next = &node[int]{v: 1}
	assert.NotNil(t, s.next)
	s.Clear()
	assert.Nil(t, s.next)
}

func TestSliceStackClear(t *testing.T) {
	s := SliceStack[int]{[]int{1}}
	assert.NotEmpty(t, s.slice)
	s.Clear()
	assert.Empty(t, s.slice)
}
