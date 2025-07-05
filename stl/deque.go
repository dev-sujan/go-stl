package stl

import (
	"fmt"
	"math"
)

// Deque represents a double-ended queue
type Deque[T any] struct {
	data  []T
	front int
	back  int
	size  int
}

// NewDeque creates a new empty deque with initial capacity
func NewDeque[T any](initialCapacity int) *Deque[T] {
	if initialCapacity <= 0 {
		initialCapacity = 16
	}
	return &Deque[T]{
		data:  make([]T, initialCapacity),
		front: 0,
		back:  0,
		size:  0,
	}
}

// NewDequeFromSlice creates a deque from a slice
func NewDequeFromSlice[T any](slice []T) *Deque[T] {
	capacity := len(slice)
	if capacity == 0 {
		capacity = 16
	} else {
		// Round up to next power of 2 for better performance
		capacity = int(math.Pow(2, math.Ceil(math.Log2(float64(capacity)))))
	}

	d := &Deque[T]{
		data:  make([]T, capacity),
		front: 0,
		back:  0,
		size:  0,
	}

	for _, item := range slice {
		d.PushBack(item)
	}

	return d
}

// ensureCapacity ensures the deque has enough capacity
func (d *Deque[T]) ensureCapacity() {
	if d.size == len(d.data) {
		// Need to grow
		newCapacity := len(d.data) * 2
		newData := make([]T, newCapacity)

		// Copy elements to new array
		for i := 0; i < d.size; i++ {
			newData[i] = d.data[(d.front+i)%len(d.data)]
		}

		d.data = newData
		d.front = 0
		d.back = d.size
	}
}

// PushFront adds an element to the front of the deque
func (d *Deque[T]) PushFront(element T) {
	d.ensureCapacity()

	d.front = (d.front - 1 + len(d.data)) % len(d.data)
	d.data[d.front] = element
	d.size++
}

// PushBack adds an element to the back of the deque
func (d *Deque[T]) PushBack(element T) {
	d.ensureCapacity()

	d.data[d.back] = element
	d.back = (d.back + 1) % len(d.data)
	d.size++
}

// PopFront removes and returns the element from the front of the deque
func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	element := d.data[d.front]
	d.front = (d.front + 1) % len(d.data)
	d.size--

	return element, true
}

// PopBack removes and returns the element from the back of the deque
func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}

	d.back = (d.back - 1 + len(d.data)) % len(d.data)
	element := d.data[d.back]
	d.size--

	return element, true
}

// Front returns the element at the front of the deque without removing it
func (d *Deque[T]) Front() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.data[d.front], true
}

// Back returns the element at the back of the deque without removing it
func (d *Deque[T]) Back() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	backIndex := (d.back - 1 + len(d.data)) % len(d.data)
	return d.data[backIndex], true
}

// At returns the element at the specified index
func (d *Deque[T]) At(index int) (T, bool) {
	if index < 0 || index >= d.size {
		var zero T
		return zero, false
	}
	actualIndex := (d.front + index) % len(d.data)
	return d.data[actualIndex], true
}

// Set sets the element at the specified index
func (d *Deque[T]) Set(index int, element T) bool {
	if index < 0 || index >= d.size {
		return false
	}
	actualIndex := (d.front + index) % len(d.data)
	d.data[actualIndex] = element
	return true
}

// Size returns the number of elements in the deque
func (d *Deque[T]) Size() int {
	return d.size
}

// IsEmpty checks if the deque is empty
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Clear removes all elements from the deque
func (d *Deque[T]) Clear() {
	d.front = 0
	d.back = 0
	d.size = 0
	// Clear the underlying array to help with garbage collection
	for i := range d.data {
		var zero T
		d.data[i] = zero
	}
}

// Capacity returns the current capacity of the deque
func (d *Deque[T]) Capacity() int {
	return len(d.data)
}

// Reserve ensures the deque has at least the specified capacity
func (d *Deque[T]) Reserve(capacity int) {
	if capacity > len(d.data) {
		newData := make([]T, capacity)

		// Copy elements to new array
		for i := 0; i < d.size; i++ {
			newData[i] = d.data[(d.front+i)%len(d.data)]
		}

		d.data = newData
		d.front = 0
		d.back = d.size
	}
}

// ShrinkToFit reduces the capacity to match the size
func (d *Deque[T]) ShrinkToFit() {
	if d.size < len(d.data) {
		newData := make([]T, d.size)

		// Copy elements to new array
		for i := 0; i < d.size; i++ {
			newData[i] = d.data[(d.front+i)%len(d.data)]
		}

		d.data = newData
		d.front = 0
		d.back = d.size
	}
}

// ToSlice converts the deque to a slice
func (d *Deque[T]) ToSlice() []T {
	result := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		result[i] = d.data[(d.front+i)%len(d.data)]
	}
	return result
}

// String returns a string representation of the deque
func (d *Deque[T]) String() string {
	return fmt.Sprintf("Deque%v", d.ToSlice())
}

