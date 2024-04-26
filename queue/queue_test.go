package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkLinkedQueueEnqueue(b *testing.B) {
	benchmarkEnqueue(b, &LinkedQueue[int]{})
}

func BenchmarkSliceQueueEnqueue(b *testing.B) {
	benchmarkEnqueue(b, &SliceQueue[int]{})
}

func benchmarkEnqueue(b *testing.B, q Queue[int]) {
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func TestQueue(t *testing.T) {
	assert.Implements(t, (*Queue[int])(nil), new(LinkedQueue[int]))
	assert.Implements(t, (*Queue[int])(nil), new(SliceQueue[int]))
}

func TestLinkedQueueEnqueue(t *testing.T) {
	q := LinkedQueue[int]{}
	q.Enqueue(1)
	assert.NotNil(t, q.first)
	assert.Equal(t, 1, q.first.v)
	q.Enqueue(2)
	assert.NotNil(t, q.first)
	assert.NotNil(t, q.last)
	assert.Equal(t, 1, q.first.v)
	assert.Equal(t, 2, q.last.v)
	q.Enqueue(3)
	assert.NotNil(t, q.first)
	assert.NotNil(t, q.last)
	assert.Equal(t, 1, q.first.v)
	assert.Equal(t, 3, q.last.v)
}

func TestSliceQueueEnqueue(t *testing.T) {
	q := SliceQueue[int]{}
	q.Enqueue(1)
	assert.Equal(t, 1, q.slice[0])
	q.Enqueue(2)
	assert.Equal(t, 1, q.slice[0])
	assert.Equal(t, 2, q.slice[1])
	q.Enqueue(3)
	assert.Equal(t, 1, q.slice[0])
	assert.Equal(t, 2, q.slice[1])
	assert.Equal(t, 3, q.slice[2])
}

func constructLinkedQueue(q *LinkedQueue[int]) {
	last := &node[int]{v: 3}
	q.first = &node[int]{v: 1, next: &node[int]{v: 2, next: last}}
	q.last = last
	q.size = 3
}

func TestLinkedQueueDequeue(t *testing.T) {
	q := LinkedQueue[int]{}
	constructLinkedQueue(&q)
	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestSliceQueueDequeue(t *testing.T) {
	q := SliceQueue[int]{slice: []int{1, 2, 3}}
	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	v, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestLinkedQueuePeek(t *testing.T) {
	q := LinkedQueue[int]{}

	v, ok := q.Peek()
	assert.False(t, ok)
	assert.Zero(t, v)

	constructLinkedQueue(&q)
	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestSliceQueuePeek(t *testing.T) {
	q := SliceQueue[int]{}

	v, ok := q.Peek()
	assert.False(t, ok)
	assert.Zero(t, v)

	q.slice = []int{1}
	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestLinkedQueueClear(t *testing.T) {
	q := LinkedQueue[int]{}
	constructLinkedQueue(&q)

	assert.Positive(t, q.size)
	assert.NotNil(t, q.first)
	assert.NotNil(t, q.last)

	q.Clear()

	assert.Equal(t, 0, q.size)
	assert.Nil(t, q.first)
	assert.Nil(t, q.last)
}

func TestSliceQueueClear(t *testing.T) {
	q := SliceQueue[int]{slice: []int{1}}

	assert.Positive(t, len(q.slice))
	q.Clear()

	assert.Equal(t, 0, len(q.slice))
}

func TestLinkedQueueSize(t *testing.T) {
	q := LinkedQueue[int]{}
	assert.Equal(t, 0, q.Size())

	constructLinkedQueue(&q)
	assert.Equal(t, 3, q.Size())
}

func TestSliceQueueSize(t *testing.T) {
	q := SliceQueue[int]{}
	assert.Equal(t, 0, q.Size())

	q.slice = []int{1, 2, 3}
	assert.Equal(t, 3, q.Size())
}
