package stl

import (
	"testing"
)

func TestBSTBasicOperations(t *testing.T) {
	// Create BST with integer comparator
	bst := NewBST[int](func(a, b int) bool {
		return a < b
	})

	// Test Insert and Contains
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	// Check size
	if bst.Size != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), bst.Size)
	}

	// Test Search
	for _, v := range values {
		if !bst.Search(v) {
			t.Errorf("BST should contain %d", v)
		}
	}

	// Test non-existent value
	if bst.Search(99) {
		t.Error("BST should not contain 99")
	}

	// Test Min and Max
	min, found := bst.Min()
	if !found || min != 2 {
		t.Errorf("Expected min value 2, got %d, found: %v", min, found)
	}

	max, found := bst.Max()
	if !found || max != 8 {
		t.Errorf("Expected max value 8, got %d, found: %v", max, found)
	}
}

func TestBSTRemove(t *testing.T) {
	bst := NewBST[int](func(a, b int) bool {
		return a < b
	})

	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	// Delete leaf node
	bst.Delete(2)
	if bst.Search(2) {
		t.Error("BST should not contain 2 after deletion")
	}
	if bst.Size != len(values)-1 {
		t.Errorf("Expected size %d after deletion, got %d", len(values)-1, bst.Size)
	}

	// Delete node with one child
	bst.Delete(3)
	if bst.Search(3) {
		t.Error("BST should not contain 3 after deletion")
	}

	// Delete node with two children
	bst.Delete(7)
	if bst.Search(7) {
		t.Error("BST should not contain 7 after deletion")
	}

	// Delete root
	bst.Delete(5)
	if bst.Search(5) {
		t.Error("BST should not contain 5 after deletion")
	}
}

// TestBSTTraversal is skipped as the methods are not yet implemented.
func TestBSTTraversal(t *testing.T) {
	t.Skip("Traversal methods not implemented yet")
}

func TestBSTFromSlice(t *testing.T) {
	values := []int{5, 3, 7, 2, 4, 6, 8}

	bst := NewBSTFromSlice(values, func(a, b int) bool {
		return a < b
	})

	if bst.Size != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), bst.Size)
	}

	for _, v := range values {
		if !bst.Search(v) {
			t.Errorf("BST should contain %d", v)
		}
	}
}

func TestBSTHeight(t *testing.T) {
	bst := NewBST[int](func(a, b int) bool {
		return a < b
	})

	// Empty tree
	if bst.Height() != -1 {
		t.Errorf("Height of empty tree should be -1, got %d", bst.Height())
	}

	// Single node
	bst.Insert(1)
	if bst.Height() != 0 {
		t.Errorf("Height of single node tree should be 0, got %d", bst.Height())
	}

	// Multiple levels
	bst.Insert(2)
	bst.Insert(3)
	if bst.Height() != 2 {
		t.Errorf("Height of tree should be 2, got %d", bst.Height())
	}
}

func TestBSTIsEmpty(t *testing.T) {
	bst := NewBST[int](func(a, b int) bool {
		return a < b
	})

	if !bst.IsEmpty() {
		t.Error("New BST should be empty")
	}

	bst.Insert(1)
	if bst.IsEmpty() {
		t.Error("BST with nodes should not be empty")
	}

	bst.Clear()
	if !bst.IsEmpty() {
		t.Error("Cleared BST should be empty")
	}
}

func TestBSTClear(t *testing.T) {
	bst := NewBST[int](func(a, b int) bool {
		return a < b
	})

	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	bst.Clear()
	if !bst.IsEmpty() {
		t.Error("BST should be empty after Clear()")
	}
	if bst.Size != 0 {
		t.Errorf("BST size should be 0 after Clear(), got %d", bst.Size)
	}
}

func TestBSTWithCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	bst := NewBST[Person](func(a, b Person) bool {
		return a.Age < b.Age
	})

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	for _, p := range people {
		bst.Insert(p)
	}

	if bst.Size != len(people) {
		t.Errorf("Expected size %d, got %d", len(people), bst.Size)
	}

	youngest, found := bst.Min()
	if !found || youngest.Name != "Bob" {
		t.Errorf("Expected youngest to be Bob, got %s", youngest.Name)
	}

	oldest, found := bst.Max()
	if !found || oldest.Name != "Charlie" {
		t.Errorf("Expected oldest to be Charlie, got %s", oldest.Name)
	}
}
