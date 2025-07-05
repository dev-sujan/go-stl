package stl

import (
	"testing"
)

func TestDequeBasicOperations(t *testing.T) {
	deque := NewDeque[int](10)

	// Test PushFront and PushBack
	deque.PushFront(1)
	deque.PushBack(2)
	deque.PushFront(0)
	deque.PushBack(3)

	// Test Size
	if deque.Size() != 4 {
		t.Errorf("Expected size 4, got %d", deque.Size())
	}

	// Test Front and Back
	front, ok := deque.Front()
	if !ok || front != 0 {
		t.Errorf("Expected front value 0, got %d, ok: %v", front, ok)
	}

	back, ok := deque.Back()
	if !ok || back != 3 {
		t.Errorf("Expected back value 3, got %d, ok: %v", back, ok)
	}

	// Test PopFront
	value, ok := deque.PopFront()
	if !ok || value != 0 {
		t.Errorf("Expected popped front value 0, got %d, ok: %v", value, ok)
	}
	if deque.Size() != 3 {
		t.Errorf("Expected size 3 after pop, got %d", deque.Size())
	}

	// Test PopBack
	value, ok = deque.PopBack()
	if !ok || value != 3 {
		t.Errorf("Expected popped back value 3, got %d, ok: %v", value, ok)
	}
	if deque.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", deque.Size())
	}
}

func TestDequeEmptyOperations(t *testing.T) {
	deque := NewDeque[int](10)

	// Test IsEmpty
	if !deque.IsEmpty() {
		t.Error("New deque should be empty")
	}

	// Test Front/Back on empty deque
	_, ok := deque.Front()
	if ok {
		t.Error("Front should fail on empty deque")
	}

	_, ok = deque.Back()
	if ok {
		t.Error("Back should fail on empty deque")
	}

	// Test PopFront/PopBack on empty deque
	_, ok = deque.PopFront()
	if ok {
		t.Error("PopFront should fail on empty deque")
	}

	_, ok = deque.PopBack()
	if ok {
		t.Error("PopBack should fail on empty deque")
	}
}

func TestDequeAutoResize(t *testing.T) {
	// Start with small capacity
	deque := NewDeque[int](2)

	// Add enough items to trigger resize
	for i := 0; i < 20; i++ {
		deque.PushBack(i)
	}

	if deque.Size() != 20 {
		t.Errorf("Expected size 20 after pushes, got %d", deque.Size())
	}

	// Verify values are correct after resize
	for i := 0; i < 20; i++ {
		value, ok := deque.PopFront()
		if !ok || value != i {
			t.Errorf("Expected value %d, got %d, ok: %v", i, value, ok)
		}
	}

	if !deque.IsEmpty() {
		t.Error("Deque should be empty after popping all items")
	}
}

func TestDequeCircularBuffer(t *testing.T) {
	// Test the circular buffer behavior
	deque := NewDeque[int](4) // Small capacity to test wraparound

	// Fill up to capacity
	for i := 0; i < 4; i++ {
		deque.PushBack(i)
	}

	// Remove from front and add to back to force wraparound
	for i := 0; i < 4; i++ {
		deque.PopFront()
		deque.PushBack(i + 10)
	}

	// Verify values
	for i := 0; i < 4; i++ {
		value, ok := deque.PopFront()
		if !ok || value != i+10 {
			t.Errorf("Expected value %d, got %d, ok: %v", i+10, value, ok)
		}
	}
}

func TestDequeClear(t *testing.T) {
	deque := NewDeque[int](10)

	for i := 0; i < 5; i++ {
		deque.PushBack(i)
	}

	deque.Clear()

	if !deque.IsEmpty() {
		t.Error("Deque should be empty after Clear()")
	}
	if deque.Size() != 0 {
		t.Errorf("Deque size should be 0 after Clear(), got %d", deque.Size())
	}
}

func TestDequeFromSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	deque := NewDequeFromSlice(items)

	if deque.Size() != len(items) {
		t.Errorf("Expected size %d, got %d", len(items), deque.Size())
	}

	for i := 0; i < len(items); i++ {
		value, ok := deque.PopFront()
		if !ok || value != items[i] {
			t.Errorf("Expected value %d, got %d", items[i], value)
		}
	}
}

// TestDequeContains is skipped as the method is not implemented
func TestDequeContains(t *testing.T) {
	t.Skip("Contains method not implemented yet")
}
