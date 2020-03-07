package datastructs

import "errors"

type queue struct {
	head   int
	tail   int
	values []int
}

// Queue interface
type Queue interface {
	IsEmpty() bool
	Enqueue(value int)
	Dequeue() (int, error)
}

// EmptyQueue creates new empty queue
func EmptyQueue() Queue {
	return &queue{}
}

// NewQueue creates empty queue with given starting capacity
func NewQueue(capacity int) Queue {
	return &queue{values: make([]int, capacity)}
}

func (q *queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *queue) Enqueue(value int) {
	if q.tail >= len(q.values) {
		q.values = append(q.values, value)
	} else {
		q.values[q.tail] = value
	}
	q.tail++
}

func (q *queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		err := errors.New("The queue is empty")
		return 0, err
	}
	val := q.values[q.head]
	q.head++
	if q.head >= q.tail {
		q.head = 0
		q.tail = 0
	}
	return val, nil
}
