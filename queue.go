package queue

import (
	"container/list"
	"fmt"
)

// Queue is a basic FIFO queue
type Queue interface {
	Add(value interface{}) error
	Take() interface{}
	Clear()
	Len() int
}

// BoundedQueue is a queue that has a specified size.
type BoundedQueue struct {
	nodes chan interface{}
}

// New returns a new BoundedQueue with the specified buffer size.
// Size is forced to be greater than zero, will default to 1.
func NewBounded(size int) *BoundedQueue {
	if size < 1 {
		size = 1
	}
	return &BoundedQueue{
		nodes: make(chan interface{}, size),
	}
}

// Push adds a node to the queue. Concurrency safe.
func (q *BoundedQueue) Add(value interface{}) error {
	select {
	case q.nodes <- value:
		return nil
	default:
		return fmt.Errorf("Unable to add value, queue is full.")
	}
}

// Take removes and returns a node from the queue in first to last order.
// Returns nil if there is nothing in the queue. Concurrency safe.
func (q *BoundedQueue) Take() interface{} {
	select {
	case node := <-q.nodes:
		return node
	default:
		return nil
	}
}

// Clear empties the queue and leaves it at a freshly initialized state.
// Not concurrency safe.
func (q *BoundedQueue) Clear() {
	close(q.nodes)
	q.nodes = make(chan interface{}, cap(q.nodes))
}

// Len returns the length of the items in the queue. Concurrency safe.
func (q *BoundedQueue) Len() int {
	return len(q.nodes)
}

// Unbounded queue is a queue that can continuously be added to.
type UnboundedQueue struct {
	nodes *list.List
}

// New returns a new UnboundedQueue with the specified buffer size, must be greater than 1
func NewUnbounded() *UnboundedQueue {
	return &UnboundedQueue{
		nodes: list.New(),
	}
}

// Push adds a node to the queue. Not concurrency safe.
func (q *UnboundedQueue) Add(value interface{}) error {
	q.nodes.PushBack(value)
	return nil
}

// Take removes and returns a node from the queue in first to last order.
// Returns nil if there is nothing in the queue. Not concurrency safe.
func (q *UnboundedQueue) Take() interface{} {
	// Get our reference
	node := q.nodes.Front()

	if node == nil {
		return nil
	}

	// Delete the head of the list
	q.nodes.Remove(node)

	return node.Value
}

// Clear empties the queue and leaves it at a freshly initialized state.
// Not concurrency safe.
func (q *UnboundedQueue) Clear() {
	q.nodes.Init()
}

// Len returns the length of the items in the queue. Not concurrency safe.
func (q *UnboundedQueue) Len() int {
	return q.nodes.Len()
}
