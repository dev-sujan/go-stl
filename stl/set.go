package stl

import (
	"fmt"
)

// Set represents an unordered collection of unique elements.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates a new empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// NewSetFromSlice creates a set from a slice, removing duplicates.
func NewSetFromSlice[T comparable](slice []T) *Set[T] {
	s := NewSet[T]()
	for _, item := range slice {
		s.Add(item)
	}
	return s
}

// Add adds an element to the set.
func (s *Set[T]) Add(element T) {
	s.data[element] = struct{}{}
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

// Contains checks if an element exists in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// ToSlice converts the set to a slice.
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.data))
	for element := range s.data {
		result = append(result, element)
	}
	return result
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	// Add all elements from current set
	for element := range s.data {
		result.Add(element)
	}

	// Add all elements from other set
	for element := range other.data {
		result.Add(element)
	}

	return result
}

// Intersection returns a new set containing elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for element := range s.data {
		if other.Contains(element) {
			result.Add(element)
		}
	}

	return result
}

// Difference returns a new set containing elements in s but not in other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}

	return result
}

// SymmetricDifference returns a new set containing elements in either set but not both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	union := s.Union(other)
	intersection := s.Intersection(other)
	return union.Difference(intersection)
}

// IsSubset checks if s is a subset of other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSuperset checks if s is a superset of other.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// IsDisjoint checks if s and other have no elements in common.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	for element := range s.data {
		if other.Contains(element) {
			return false
		}
	}
	return true
}

// Equals checks if two sets contain the same elements.
func (s *Set[T]) Equals(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	return s.IsSubset(other)
}

// Clone creates a deep copy of the set.
func (s *Set[T]) Clone() *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		result.Add(element)
	}
	return result
}

// String returns a string representation of the set.
func (s *Set[T]) String() string {
	return fmt.Sprintf("Set%v", s.ToSlice())
}

// ForEach applies a function to each element in the set.
func (s *Set[T]) ForEach(fn func(T)) {
	for element := range s.data {
		fn(element)
	}
}

// Filter returns a new set containing elements that satisfy the predicate.
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if predicate(element) {
			result.Add(element)
		}
	}
	return result
}

// Any returns true if any element satisfies the predicate.
func (s *Set[T]) Any(predicate func(T) bool) bool {
	for element := range s.data {
		if predicate(element) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate.
func (s *Set[T]) All(predicate func(T) bool) bool {
	for element := range s.data {
		if !predicate(element) {
			return false
		}
	}
	return true
}
