package datastructs

type compare func(int, int) int

// PriorityQueue interface
type PriorityQueue interface {
	Peek() int
	Size() int
	Enqueue(value int)
	Dequeue() int
}

type heap struct {
	size     int
	elements []int
	compare  compare
}

// NewPQ creates new priority queue with init size = 0
func NewPQ(comparator compare) PriorityQueue {
	return NewSizedPQ(comparator, 0)
}

// NewMinPQ creates new min priority queue with init size = 0
func NewMinPQ() PriorityQueue {
	return NewPQ(func(a, b int) int { return a - b })
}

// NewMaxPQ creates new max priority queue with init size = 0
func NewMaxPQ() PriorityQueue {
	return NewPQ(func(a, b int) int { return b - a })
}

// NewSizedPQ creates new priority queue with given size and comparator
func NewSizedPQ(comparator compare, size int) PriorityQueue {
	return &heap{elements: make([]int, size), compare: comparator, size: size}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *heap) Peek() int {
	return h.elements[0]
}

func (h *heap) Size() int {
	return h.size
}

func (h *heap) Enqueue(value int) {
	if h.size < len(h.elements) {
		h.elements[h.size] = value
	} else {
		h.elements = append(h.elements, value)
	}
	h.size++
	index := h.size - 1
	for parent(index) >= 0 && h.compare(h.elements[parent(index)], h.elements[index]) > 0 {
		h.elements[parent(index)], h.elements[index] = h.elements[index], h.elements[parent(index)]
		index = parent(index)
	}
}

func (h *heap) Dequeue() int {
	result := h.elements[0]
	h.elements[0], h.elements[h.size-1] = h.elements[h.size-1], h.elements[0]
	h.size--
	i := 0
	for {
		j := i
		if left(i) < h.size && h.compare(h.elements[left(i)], h.elements[j]) < 0 {
			j = left(i)
		}
		if right(i) < h.size && h.compare(h.elements[right(i)], h.elements[j]) < 0 {
			j = right(i)
		}
		if i == j {
			break
		}
		h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
		i = j
	}
	return result
}
