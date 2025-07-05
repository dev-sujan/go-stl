package stl

import (
	"testing"
)

func TestSetBasicOperations(t *testing.T) {
	set := NewSet[int]()

	// Test Add and Contains
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if !set.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !set.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !set.Contains(3) {
		t.Error("Set should contain 3")
	}
	if set.Contains(4) {
		t.Error("Set should not contain 4")
	}

	// Test Size
	if set.Size() != 3 {
		t.Errorf("Expected size 3, got %d", set.Size())
	}

	// Test Remove
	set.Remove(2)
	if set.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}
	if set.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", set.Size())
	}

	// Test IsEmpty
	if set.IsEmpty() {
		t.Error("Set should not be empty")
	}

	// Test Clear
	set.Clear()
	if !set.IsEmpty() {
		t.Error("Set should be empty after clear")
	}
	if set.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", set.Size())
	}
}

func TestSetOperations(t *testing.T) {
	set1 := NewSet[int]()
	set2 := NewSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	// Test Union
	union := set1.Union(set2)
	expectedUnion := []int{1, 2, 3, 4}
	for _, val := range expectedUnion {
		if !union.Contains(val) {
			t.Errorf("Union should contain %d", val)
		}
	}

	// Test Intersection
	intersection := set1.Intersection(set2)
	expectedIntersection := []int{2, 3}
	for _, val := range expectedIntersection {
		if !intersection.Contains(val) {
			t.Errorf("Intersection should contain %d", val)
		}
	}
	if intersection.Size() != 2 {
		t.Errorf("Expected intersection size 2, got %d", intersection.Size())
	}

	// Test Difference
	difference := set1.Difference(set2)
	if !difference.Contains(1) {
		t.Error("Difference should contain 1")
	}
	if difference.Contains(2) || difference.Contains(3) {
		t.Error("Difference should not contain 2 or 3")
	}

	// Test SymmetricDifference
	symmetricDiff := set1.SymmetricDifference(set2)
	expectedSymmetricDiff := []int{1, 4}
	for _, val := range expectedSymmetricDiff {
		if !symmetricDiff.Contains(val) {
			t.Errorf("Symmetric difference should contain %d", val)
		}
	}
}

func TestSetFromSlice(t *testing.T) {
	slice := []int{1, 2, 2, 3, 3, 3}
	set := NewSetFromSlice(slice)

	if set.Size() != 3 {
		t.Errorf("Expected size 3 (unique elements), got %d", set.Size())
	}

	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("Set should contain all unique elements from slice")
	}
}

func TestSetFunctionalOperations(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	// Test Filter
	evenNumbers := set.Filter(func(x int) bool { return x%2 == 0 })
	if evenNumbers.Size() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", evenNumbers.Size())
	}
	if !evenNumbers.Contains(2) || !evenNumbers.Contains(4) {
		t.Error("Filtered set should contain 2 and 4")
	}

	// Test Any
	hasEven := set.Any(func(x int) bool { return x%2 == 0 })
	if !hasEven {
		t.Error("Set should have even numbers")
	}

	hasNegative := set.Any(func(x int) bool { return x < 0 })
	if hasNegative {
		t.Error("Set should not have negative numbers")
	}

	// Test All
	allPositive := set.All(func(x int) bool { return x > 0 })
	if !allPositive {
		t.Error("All numbers should be positive")
	}

	allEven := set.All(func(x int) bool { return x%2 == 0 })
	if allEven {
		t.Error("Not all numbers should be even")
	}
}

func TestSetClone(t *testing.T) {
	original := NewSet[int]()
	original.Add(1)
	original.Add(2)
	original.Add(3)

	cloned := original.Clone()

	// Test that they are equal
	if !original.Equals(cloned) {
		t.Error("Cloned set should equal original")
	}

	// Test that they are independent
	original.Add(4)
	if cloned.Contains(4) {
		t.Error("Cloned set should not be affected by original changes")
	}

	cloned.Add(5)
	if original.Contains(5) {
		t.Error("Original set should not be affected by cloned changes")
	}
}

func TestSetEquals(t *testing.T) {
	set1 := NewSet[int]()
	set2 := NewSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(3)
	set2.Add(2)
	set2.Add(1)

	// Test equality (order doesn't matter)
	if !set1.Equals(set2) {
		t.Error("Sets should be equal regardless of insertion order")
	}

	// Test inequality
	set2.Add(4)
	if set1.Equals(set2) {
		t.Error("Sets should not be equal when they have different elements")
	}
}

func BenchmarkSetAdd(b *testing.B) {
	set := NewSet[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkSetContains(b *testing.B) {
	set := NewSet[int]()
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i % 1000)
	}
}
