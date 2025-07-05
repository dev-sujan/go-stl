package stl

import (
	"testing"
)

func TestMultiSetBasicOperations(t *testing.T) {
	ms := NewMultiSet[string]()

	// Test Add and Count
	ms.Add("apple")
	ms.Add("apple")
	ms.Add("banana")

	if ms.Count("apple") != 2 {
		t.Errorf("Expected count 2 for 'apple', got %d", ms.Count("apple"))
	}
	if ms.Count("banana") != 1 {
		t.Errorf("Expected count 1 for 'banana', got %d", ms.Count("banana"))
	}
	if ms.Count("orange") != 0 {
		t.Errorf("Expected count 0 for 'orange', got %d", ms.Count("orange"))
	}

	// Test Size and UniqueSize
	if ms.Size() != 3 {
		t.Errorf("Expected total size 3, got %d", ms.Size())
	}
	if ms.UniqueSize() != 2 {
		t.Errorf("Expected unique size 2, got %d", ms.UniqueSize())
	}
}

func TestMultiSetAddCount(t *testing.T) {
	ms := NewMultiSet[string]()

	ms.AddCount("apple", 3)
	if ms.Count("apple") != 3 {
		t.Errorf("Expected count 3 for 'apple', got %d", ms.Count("apple"))
	}

	ms.AddCount("apple", 2)
	if ms.Count("apple") != 5 {
		t.Errorf("Expected count 5 for 'apple', got %d", ms.Count("apple"))
	}
}

func TestMultiSetRemoveOperations(t *testing.T) {
	ms := NewMultiSet[string]()

	// Add items
	ms.AddCount("apple", 3)
	ms.AddCount("banana", 2)

	// Test Remove (single)
	ms.Remove("apple")
	if ms.Count("apple") != 2 {
		t.Errorf("Expected count 2 for 'apple' after remove, got %d", ms.Count("apple"))
	}

	// Test RemoveCount
	ms.RemoveCount("apple", 2)
	if ms.Count("apple") != 0 {
		t.Errorf("Expected count 0 for 'apple' after removeCount, got %d", ms.Count("apple"))
	}
	if ms.Contains("apple") {
		t.Error("MultiSet should not contain 'apple' after removing all occurrences")
	}

	// Test RemoveAll
	ms.RemoveAll("banana")
	if ms.Contains("banana") {
		t.Error("MultiSet should not contain 'banana' after removeAll")
	}
}

func TestMultiSetFromSlice(t *testing.T) {
	slice := []string{"apple", "apple", "banana", "apple", "orange"}
	ms := NewMultiSetFromSlice(slice)

	if ms.Count("apple") != 3 {
		t.Errorf("Expected count 3 for 'apple', got %d", ms.Count("apple"))
	}
	if ms.Count("banana") != 1 {
		t.Errorf("Expected count 1 for 'banana', got %d", ms.Count("banana"))
	}
	if ms.Count("orange") != 1 {
		t.Errorf("Expected count 1 for 'orange', got %d", ms.Count("orange"))
	}
}

func TestMultiSetSetOperations(t *testing.T) {
	ms1 := NewMultiSet[string]()
	ms1.AddCount("apple", 3)
	ms1.AddCount("banana", 2)

	ms2 := NewMultiSet[string]()
	ms2.AddCount("apple", 2)
	ms2.AddCount("orange", 1)

	// Test Union
	union := ms1.Union(ms2)
	if union.Count("apple") != 5 {
		t.Errorf("Expected union count 5 for 'apple', got %d", union.Count("apple"))
	}
	if union.Count("banana") != 2 {
		t.Errorf("Expected union count 2 for 'banana', got %d", union.Count("banana"))
	}
	if union.Count("orange") != 1 {
		t.Errorf("Expected union count 1 for 'orange', got %d", union.Count("orange"))
	}

	// Test Intersection
	intersection := ms1.Intersection(ms2)
	if intersection.Count("apple") != 2 {
		t.Errorf("Expected intersection count 2 for 'apple', got %d", intersection.Count("apple"))
	}
	if intersection.Contains("banana") || intersection.Contains("orange") {
		t.Error("Intersection should not contain 'banana' or 'orange'")
	}

	// Test Difference
	difference := ms1.Difference(ms2)
	if difference.Count("apple") != 1 {
		t.Errorf("Expected difference count 1 for 'apple', got %d", difference.Count("apple"))
	}
	if difference.Count("banana") != 2 {
		t.Errorf("Expected difference count 2 for 'banana', got %d", difference.Count("banana"))
	}
}

func TestMultiSetSubsetSuperset(t *testing.T) {
	subset := NewMultiSet[string]()
	subset.AddCount("apple", 2)
	subset.Add("banana")

	superset := NewMultiSet[string]()
	superset.AddCount("apple", 3)
	superset.AddCount("banana", 2)
	superset.Add("orange")

	if !subset.IsSubset(superset) {
		t.Error("subset should be a subset of superset")
	}
	if subset.IsSuperset(superset) {
		t.Error("subset should not be a superset of superset")
	}
	if !superset.IsSuperset(subset) {
		t.Error("superset should be a superset of subset")
	}
}

func TestMultiSetFrequencyOperations(t *testing.T) {
	ms := NewMultiSet[string]()
	ms.AddCount("apple", 5)
	ms.AddCount("banana", 3)
	ms.AddCount("orange", 1)
	ms.AddCount("grape", 4)

	// Test MostCommon
	most := ms.MostCommon(2)
	if len(most) != 2 {
		t.Errorf("Expected 2 most common elements, got %d", len(most))
	}
	if most[0] != "apple" {
		t.Errorf("Expected most common element to be 'apple', got %s", most[0])
	}

	// Test LeastCommon
	least := ms.LeastCommon(1)
	if len(least) != 1 {
		t.Errorf("Expected 1 least common element, got %d", len(least))
	}
	if least[0] != "orange" {
		t.Errorf("Expected least common element to be 'orange', got %s", least[0])
	}
}

func TestMultiSetClearAndEmpty(t *testing.T) {
	ms := NewMultiSet[string]()

	if !ms.IsEmpty() {
		t.Error("New MultiSet should be empty")
	}

	ms.Add("apple")
	if ms.IsEmpty() {
		t.Error("MultiSet should not be empty after adding element")
	}

	ms.Clear()
	if !ms.IsEmpty() {
		t.Error("MultiSet should be empty after Clear()")
	}
	if ms.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", ms.Size())
	}
}

func TestMultiSetConversions(t *testing.T) {
	ms := NewMultiSet[string]()
	ms.AddCount("apple", 2)
	ms.Add("banana")

	// Test ToSlice
	slice := ms.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Test ToUniqueSlice
	uniqueSlice := ms.ToUniqueSlice()
	if len(uniqueSlice) != 2 {
		t.Errorf("Expected unique slice length 2, got %d", len(uniqueSlice))
	}

	// Test ToCountMap
	countMap := ms.ToCountMap()
	if countMap["apple"] != 2 {
		t.Errorf("Expected count map value 2 for 'apple', got %d", countMap["apple"])
	}
	if countMap["banana"] != 1 {
		t.Errorf("Expected count map value 1 for 'banana', got %d", countMap["banana"])
	}
}