// ForEach applies a function to each element in the deque
func (d *Deque[T]) ForEach(fn func(T)) {
	for i := 0; i < d.size; i++ {
		fn(d.data[(d.front+i)%len(d.data)])
	}
}

// ForEachIndex applies a function to each element and its index in the deque
func (d *Deque[T]) ForEachIndex(fn func(int, T)) {
	for i := 0; i < d.size; i++ {
		fn(i, d.data[(d.front+i)%len(d.data)])
	}
}

// Filter returns a new deque containing elements that satisfy the predicate
func (d *Deque[T]) Filter(predicate func(T) bool) *Deque[T] {
	result := NewDeque[T](d.size)
	for i := 0; i < d.size; i++ {
		element := d.data[(d.front+i)%len(d.data)]
		if predicate(element) {
			result.PushBack(element)
		}
	}
	return result
}

// Any returns true if any element satisfies the predicate
func (d *Deque[T]) Any(predicate func(T) bool) bool {
	for i := 0; i < d.size; i++ {
		element := d.data[(d.front+i)%len(d.data)]
		if predicate(element) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate
func (d *Deque[T]) All(predicate func(T) bool) bool {
	for i := 0; i < d.size; i++ {
		element := d.data[(d.front+i)%len(d.data)]
		if !predicate(element) {
			return false
		}
	}
	return true
}

// Clone creates a deep copy of the deque
func (d *Deque[T]) Clone() *Deque[T] {
	result := NewDeque[T](d.size)
	for i := 0; i < d.size; i++ {
		result.PushBack(d.data[(d.front+i)%len(d.data)])
	}
	return result
}

// Equals checks if two deques contain the same elements in the same order
func (d *Deque[T]) Equals(other *Deque[T]) bool {
	if d.size != other.size {
		return false
	}

	for i := 0; i < d.size; i++ {
		element1 := d.data[(d.front+i)%len(d.data)]
		element2 := other.data[(other.front+i)%len(other.data)]
		if fmt.Sprintf("%v", element1) != fmt.Sprintf("%v", element2) {
			return false
		}
	}

	return true
}

// Reverse reverses the order of elements in the deque
func (d *Deque[T]) Reverse() {
	if d.size <= 1 {
		return
	}

	// Create a temporary slice and reverse it
	temp := d.ToSlice()
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	// Copy back to deque
	for i, element := range temp {
		d.data[(d.front+i)%len(d.data)] = element
	}
}

// RotateLeft rotates the deque left by n positions
func (d *Deque[T]) RotateLeft(n int) {
	if d.size <= 1 || n == 0 {
		return
	}

	n = n % d.size
	if n < 0 {
		n += d.size
	}

	for i := 0; i < n; i++ {
		if element, ok := d.PopFront(); ok {
			d.PushBack(element)
		}
	}
}

// RotateRight rotates the deque right by n positions
func (d *Deque[T]) RotateRight(n int) {
	if d.size <= 1 || n == 0 {
		return
	}

	n = n % d.size
	if n < 0 {
		n += d.size
	}

	for i := 0; i < n; i++ {
		if element, ok := d.PopBack(); ok {
			d.PushFront(element)
		}
	}
}

// Swap swaps elements at the specified indices
func (d *Deque[T]) Swap(i, j int) bool {
	if i < 0 || i >= d.size || j < 0 || j >= d.size {
		return false
	}

	index1 := (d.front + i) % len(d.data)
	index2 := (d.front + j) % len(d.data)

	d.data[index1], d.data[index2] = d.data[index2], d.data[index1]
	return true
}

// Insert inserts an element at the specified index
func (d *Deque[T]) Insert(index int, element T) bool {
	if index < 0 || index > d.size {
		return false
	}

	if index == 0 {
		d.PushFront(element)
		return true
	}

	if index == d.size {
		d.PushBack(element)
		return true
	}

	// Insert in the middle - need to shift elements
	d.ensureCapacity()

	// Shift elements to make room
	for i := d.size; i > index; i-- {
		srcIndex := (d.front + i - 1) % len(d.data)
		dstIndex := (d.front + i) % len(d.data)
		d.data[dstIndex] = d.data[srcIndex]
	}

	// Insert the new element
	insertIndex := (d.front + index) % len(d.data)
	d.data[insertIndex] = element
	d.back = (d.back + 1) % len(d.data)
	d.size++

	return true
}

// Remove removes the element at the specified index
func (d *Deque[T]) Remove(index int) (T, bool) {
	if index < 0 || index >= d.size {
		var zero T
		return zero, false
	}

	if index == 0 {
		return d.PopFront()
	}

	if index == d.size-1 {
		return d.PopBack()
	}

	// Remove from the middle - need to shift elements
	removeIndex := (d.front + index) % len(d.data)
	element := d.data[removeIndex]

	// Shift elements to fill the gap
	for i := index; i < d.size-1; i++ {
		srcIndex := (d.front + i + 1) % len(d.data)
		dstIndex := (d.front + i) % len(d.data)
		d.data[dstIndex] = d.data[srcIndex]
	}

	d.back = (d.back - 1 + len(d.data)) % len(d.data)
	d.size--

	return element, true
}
