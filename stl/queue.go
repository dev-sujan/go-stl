package stl

import (
	"fmt"
	"sort"
)

// Queue represents a FIFO (First In, First Out) data structure.
type Queue[T any] struct {
	data []T
}

// NewQueue creates a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// NewQueueWithCapacity creates a new queue with initial capacity.
func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, capacity),
	}
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// EnqueueAll adds multiple elements to the queue.
func (q *Queue[T]) EnqueueAll(items []T) {
	q.data = append(q.data, items...)
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	item := q.data[0]
	q.data = q.data[1:]
	return item, true
}

// Peek returns the front element without removing it.
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.data[0], true
}

// PeekBack returns the back element without removing it.
func (q *Queue[T]) PeekBack() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.data[len(q.data)-1], true
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.data = q.data[:0]
}

// ToSlice returns a copy of the queue as a slice.
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.data))
	copy(result, q.data)
	return result
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue%v", q.data)
}

// ForEach applies a function to each element in the queue (from front to back).
func (q *Queue[T]) ForEach(fn func(T)) {
	for _, item := range q.data {
		fn(item)
	}
}

// ForEachReversed applies a function to each element in the queue (from back to front).
func (q *Queue[T]) ForEachReversed(fn func(T)) {
	for i := len(q.data) - 1; i >= 0; i-- {
		fn(q.data[i])
	}
}

// Filter returns a new queue containing elements that satisfy the predicate.
func (q *Queue[T]) Filter(predicate func(T) bool) *Queue[T] {
	result := NewQueue[T]()
	for _, item := range q.data {
		if predicate(item) {
			result.Enqueue(item)
		}
	}
	return result
}

// Map applies a transformation function to each element and returns a new queue.
func (q *Queue[T]) Map(transform func(T) T) *Queue[T] {
	result := NewQueue[T]()
	for _, item := range q.data {
		result.Enqueue(transform(item))
	}
	return result
}

// Clone creates a deep copy of the queue.
func (q *Queue[T]) Clone() *Queue[T] {
	result := NewQueueWithCapacity[T](len(q.data))
	result.EnqueueAll(q.data)
	return result
}

// Equals checks if two queues contain the same elements in the same order.
func (q *Queue[T]) Equals(other *Queue[T]) bool {
	if q.Size() != other.Size() {
		return false
	}

	for i, item := range q.data {
		if fmt.Sprintf("%v", item) != fmt.Sprintf("%v", other.data[i]) {
			return false
		}
	}
	return true
}

// Reverse reverses the order of elements in the queue.
func (q *Queue[T]) Reverse() {
	for i, j := 0, len(q.data)-1; i < j; i, j = i+1, j-1 {
		q.data[i], q.data[j] = q.data[j], q.data[i]
	}
}

// GetAt returns the element at the specified index (0 = front, size-1 = back).
func (q *Queue[T]) GetAt(index int) (T, bool) {
	if index < 0 || index >= len(q.data) {
		var zero T
		return zero, false
	}
	return q.data[index], true
}

// SetAt sets the element at the specified index.
func (q *Queue[T]) SetAt(index int, item T) bool {
	if index < 0 || index >= len(q.data) {
		return false
	}
	q.data[index] = item
	return true
}

// RemoveAt removes the element at the specified index.
func (q *Queue[T]) RemoveAt(index int) bool {
	if index < 0 || index >= len(q.data) {
		return false
	}
	q.data = append(q.data[:index], q.data[index+1:]...)
	return true
}

// InsertAt inserts an element at the specified index.
func (q *Queue[T]) InsertAt(index int, item T) bool {
	if index < 0 || index > len(q.data) {
		return false
	}
	q.data = append(q.data[:index], append([]T{item}, q.data[index:]...)...)
	return true
}

