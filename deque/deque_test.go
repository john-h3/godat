package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDequeEnqueueFront(b *testing.B) {
	q := Deque[int]{}
	for i := 0; i < b.N; i++ {
		q.EnqueueFront(i)
	}
}

func BenchmarkDequeEnqueueBack(b *testing.B) {
	q := Deque[int]{}
	for i := 0; i < b.N; i++ {
		q.EnqueueBack(i)
	}
}

func TestDequeEnqueueFront(t *testing.T) {
	q := Deque[int]{}
	q.EnqueueFront(1)
	assert.NotNil(t, q.first)
	assert.Equal(t, 1, q.first.v)

	q.EnqueueFront(2)
	assert.Equal(t, 2, q.first.v)

	q.EnqueueFront(3)
	assert.Equal(t, 3, q.first.v)
}

func TestDequeEnqueueBack(t *testing.T) {
	q := Deque[int]{}
	q.EnqueueBack(1)
	assert.NotNil(t, q.last)
	assert.Equal(t, 1, q.last.v)

	q.EnqueueBack(2)
	assert.Equal(t, 2, q.last.v)

	q.EnqueueBack(3)
	assert.Equal(t, 3, q.last.v)
}

func constructDeque(q *Deque[int]) {
	second := &node[int]{v: 2}
	q.first = &node[int]{v: 1, next: second}
	q.last = &node[int]{v: 3, prev: second}
	second.prev, second.next = q.first, q.last
	q.size = 3
}

func TestDequeDequeueFront(t *testing.T) {
	q := &Deque[int]{}
	_, ok := q.DequeueFront()
	assert.False(t, ok)

	constructDeque(q)

	v, ok := q.DequeueFront()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = q.DequeueFront()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.DequeueFront()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}

func TestDequeDequeueBack(t *testing.T) {
	q := &Deque[int]{}
	_, ok := q.DequeueBack()
	assert.False(t, ok)

	constructDeque(q)

	v, ok := q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.DequeueBack()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestDequePeekFront(t *testing.T) {
	q := &Deque[int]{}
	_, ok := q.PeekFront()
	assert.False(t, ok)

	constructDeque(q)
	v, ok := q.PeekFront()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestDequePeekBack(t *testing.T) {
	q := &Deque[int]{}
	_, ok := q.PeekBack()
	assert.False(t, ok)

	constructDeque(q)
	v, ok := q.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}

func TestDequeClear(t *testing.T) {
	q := &Deque[int]{}
	constructDeque(q)
	assert.NotNil(t, q.first)
	assert.NotNil(t, q.last)
	assert.Positive(t, 3, q.size)

	q.Clear()
	assert.Nil(t, q.first)
	assert.Nil(t, q.last)
	assert.Equal(t, 0, q.size)
}

func TestDequeSize(t *testing.T) {
	q := &Deque[int]{}
	assert.Equal(t, 0, q.size)

	constructDeque(q)
	assert.Equal(t, 3, q.size)
}
