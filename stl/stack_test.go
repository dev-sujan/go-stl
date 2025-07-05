package stl

import (
	"testing"
)

func TestStackBasicOperations(t *testing.T) {
	stack := NewStack[int]()

	// Test Push and Peek
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test Size
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}

	// Test Peek
	value, ok := stack.Peek()
	if !ok {
		t.Error("Peek should succeed on non-empty stack")
	}
	if value != 3 {
		t.Errorf("Expected peek value 3, got %d", value)
	}

	// Test Pop
	value, ok = stack.Pop()
	if !ok {
		t.Error("Pop should succeed on non-empty stack")
	}
	if value != 3 {
		t.Errorf("Expected pop value 3, got %d", value)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}

	// Test IsEmpty
	if stack.IsEmpty() {
		t.Error("Stack should not be empty")
	}

	// Pop remaining elements
	stack.Pop()
	stack.Pop()

	// Test IsEmpty on empty stack
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	// Test Pop on empty stack
	_, ok = stack.Pop()
	if ok {
		t.Error("Pop should fail on empty stack")
	}
}

func TestStackPushAll(t *testing.T) {
	stack := NewStack[int]()
	items := []int{1, 2, 3, 4, 5}

	stack.PushAll(items)

	if stack.Size() != 5 {
		t.Errorf("Expected size 5 after PushAll, got %d", stack.Size())
	}

	value, _ := stack.Peek()
	if value != 5 {
		t.Errorf("Expected top element to be 5, got %d", value)
	}
}

func TestStackWithCapacity(t *testing.T) {
	stack := NewStackWithCapacity[int](10)

	for i := 0; i < 15; i++ {
		stack.Push(i)
	}

	if stack.Size() != 15 {
		t.Errorf("Expected size 15, got %d", stack.Size())
	}
}

func TestStackClear(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	stack.Clear()

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after Clear")
	}
}

func TestStackContains(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !stack.Contains(2) {
		t.Error("Stack should contain element 2")
	}

	if stack.Contains(4) {
		t.Error("Stack should not contain element 4")
	}
}
