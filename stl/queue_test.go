package stl

import (
	"testing"
)

func TestQueueBasicOperations(t *testing.T) {
	queue := NewQueue[int]()

	// Test Enqueue and Peek
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Test Size
	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}

	// Test Peek
	value, ok := queue.Peek()
	if !ok {
		t.Error("Peek should succeed on non-empty queue")
	}
	if value != 1 {
		t.Errorf("Expected peek value 1, got %d", value)
	}

	// Test Dequeue
	value, ok = queue.Dequeue()
	if !ok {
		t.Error("Dequeue should succeed on non-empty queue")
	}
	if value != 1 {
		t.Errorf("Expected dequeue value 1, got %d", value)
	}
	if queue.Size() != 2 {
		t.Errorf("Expected size 2 after dequeue, got %d", queue.Size())
	}

	// Test IsEmpty
	if queue.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	// Dequeue remaining elements
	queue.Dequeue()
	queue.Dequeue()

	// Test IsEmpty on empty queue
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeueing all elements")
	}

	// Test Dequeue on empty queue
	_, ok = queue.Dequeue()
	if ok {
		t.Error("Dequeue should fail on empty queue")
	}

}

func TestQueueEnqueueAll(t *testing.T) {
	queue := NewQueue[int]()
	items := []int{1, 2, 3, 4, 5}

	queue.EnqueueAll(items)

	if queue.Size() != 5 {
		t.Errorf("Expected size 5 after EnqueueAll, got %d", queue.Size())
	}

	value, _ := queue.Peek()
	if value != 1 {
		t.Errorf("Expected first element to be 1, got %d", value)
	}
}

func TestQueueWithCapacity(t *testing.T) {
	queue := NewQueueWithCapacity[int](10)

	for i := 0; i < 15; i++ {
		queue.Enqueue(i)
	}

	if queue.Size() != 15 {
		t.Errorf("Expected size 15, got %d", queue.Size())
	}
}

func TestQueueClear(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)

	queue.Clear()

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after Clear")
	}
}

func TestQueueContains(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if !queue.Contains(2) {
		t.Error("Queue should contain element 2")
	}

	if queue.Contains(4) {
		t.Error("Queue should not contain element 4")
	}
}
