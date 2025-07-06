package stl

import (
	"fmt"
	"sort"
)

// MultiMap represents a map that allows multiple values per key.
type MultiMap[K comparable, V any] struct {
	data map[K][]V
}

// NewMultiMap creates a new empty multimap.
func NewMultiMap[K comparable, V any]() *MultiMap[K, V] {
	return &MultiMap[K, V]{
		data: make(map[K][]V),
	}
}

// Put adds a value to the multimap for the given key.
func (mm *MultiMap[K, V]) Put(key K, value V) {
	mm.data[key] = append(mm.data[key], value)
}

// PutAll adds multiple values to the multimap for the given key.
func (mm *MultiMap[K, V]) PutAll(key K, values []V) {
	mm.data[key] = append(mm.data[key], values...)
}

// Get returns all values associated with the given key.
func (mm *MultiMap[K, V]) Get(key K) []V {
	if values, exists := mm.data[key]; exists {
		// Return a copy to prevent external modification
		result := make([]V, len(values))
		copy(result, values)
		return result
	}
	return []V{}
}

// GetFirst returns the first value associated with the given key.
func (mm *MultiMap[K, V]) GetFirst(key K) (V, bool) {
	if values, exists := mm.data[key]; exists && len(values) > 0 {
		return values[0], true
	}
	var zero V
	return zero, false
}

// GetLast returns the last value associated with the given key.
func (mm *MultiMap[K, V]) GetLast(key K) (V, bool) {
	if values, exists := mm.data[key]; exists && len(values) > 0 {
		return values[len(values)-1], true
	}
	var zero V
	return zero, false
}

// Remove removes a specific value from the multimap for the given key.
func (mm *MultiMap[K, V]) Remove(key K, value V) bool {
	if values, exists := mm.data[key]; exists {
		for i, v := range values {
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", value) {
				// Remove the element at index i
				mm.data[key] = append(values[:i], values[i+1:]...)
				// If no values left for this key, remove the key
				if len(mm.data[key]) == 0 {
					delete(mm.data, key)
				}
				return true
			}
		}
	}
	return false
}

// RemoveAll removes all values for the given key.
func (mm *MultiMap[K, V]) RemoveAll(key K) bool {
	if _, exists := mm.data[key]; exists {
		delete(mm.data, key)
		return true
	}
	return false
}

// RemoveKey removes all values for the given key (alias for RemoveAll).
func (mm *MultiMap[K, V]) RemoveKey(key K) bool {
	return mm.RemoveAll(key)
}

// ContainsKey checks if the multimap contains the given key.
func (mm *MultiMap[K, V]) ContainsKey(key K) bool {
	_, exists := mm.data[key]
	return exists
}

// ContainsValue checks if the multimap contains the given value.
func (mm *MultiMap[K, V]) ContainsValue(value V) bool {
	for _, values := range mm.data {
		for _, v := range values {
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", value) {
				return true
			}
		}
	}
	return false
}

// ContainsEntry checks if the multimap contains the given key-value pair.
func (mm *MultiMap[K, V]) ContainsEntry(key K, value V) bool {
	if values, exists := mm.data[key]; exists {
		for _, v := range values {
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", value) {
				return true
			}
		}
	}
	return false
}

// Size returns the total number of key-value pairs.
func (mm *MultiMap[K, V]) Size() int {
	total := 0
	for _, values := range mm.data {
		total += len(values)
	}
	return total
}

// KeySize returns the number of unique keys.
func (mm *MultiMap[K, V]) KeySize() int {
	return len(mm.data)
}

// ValueCount returns the number of values for a given key.
func (mm *MultiMap[K, V]) ValueCount(key K) int {
	if values, exists := mm.data[key]; exists {
		return len(values)
	}
	return 0
}

// IsEmpty checks if the multimap is empty.
func (mm *MultiMap[K, V]) IsEmpty() bool {
	return len(mm.data) == 0
}

// Clear removes all elements from the multimap.
func (mm *MultiMap[K, V]) Clear() {
	mm.data = make(map[K][]V)
}

// Keys returns all keys in the multimap.
func (mm *MultiMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(mm.data))
	for key := range mm.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values in the multimap.
func (mm *MultiMap[K, V]) Values() []V {
	var values []V
	for _, vals := range mm.data {
		values = append(values, vals...)
	}
	return values
}

// UniqueValues returns unique values in the multimap.
func (mm *MultiMap[K, V]) UniqueValues() []V {
	valueSet := make(map[string]V)
	for _, vals := range mm.data {
		for _, val := range vals {
			key := fmt.Sprintf("%v", val)
			valueSet[key] = val
		}
	}

	values := make([]V, 0, len(valueSet))
	for _, val := range valueSet {
		values = append(values, val)
	}
	return values
}