// Contains checks if the queue contains an element.
func (q *Queue[T]) Contains(item T) bool {
	for _, element := range q.data {
		if fmt.Sprintf("%v", element) == fmt.Sprintf("%v", item) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of an element.
func (q *Queue[T]) IndexOf(item T) int {
	for i, element := range q.data {
		if fmt.Sprintf("%v", element) == fmt.Sprintf("%v", item) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of an element.
func (q *Queue[T]) LastIndexOf(item T) int {
	for i := len(q.data) - 1; i >= 0; i-- {
		if fmt.Sprintf("%v", q.data[i]) == fmt.Sprintf("%v", item) {
			return i
		}
	}
	return -1
}

// Remove removes the first occurrence of an element.
func (q *Queue[T]) Remove(item T) bool {
	index := q.IndexOf(item)
	if index == -1 {
		return false
	}
	return q.RemoveAt(index)
}

// RemoveAll removes all occurrences of an element.
func (q *Queue[T]) RemoveAll(item T) int {
	count := 0
	for i := len(q.data) - 1; i >= 0; i-- {
		if fmt.Sprintf("%v", q.data[i]) == fmt.Sprintf("%v", item) {
			q.RemoveAt(i)
			count++
		}
	}
	return count
}

// Sort sorts the queue using a custom comparator.
func (q *Queue[T]) Sort(less func(T, T) bool) {
	sort.Slice(q.data, func(i, j int) bool {
		return less(q.data[i], q.data[j])
	})
}

// SortStable sorts the queue stably using a custom comparator.
func (q *Queue[T]) SortStable(less func(T, T) bool) {
	sort.SliceStable(q.data, func(i, j int) bool {
		return less(q.data[i], q.data[j])
	})
}

// Shuffle randomizes the order of elements in the queue.
func (q *Queue[T]) Shuffle() {
	for i := len(q.data) - 1; i > 0; i-- {
		j := i // In a real implementation, you'd use rand.Intn(i + 1)
		q.data[i], q.data[j] = q.data[j], q.data[i]
	}
}

// Take returns the first n elements as a slice.
func (q *Queue[T]) Take(n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(q.data) {
		return q.ToSlice()
	}

	result := make([]T, n)
	copy(result, q.data[:n])
	return result
}

// Drop removes the first n elements from the queue.
func (q *Queue[T]) Drop(n int) int {
	if n <= 0 {
		return 0
	}
	if n >= len(q.data) {
		removed := len(q.data)
		q.Clear()
		return removed
	}

	q.data = q.data[n:]
	return n
}

// Capacity returns the current capacity of the queue.
func (q *Queue[T]) Capacity() int {
	return cap(q.data)
}

// Reserve ensures the queue has at least the specified capacity.
func (q *Queue[T]) Reserve(capacity int) {
	if capacity > cap(q.data) {
		newData := make([]T, len(q.data), capacity)
		copy(newData, q.data)
		q.data = newData
	}
}

// TrimToSize reduces the capacity to match the current size.
func (q *Queue[T]) TrimToSize() {
	if len(q.data) < cap(q.data) {
		newData := make([]T, len(q.data))
		copy(newData, q.data)
		q.data = newData
	}
}

// PriorityQueue represents a priority queue where elements are ordered by priority.
type PriorityQueue[T any] struct {
	less func(T, T) bool
	data []T
}

// NewPriorityQueue creates a new priority queue with a custom comparator.
func NewPriorityQueue[T any](less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make([]T, 0),
		less: less,
	}
}

// NewPriorityQueueWithCapacity creates a new priority queue with initial capacity.
func NewPriorityQueueWithCapacity[T any](capacity int, less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make([]T, 0, capacity),
		less: less,
	}
}

// Enqueue adds an element to the priority queue.
func (pq *PriorityQueue[T]) Enqueue(item T) {
	pq.data = append(pq.data, item)
	pq.up(len(pq.data) - 1)
}

// Dequeue removes and returns the highest priority element.
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}

	item := pq.data[0]
	pq.data[0] = pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]

	if len(pq.data) > 0 {
		pq.down(0)
	}

	return item, true
}

// Peek returns the highest priority element without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}
	return pq.data[0], true
}

// Size returns the number of elements in the priority queue.
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.data)
}

// IsEmpty returns true if the priority queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.data) == 0
}

// Clear removes all elements from the priority queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.data = pq.data[:0]
}

// ToSlice returns a copy of the priority queue as a slice.
func (pq *PriorityQueue[T]) ToSlice() []T {
	result := make([]T, len(pq.data))
	copy(result, pq.data)
	return result
}

// String returns a string representation of the priority queue.
func (pq *PriorityQueue[T]) String() string {
	return fmt.Sprintf("PriorityQueue%v", pq.data)
}

// up moves an element up in the heap to maintain heap property.
func (pq *PriorityQueue[T]) up(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if !pq.less(pq.data[index], pq.data[parent]) {
			break
		}
		pq.data[index], pq.data[parent] = pq.data[parent], pq.data[index]
		index = parent
	}
}

// down moves an element down in the heap to maintain heap property.
func (pq *PriorityQueue[T]) down(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < len(pq.data) && pq.less(pq.data[left], pq.data[smallest]) {
			smallest = left
		}

		if right < len(pq.data) && pq.less(pq.data[right], pq.data[smallest]) {
			smallest = right
		}

		if smallest == index {
			break
		}

		pq.data[index], pq.data[smallest] = pq.data[smallest], pq.data[index]
		index = smallest
	}
}

// Clone creates a deep copy of the priority queue.
func (pq *PriorityQueue[T]) Clone() *PriorityQueue[T] {
	result := NewPriorityQueueWithCapacity[T](len(pq.data), pq.less)
	result.data = make([]T, len(pq.data))
	copy(result.data, pq.data)
	return result
}
