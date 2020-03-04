package datastructs

import "testing"

func TestQueueReturnsValuesInFifo(t *testing.T) {
	q := EmptyQueue()
	order := []int{1, 2, 3}
	for _, val := range order {
		q.Enqueue(val)
	}

	for i := 0; !q.IsEmpty(); i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if val != order[i] {
			t.Errorf("expected %v but got %v", order[i], val)
		}
	}
}

func TestEmptyQueueReturnsErrorOnDequeue(t *testing.T) {
	q := EmptyQueue()

	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Empty queue should return an error, but got %v", err)
	}
}

func TestQueueIsEmptyFunctionReturnsTrue(t *testing.T) {
	q := EmptyQueue()
	q.Enqueue(42)
	q.Dequeue()

	if !q.IsEmpty() {
		t.Error("IsEmpty() should return true on a empty queue")
	}
}

func TestQueueEnqDeqSequences(t *testing.T) {
	q := EmptyQueue()
	q.Enqueue(42)
	q.Enqueue(43)
	q.Dequeue()
	q.Dequeue()

	for _, val := range []int{3, 6, 7, 9} {
		q.Enqueue(val)
	}

	q.Dequeue()
	q.Dequeue()
	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if val != 7 {
		t.Errorf("Expected %v, but got %v", 7, val)
	}
}