// Entries returns all key-value pairs as a slice of Entry structs.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func (mm *MultiMap[K, V]) Entries() []Entry[K, V] {
	var entries []Entry[K, V]
	for key, values := range mm.data {
		for _, value := range values {
			entries = append(entries, Entry[K, V]{Key: key, Value: value})
		}
	}
	return entries
}

// ToMap converts the multimap to a regular map (keeping only the last value for each key).
func (mm *MultiMap[K, V]) ToMap() map[K]V {
	result := make(map[K]V)
	for key, values := range mm.data {
		if len(values) > 0 {
			result[key] = values[len(values)-1]
		}
	}
	return result
}

// ToMapOfSlices converts the multimap to a map of slices.
func (mm *MultiMap[K, V]) ToMapOfSlices() map[K][]V {
	result := make(map[K][]V)
	for key, values := range mm.data {
		result[key] = make([]V, len(values))
		copy(result[key], values)
	}
	return result
}

// String returns a string representation of the multimap.
func (mm *MultiMap[K, V]) String() string {
	return fmt.Sprintf("MultiMap%v", mm.ToMapOfSlices())
}

// ForEach applies a function to each key-value pair.
func (mm *MultiMap[K, V]) ForEach(fn func(K, V)) {
	for key, values := range mm.data {
		for _, value := range values {
			fn(key, value)
		}
	}
}

// ForEachKey applies a function to each key and its associated values.
func (mm *MultiMap[K, V]) ForEachKey(fn func(K, []V)) {
	for key, values := range mm.data {
		valuesCopy := make([]V, len(values))
		copy(valuesCopy, values)
		fn(key, valuesCopy)
	}
}

// Filter returns a new multimap containing entries that satisfy the predicate.
func (mm *MultiMap[K, V]) Filter(predicate func(K, V) bool) *MultiMap[K, V] {
	result := NewMultiMap[K, V]()
	for key, values := range mm.data {
		for _, value := range values {
			if predicate(key, value) {
				result.Put(key, value)
			}
		}
	}
	return result
}

// FilterKeys returns a new multimap containing entries with keys that satisfy the predicate.
func (mm *MultiMap[K, V]) FilterKeys(predicate func(K) bool) *MultiMap[K, V] {
	result := NewMultiMap[K, V]()
	for key, values := range mm.data {
		if predicate(key) {
			for _, value := range values {
				result.Put(key, value)
			}
		}
	}
	return result
}

// FilterValues returns a new multimap containing entries with values that satisfy the predicate.
func (mm *MultiMap[K, V]) FilterValues(predicate func(V) bool) *MultiMap[K, V] {
	result := NewMultiMap[K, V]()
	for key, values := range mm.data {
		for _, value := range values {
			if predicate(value) {
				result.Put(key, value)
			}
		}
	}
	return result
}

// Clone creates a deep copy of the multimap.
func (mm *MultiMap[K, V]) Clone() *MultiMap[K, V] {
	result := NewMultiMap[K, V]()
	for key, values := range mm.data {
		for _, value := range values {
			result.Put(key, value)
		}
	}
	return result
}

// Equals checks if two multimaps contain the same key-value pairs.
func (mm *MultiMap[K, V]) Equals(other *MultiMap[K, V]) bool {
	if mm.KeySize() != other.KeySize() {
		return false
	}

	for key, values1 := range mm.data {
		values2 := other.Get(key)
		if len(values1) != len(values2) {
			return false
		}

		// Create maps to compare values (order doesn't matter)
		valueMap1 := make(map[string]int)
		valueMap2 := make(map[string]int)

		for _, v := range values1 {
			key := fmt.Sprintf("%v", v)
			valueMap1[key]++
		}

		for _, v := range values2 {
			key := fmt.Sprintf("%v", v)
			valueMap2[key]++
		}

		if len(valueMap1) != len(valueMap2) {
			return false
		}

		for k, count := range valueMap1 {
			if valueMap2[k] != count {
				return false
			}
		}
	}

	return true
}

// GetSortedKeys returns keys sorted by a custom comparator.
func (mm *MultiMap[K, V]) GetSortedKeys(less func(K, K) bool) []K {
	keys := mm.Keys()
	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})
	return keys
}

// GetSortedValues returns values for a key sorted by a custom comparator.
func (mm *MultiMap[K, V]) GetSortedValues(key K, less func(V, V) bool) []V {
	values := mm.Get(key)
	sort.Slice(values, func(i, j int) bool {
		return less(values[i], values[j])
	})
	return values
}
