package datastructs

type compare func(interface{}, interface{}) int

// PriorityQueue interface
type PriorityQueue interface {
	Peek() interface{}
	Size() int
	IsEmpty() bool
	Enqueue(value interface{})
	Dequeue() interface{}
}

type heap struct {
	size     int
	elements []interface{}
	compare  compare
}

// NewPQ creates new priority queue with init size = 0
func NewPQ(comparator compare) PriorityQueue {
	return NewSizedPQ(comparator, 0)
}

// NewMinPQ creates new min priority queue with init size = 0
func NewMinPQ() PriorityQueue {
	return NewPQ(func(a, b interface{}) int { return a.(int) - b.(int) })
}

// NewMaxPQ creates new max priority queue with init size = 0
func NewMaxPQ() PriorityQueue {
	return NewPQ(func(a, b interface{}) int { return b.(int) - a.(int) })
}

// NewSizedPQ creates new priority queue with given size and comparator
func NewSizedPQ(comparator compare, size int) PriorityQueue {
	return &heap{elements: make([]interface{}, size), compare: comparator, size: size}
}

// FindMedian finds median mod len(arr) in array using two heaps (assigment answer is 1213)
func FindMedian(arr []int) int {
	h1 := NewMaxPQ()
	h2 := NewMinPQ()

	result := 0
	i := 0
	for _, val := range arr {
		i++
		if h1.IsEmpty() || val < h1.Peek().(int) {
			h1.Enqueue(val)
		} else if h2.IsEmpty() || val > h2.Peek().(int) {
			h2.Enqueue(val)
		} else {
			h1.Enqueue(val)
		}
		if i%2 == 0 {
			if h1.Size()-h2.Size() >= 2 {
				h2.Enqueue(h1.Dequeue())
			} else if h2.Size()-h1.Size() >= 2 {
				h1.Enqueue(h2.Dequeue())
			}
		}
		if h1.Size() == h2.Size() || h1.Size() > h2.Size() {
			result += h1.Peek().(int)
		} else {
			result += h2.Peek().(int)
		}
	}
	return result % len(arr)
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

func (h *heap) Peek() interface{} {
	if h.IsEmpty() {
		return 0
	}
	return h.elements[0]
}

func (h *heap) Size() int {
	return h.size
}

func (h *heap) IsEmpty() bool {
	return h.size == 0
}

func (h *heap) Enqueue(value interface{}) {
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

func (h *heap) Dequeue() interface{} {
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
