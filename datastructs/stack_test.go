package datastructs

import "testing"

func TestStackReturnsValuesInLifo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	stack := Stack()
	for _, v := range values {
		stack.Push(v)
	}
	for i := len(values) - 1; i >= 0; i-- {
		v, err := stack.Pop()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if v != values[i] {
			t.Errorf("expected %v, but got %v", values[i], v)
		}
	}
}

func TestEmptyStackReturnsErrorOnPop(t *testing.T) {
	stack := Stack()
	_, err := stack.Pop()
	if err == nil {
		t.Error("Pop() shuold return an error if stack is empty")
	}
}

func TestStackIsEmptyReturnsTrue(t *testing.T) {
	stack := Stack()
	stack.Push(42)
	stack.Pop()

	if !stack.IsEmpty() {
		t.Error("IsEmpty() should return true if stack us empty")
	}
}

func TestStackSeqPushPop(t *testing.T) {
	stack := Stack()
	stack.Push(42)
	stack.Push(34)
	stack.Pop()
	stack.Pop()

	for _, v := range []int{1, 4, 5, 6, 7} {
		stack.Push(v)
	}
	stack.Pop()
	stack.Pop()

	if val, _ := stack.Pop(); val != 5 {
		t.Errorf("expected %v, but got %v", 5, val)
	}
}
