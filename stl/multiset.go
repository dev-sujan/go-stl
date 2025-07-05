package stl

import (
	"fmt"
	"sort"
)

// MultiSet represents a collection that allows duplicate elements with count tracking
type MultiSet[T comparable] struct {
	data map[T]int
}

// NewMultiSet creates a new empty multiset
func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{
		data: make(map[T]int),
	}
}

// NewMultiSetFromSlice creates a multiset from a slice
func NewMultiSetFromSlice[T comparable](slice []T) *MultiSet[T] {
	ms := NewMultiSet[T]()
	for _, item := range slice {
		ms.Add(item)
	}
	return ms
}

// Add adds an element to the multiset
func (ms *MultiSet[T]) Add(element T) {
	ms.data[element]++
}

// AddCount adds multiple occurrences of an element
func (ms *MultiSet[T]) AddCount(element T, count int) {
	if count > 0 {
		ms.data[element] += count
	}
}

// Remove removes one occurrence of an element
func (ms *MultiSet[T]) Remove(element T) bool {
	if count, exists := ms.data[element]; exists {
		if count == 1 {
			delete(ms.data, element)
		} else {
			ms.data[element] = count - 1
		}
		return true
	}
	return false
}

// RemoveAll removes all occurrences of an element
func (ms *MultiSet[T]) RemoveAll(element T) bool {
	if _, exists := ms.data[element]; exists {
		delete(ms.data, element)
		return true
	}
	return false
}

// RemoveCount removes a specific number of occurrences of an element
func (ms *MultiSet[T]) RemoveCount(element T, count int) bool {
	if currentCount, exists := ms.data[element]; exists {
		if count >= currentCount {
			delete(ms.data, element)
		} else {
			ms.data[element] = currentCount - count
		}
		return true
	}
	return false
}

// Count returns the number of occurrences of an element
func (ms *MultiSet[T]) Count(element T) int {
	return ms.data[element]
}

// Contains checks if an element exists in the multiset
func (ms *MultiSet[T]) Contains(element T) bool {
	_, exists := ms.data[element]
	return exists
}

// Size returns the total number of elements (including duplicates)
func (ms *MultiSet[T]) Size() int {
	total := 0
	for _, count := range ms.data {
		total += count
	}
	return total
}

// UniqueSize returns the number of unique elements
func (ms *MultiSet[T]) UniqueSize() int {
	return len(ms.data)
}

// IsEmpty checks if the multiset is empty
func (ms *MultiSet[T]) IsEmpty() bool {
	return len(ms.data) == 0
}

// Clear removes all elements from the multiset
func (ms *MultiSet[T]) Clear() {
	ms.data = make(map[T]int)
}

// ToSlice converts the multiset to a slice (with duplicates)
func (ms *MultiSet[T]) ToSlice() []T {
	result := make([]T, 0, ms.Size())
	for element, count := range ms.data {
		for i := 0; i < count; i++ {
			result = append(result, element)
		}
	}
	return result
}

// ToUniqueSlice converts the multiset to a slice of unique elements
func (ms *MultiSet[T]) ToUniqueSlice() []T {
	result := make([]T, 0, len(ms.data))
	for element := range ms.data {
		result = append(result, element)
	}
	return result
}

// ToCountMap returns a map of elements to their counts
func (ms *MultiSet[T]) ToCountMap() map[T]int {
	result := make(map[T]int)
	for element, count := range ms.data {
		result[element] = count
	}
	return result
}

// Union returns a new multiset containing elements from both multisets
func (ms *MultiSet[T]) Union(other *MultiSet[T]) *MultiSet[T] {
	result := NewMultiSet[T]()

	// Add all elements from current multiset
	for element, count := range ms.data {
		result.AddCount(element, count)
	}

	// Add all elements from other multiset
	for element, count := range other.data {
		result.AddCount(element, count)
	}

	return result
}

// Intersection returns a new multiset containing elements present in both multisets
func (ms *MultiSet[T]) Intersection(other *MultiSet[T]) *MultiSet[T] {
	result := NewMultiSet[T]()

	for element, count1 := range ms.data {
		if count2, exists := other.data[element]; exists {
			minCount := count1
			if count2 < minCount {
				minCount = count2
			}
			result.AddCount(element, minCount)
		}
	}

	return result
}

// Difference returns a new multiset containing elements in ms but not in other
func (ms *MultiSet[T]) Difference(other *MultiSet[T]) *MultiSet[T] {
	result := NewMultiSet[T]()

	for element, count1 := range ms.data {
		count2 := other.Count(element)
		if count1 > count2 {
			result.AddCount(element, count1-count2)
		}
	}

	return result
}

// IsSubset checks if ms is a subset of other
func (ms *MultiSet[T]) IsSubset(other *MultiSet[T]) bool {
	for element, count := range ms.data {
		if other.Count(element) < count {
			return false
		}
	}
	return true
}

// IsSuperset checks if ms is a superset of other
func (ms *MultiSet[T]) IsSuperset(other *MultiSet[T]) bool {
	return other.IsSubset(ms)
}

// Equals checks if two multisets contain the same elements with the same counts
func (ms *MultiSet[T]) Equals(other *MultiSet[T]) bool {
	if len(ms.data) != len(other.data) {
		return false
	}

	for element, count := range ms.data {
		if other.Count(element) != count {
			return false
		}
	}

	return true
}

// Clone creates a deep copy of the multiset
func (ms *MultiSet[T]) Clone() *MultiSet[T] {
	result := NewMultiSet[T]()
	for element, count := range ms.data {
		result.AddCount(element, count)
	}
	return result
}

// String returns a string representation of the multiset
func (ms *MultiSet[T]) String() string {
	return fmt.Sprintf("MultiSet%v", ms.ToCountMap())
}

// ForEach applies a function to each element in the multiset (including duplicates)
func (ms *MultiSet[T]) ForEach(fn func(T)) {
	for element, count := range ms.data {
		for i := 0; i < count; i++ {
			fn(element)
		}
	}
}

// ForEachUnique applies a function to each unique element in the multiset
func (ms *MultiSet[T]) ForEachUnique(fn func(T, int)) {
	for element, count := range ms.data {
		fn(element, count)
	}
}

// Filter returns a new multiset containing elements that satisfy the predicate
func (ms *MultiSet[T]) Filter(predicate func(T) bool) *MultiSet[T] {
	result := NewMultiSet[T]()
	for element, count := range ms.data {
		if predicate(element) {
			result.AddCount(element, count)
		}
	}
	return result
}

// MostCommon returns the most frequently occurring elements
// and LeastCommon returns the least frequently occurring elements.
// To avoid code duplication, the core logic is factored into a helper.
func (ms *MultiSet[T]) mostOrLeastCommon(n int, least bool) []T {
	if n <= 0 {
		return []T{}
	}
	type elementCount struct {
		element T
		count   int
	}
	var elements []elementCount
	for element, count := range ms.data {
		elements = append(elements, elementCount{element, count})
	}
	if least {
		sort.Slice(elements, func(i, j int) bool {
			return elements[i].count < elements[j].count
		})
	} else {
		sort.Slice(elements, func(i, j int) bool {
			return elements[i].count > elements[j].count
		})
	}
	result := make([]T, 0, n)
	for i := 0; i < n && i < len(elements); i++ {
		result = append(result, elements[i].element)
	}
	return result
}

func (ms *MultiSet[T]) MostCommon(n int) []T {
	return ms.mostOrLeastCommon(n, false)
}

func (ms *MultiSet[T]) LeastCommon(n int) []T {
	return ms.mostOrLeastCommon(n, true)
}
