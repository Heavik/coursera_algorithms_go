package datastructs

import "testing"

func TestMinPriorityQueue(t *testing.T) {
	pq := NewMinPQ()
	vals := []int{2, 4, 13, 1, 7, 3}
	expect := []int{1, 2, 3, 4, 7, 13}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	for _, val := range expect {
		min := pq.Dequeue()
		if min != val {
			t.Errorf("Expected %v, but got %v", val, min)
		}
	}
}

func TestMinPqWithDuplicates(t *testing.T) {
	pq := NewMinPQ()
	vals := []int{2, 4, 13, 2, 1, 7, 3, 3}
	expect := []int{1, 2, 2, 3, 3, 4, 7, 13}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	for _, val := range expect {
		min := pq.Dequeue()
		if min != val {
			t.Errorf("Expected %v, but got %v", val, min)
		}
	}
}

func TestMinPriorityQueuePeek(t *testing.T) {
	pq := NewMinPQ()
	vals := []int{45, 4, 13, 1, 0, 7, 3, 78, 11}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	if pq.Peek() != 0 {
		t.Errorf("Expected %v, but got %v", 0, pq.Peek())
	}
}

func TestMaxPriorityQueuePeek(t *testing.T) {
	pq := NewMaxPQ()
	vals := []int{45, 4, 13, 1, 0, 7, 3, 78, 11}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	if pq.Peek() != 78 {
		t.Errorf("Expected %v, but got %v", 0, pq.Peek())
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	pq := NewMaxPQ()
	vals := []int{2, 4, 13, 1, 7, 3}
	expect := []int{13, 7, 4, 3, 2, 1}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	for _, val := range expect {
		min := pq.Dequeue()
		if min != val {
			t.Errorf("Expected %v, but got %v", val, min)
		}
	}
}

func TestMinPqEnqueueDequeueSequence(t *testing.T) {
	pq := NewMinPQ()
	vals := []int{2, 4, 13, 1, 7, 3}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	min := pq.Dequeue()
	min = pq.Dequeue()
	pq.Enqueue(0)
	pq.Enqueue(34)
	pq.Enqueue(12)
	min = pq.Dequeue()
	min = pq.Dequeue()
	if min != 3 {
		t.Errorf("Expected %v, but got %v", 3, min)
	}
}

func TestMaxPqEnqueueDequeueSequence(t *testing.T) {
	pq := NewMaxPQ()
	vals := []int{2, 4, 13, 1, 7, 3}
	for _, val := range vals {
		pq.Enqueue(val)
	}
	min := pq.Dequeue()
	min = pq.Dequeue()
	pq.Enqueue(0)
	pq.Enqueue(34)
	pq.Enqueue(12)
	min = pq.Dequeue()
	min = pq.Dequeue()
	if min != 12 {
		t.Errorf("Expected %v, but got %v", 12, min)
	}
}
