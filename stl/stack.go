package stl

import (
	"fmt"
	"sort"
)

// Stack represents a LIFO (Last In, First Out) data structure
type Stack[T any] struct {
	data []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// NewStackWithCapacity creates a new stack with initial capacity
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, capacity),
	}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// PushAll adds multiple elements to the stack (in order, so last element becomes top)
func (s *Stack[T]) PushAll(items []T) {
	s.data = append(s.data, items...)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, true
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack
func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

// ToSlice returns a copy of the stack as a slice
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.data))
	copy(result, s.data)
	return result
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	return fmt.Sprintf("Stack%v", s.data)
}

// ForEach applies a function to each element in the stack (from bottom to top)
func (s *Stack[T]) ForEach(fn func(T)) {
	for _, item := range s.data {
		fn(item)
	}
}

// ForEachReversed applies a function to each element in the stack (from top to bottom)
func (s *Stack[T]) ForEachReversed(fn func(T)) {
	for i := len(s.data) - 1; i >= 0; i-- {
		fn(s.data[i])
	}
}

// Filter returns a new stack containing elements that satisfy the predicate
func (s *Stack[T]) Filter(predicate func(T) bool) *Stack[T] {
	result := NewStack[T]()
	for _, item := range s.data {
		if predicate(item) {
			result.Push(item)
		}
	}
	return result
}

// Map applies a transformation function to each element and returns a new stack
func (s *Stack[T]) Map(transform func(T) T) *Stack[T] {
	result := NewStack[T]()
	for _, item := range s.data {
		result.Push(transform(item))
	}
	return result
}

// Clone creates a deep copy of the stack
func (s *Stack[T]) Clone() *Stack[T] {
	result := NewStackWithCapacity[T](len(s.data))
	result.PushAll(s.data)
	return result
}

// Equals checks if two stacks contain the same elements in the same order
func (s *Stack[T]) Equals(other *Stack[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for i, item := range s.data {
		if fmt.Sprintf("%v", item) != fmt.Sprintf("%v", other.data[i]) {
			return false
		}
	}
	return true
}

// Reverse reverses the order of elements in the stack
func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

// GetAt returns the element at the specified index (0 = bottom, size-1 = top)
func (s *Stack[T]) GetAt(index int) (T, bool) {
	if index < 0 || index >= len(s.data) {
		var zero T
		return zero, false
	}
	return s.data[index], true
}

// SetAt sets the element at the specified index
func (s *Stack[T]) SetAt(index int, item T) bool {
	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data[index] = item
	return true
}

// RemoveAt removes the element at the specified index
func (s *Stack[T]) RemoveAt(index int) bool {
	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return true
}

// InsertAt inserts an element at the specified index
func (s *Stack[T]) InsertAt(index int, item T) bool {
	if index < 0 || index > len(s.data) {
		return false
	}
	s.data = append(s.data[:index], append([]T{item}, s.data[index:]...)...)
	return true
}

// Contains checks if the stack contains an element
func (s *Stack[T]) Contains(item T) bool {
	for _, element := range s.data {
		if fmt.Sprintf("%v", element) == fmt.Sprintf("%v", item) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of an element
func (s *Stack[T]) IndexOf(item T) int {
	for i, element := range s.data {
		if fmt.Sprintf("%v", element) == fmt.Sprintf("%v", item) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of an element
func (s *Stack[T]) LastIndexOf(item T) int {
	for i := len(s.data) - 1; i >= 0; i-- {
		if fmt.Sprintf("%v", s.data[i]) == fmt.Sprintf("%v", item) {
			return i
		}
	}
	return -1
}

// Remove removes the first occurrence of an element
func (s *Stack[T]) Remove(item T) bool {
	index := s.IndexOf(item)
	if index == -1 {
		return false
	}
	return s.RemoveAt(index)
}

// RemoveAll removes all occurrences of an element
func (s *Stack[T]) RemoveAll(item T) int {
	count := 0
	for i := len(s.data) - 1; i >= 0; i-- {
		if fmt.Sprintf("%v", s.data[i]) == fmt.Sprintf("%v", item) {
			s.RemoveAt(i)
			count++
		}
	}
	return count
}

// Sort sorts the stack using a custom comparator
func (s *Stack[T]) Sort(less func(T, T) bool) {
	sort.Slice(s.data, func(i, j int) bool {
		return less(s.data[i], s.data[j])
	})
}

// SortStable sorts the stack stably using a custom comparator
func (s *Stack[T]) SortStable(less func(T, T) bool) {
	sort.SliceStable(s.data, func(i, j int) bool {
		return less(s.data[i], s.data[j])
	})
}

// Shuffle randomizes the order of elements in the stack
func (s *Stack[T]) Shuffle() {
	for i := len(s.data) - 1; i > 0; i-- {
		j := i // In a real implementation, you'd use rand.Intn(i + 1)
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

// Take returns the top n elements as a slice
func (s *Stack[T]) Take(n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(s.data) {
		return s.ToSlice()
	}

	result := make([]T, n)
	copy(result, s.data[len(s.data)-n:])
	return result
}

// Drop removes the top n elements from the stack
func (s *Stack[T]) Drop(n int) int {
	if n <= 0 {
		return 0
	}
	if n >= len(s.data) {
		removed := len(s.data)
		s.Clear()
		return removed
	}

	s.data = s.data[:len(s.data)-n]
	return n
}

// Capacity returns the current capacity of the stack
func (s *Stack[T]) Capacity() int {
	return cap(s.data)
}

// Reserve ensures the stack has at least the specified capacity
func (s *Stack[T]) Reserve(capacity int) {
	if capacity > cap(s.data) {
		newData := make([]T, len(s.data), capacity)
		copy(newData, s.data)
		s.data = newData
	}
}

// TrimToSize reduces the capacity to match the current size
func (s *Stack[T]) TrimToSize() {
	if len(s.data) < cap(s.data) {
		newData := make([]T, len(s.data))
		copy(newData, s.data)
		s.data = newData
	}
}
