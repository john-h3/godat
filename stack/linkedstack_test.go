package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPush(b *testing.B) {
	s := LinkedStack[int]{}
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func TestPush(t *testing.T) {
	s := LinkedStack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	node := s.sentinel.next
	assert.NotNil(t, node)
	assert.NotNil(t, node.v)
	assert.Equal(t, 3, node.v)
	node = node.next
	assert.NotNil(t, node)
	assert.NotNil(t, node.v)
	assert.Equal(t, 2, node.v)
	node = node.next
	assert.NotNil(t, node)
	assert.NotNil(t, node.v)
	assert.Equal(t, 1, node.v)
}

func TestPop(t *testing.T) {
	s := LinkedStack[int]{}
	s.sentinel.next = &node[int]{v: 1}
	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestPeek(t *testing.T) {
	s := LinkedStack[int]{}
	s.sentinel.next = &node[int]{v: 1}
	v, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestIsEmpty(t *testing.T) {
	s := LinkedStack[int]{}
	assert.True(t, s.IsEmpty())
	s.sentinel.next = &node[int]{v: 1}
	assert.False(t, s.IsEmpty())
}

func TestClear(t *testing.T) {
	s := LinkedStack[int]{}
	s.sentinel.next = &node[int]{v: 1}
	s.Clear()
	assert.Nil(t, s.sentinel.next)
}
