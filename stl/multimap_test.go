package stl

import (
	"testing"
)

func TestMultiMapBasicOperations(t *testing.T) {
	mm := NewMultiMap[string, int]()

	// Test Put and Get
	mm.Put("fruits", 1)
	mm.Put("fruits", 2)
	mm.Put("vegetables", 3)

	fruits := mm.Get("fruits")
	if len(fruits) != 2 {
		t.Errorf("Expected 2 values for 'fruits', got %d", len(fruits))
	}
	if fruits[0] != 1 || fruits[1] != 2 {
		t.Errorf("Expected values [1, 2] for 'fruits', got %v", fruits)
	}

	vegetables := mm.Get("vegetables")
	if len(vegetables) != 1 || vegetables[0] != 3 {
		t.Errorf("Expected values [3] for 'vegetables', got %v", vegetables)
	}

	// Test non-existent key
	nonExistent := mm.Get("grains")
	if len(nonExistent) != 0 {
		t.Errorf("Expected empty slice for non-existent key, got %v", nonExistent)
	}
}

func TestMultiMapGetFirstLast(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("test", 1)
	mm.Put("test", 2)
	mm.Put("test", 3)

	// Test GetFirst
	first, exists := mm.GetFirst("test")
	if !exists {
		t.Error("GetFirst should return true for existing key")
	}
	if first != 1 {
		t.Errorf("Expected first value 1, got %d", first)
	}

	// Test GetLast
	last, exists := mm.GetLast("test")
	if !exists {
		t.Error("GetLast should return true for existing key")
	}
	if last != 3 {
		t.Errorf("Expected last value 3, got %d", last)
	}

	// Test non-existent key
	_, exists = mm.GetFirst("nonexistent")
	if exists {
		t.Error("GetFirst should return false for non-existent key")
	}
}

func TestMultiMapRemoveOperations(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("test", 1)
	mm.Put("test", 2)
	mm.Put("test", 3)
	mm.Put("other", 4)

	// Test Remove specific value
	mm.Remove("test", 2)
	values := mm.Get("test")
	if len(values) != 2 {
		t.Errorf("Expected 2 values after removal, got %d", len(values))
	}
	if values[0] != 1 || values[1] != 3 {
		t.Errorf("Expected values [1, 3] after removal, got %v", values)
	}

	// Test RemoveAll
	mm.RemoveAll("test")
	if mm.ContainsKey("test") {
		t.Error("MultiMap should not contain 'test' after RemoveAll")
	}
	if !mm.ContainsKey("other") {
		t.Error("MultiMap should still contain 'other' after RemoveAll")
	}
}

func TestMultiMapContainsOperations(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("test", 1)
	mm.Put("test", 2)

	// Test ContainsKey
	if !mm.ContainsKey("test") {
		t.Error("MultiMap should contain key 'test'")
	}
	if mm.ContainsKey("nonexistent") {
		t.Error("MultiMap should not contain key 'nonexistent'")
	}

	// Test ContainsValue
	if !mm.ContainsValue(1) {
		t.Error("MultiMap should contain value 1")
	}
	if mm.ContainsValue(3) {
		t.Error("MultiMap should not contain value 3")
	}

	// Test ContainsEntry
	if !mm.ContainsEntry("test", 1) {
		t.Error("MultiMap should contain entry ('test', 1)")
	}
	if mm.ContainsEntry("test", 3) {
		t.Error("MultiMap should not contain entry ('test', 3)")
	}
}

func TestMultiMapSizeOperations(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("a", 1)
	mm.Put("a", 2)
	mm.Put("b", 3)

	if mm.Size() != 3 {
		t.Errorf("Expected total size 3, got %d", mm.Size())
	}
	if mm.KeySize() != 2 {
		t.Errorf("Expected key size 2, got %d", mm.KeySize())
	}
	if mm.ValueCount("a") != 2 {
		t.Errorf("Expected value count 2 for 'a', got %d", mm.ValueCount("a"))
	}
}

func TestMultiMapCollections(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("a", 1)
	mm.Put("a", 2)
	mm.Put("b", 3)

	// Test Keys
	keys := mm.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	// Test Values
	values := mm.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Test UniqueValues
	uniqueValues := mm.UniqueValues()
	if len(uniqueValues) != 3 {
		t.Errorf("Expected 3 unique values, got %d", len(uniqueValues))
	}

	// Test Entries
	entries := mm.Entries()
	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}
}

func TestMultiMapConversions(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("a", 1)
	mm.Put("a", 2)
	mm.Put("b", 3)

	// Test ToMap
	singleMap := mm.ToMap()
	if len(singleMap) != 2 {
		t.Errorf("Expected map size 2, got %d", len(singleMap))
	}
	if singleMap["a"] != 2 {
		t.Errorf("Expected last value 2 for key 'a', got %d", singleMap["a"])
	}

	// Test ToMapOfSlices
	sliceMap := mm.ToMapOfSlices()
	if len(sliceMap["a"]) != 2 {
		t.Errorf("Expected slice length 2 for key 'a', got %d", len(sliceMap["a"]))
	}
	if len(sliceMap["b"]) != 1 {
		t.Errorf("Expected slice length 1 for key 'b', got %d", len(sliceMap["b"]))
	}
}

func TestMultiMapClearAndEmpty(t *testing.T) {
	mm := NewMultiMap[string, int]()

	if !mm.IsEmpty() {
		t.Error("New MultiMap should be empty")
	}

	mm.Put("test", 1)
	if mm.IsEmpty() {
		t.Error("MultiMap should not be empty after adding entry")
	}

	mm.Clear()
	if !mm.IsEmpty() {
		t.Error("MultiMap should be empty after Clear()")
	}
	if mm.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", mm.Size())
	}
}

func TestMultiMapFilter(t *testing.T) {
	mm := NewMultiMap[string, int]()
	mm.Put("even", 2)
	mm.Put("even", 4)
	mm.Put("odd", 1)
	mm.Put("odd", 3)

	// Test Filter
	evenOnly := mm.Filter(func(k string, v int) bool {
		return v%2 == 0
	})
	if evenOnly.Size() != 2 {
		t.Errorf("Expected filtered size 2, got %d", evenOnly.Size())
	}

	// Test FilterKeys
	evenKey := mm.FilterKeys(func(k string) bool {
		return k == "even"
	})
	if evenKey.Size() != 2 {
		t.Errorf("Expected key-filtered size 2, got %d", evenKey.Size())
	}

	// Test FilterValues
	oddValues := mm.FilterValues(func(v int) bool {
		return v%2 != 0
	})
	if oddValues.Size() != 2 {
		t.Errorf("Expected value-filtered size 2, got %d", oddValues.Size())
	}
}
