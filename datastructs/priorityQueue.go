package datastructs

type compare func(interface{}, interface{}) int

// IntVal interface for custom types that need to be stored in PriorityQueue
type IntVal interface {
	GetVal() int
}

// PriorityQueue interface
type PriorityQueue interface {
	Peek() interface{}
	Size() int
	IsEmpty() bool
	Enqueue(value interface{})
	Dequeue() interface{}
	Remove(item interface{}) interface{}
}

type heap struct {
	size     int
	elements []interface{}
	compare  compare
	indexMap map[int]int
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
	return &heap{
		elements: make([]interface{}, size),
		compare:  comparator,
		size:     size,
		indexMap: make(map[int]int),
	}
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

func setIndex(h *heap, item interface{}, index int) {
	switch t := item.(type) {
	case IntVal:
		h.indexMap[t.GetVal()] = index
	case int:
		h.indexMap[t] = index
	default:
		// unknown type
	}
}

func getIndex(h *heap, item interface{}) int {
	switch t := item.(type) {
	case IntVal:
		return h.indexMap[t.GetVal()]
	case int:
		return h.indexMap[t]
	default:
		// unknown type
		return -1
	}
}

func swapIndexes(h *heap, item1, item2 interface{}) {
	i1 := getIndex(h, item1)
	i2 := getIndex(h, item2)
	setIndex(h, item2, i1)
	setIndex(h, item1, i2)
}

func siftUp(h *heap, index int) {
	for parent(index) >= 0 && h.compare(h.elements[parent(index)], h.elements[index]) > 0 {
		swapIndexes(h, h.elements[parent(index)], h.elements[index])
		h.elements[parent(index)], h.elements[index] = h.elements[index], h.elements[parent(index)]
		index = parent(index)
	}
}

func siftDown(h *heap, i int) bool {
	iStart := i
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
		swapIndexes(h, h.elements[i], h.elements[j])
		h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
		i = j
	}
	return i > iStart
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
	setIndex(h, value, h.size-1)
	siftUp(h, h.size-1)
}

func (h *heap) Dequeue() interface{} {
	result := h.elements[0]
	swapIndexes(h, h.elements[0], h.elements[h.size-1])
	h.elements[0], h.elements[h.size-1] = h.elements[h.size-1], h.elements[0]
	h.size--
	siftDown(h, 0)
	return result
}

func (h *heap) Remove(item interface{}) interface{} {
	index := getIndex(h, item)
	result := h.elements[index]
	if index == h.size-1 {
		h.size--
		return result
	}
	swapIndexes(h, h.elements[index], h.elements[h.size-1])
	h.elements[index], h.elements[h.size-1] = h.elements[h.size-1], h.elements[index]
	h.size--
	if !siftDown(h, index) {
		siftUp(h, index)
	}
	return result
}
