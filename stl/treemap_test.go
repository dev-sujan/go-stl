package stl

import (
	"testing"
)

const (
	testValueOne   = "one"
	testValueTwo   = "two"
	testValueThree = "three"
	testValueFour  = "four"
	testValueFive  = "five"
	testValueSeven = "seven"
	testValueNine  = "nine"
)

func lessInt(a, b int) bool {
	return a < b
}

func TestTreeMapBasicOperations(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Test Put and Get
	tm.Put(1, testValueOne)
	tm.Put(2, testValueTwo)
	tm.Put(3, testValueThree)

	if val, found := tm.Get(1); !found || val != testValueOne {
		t.Errorf("Expected 'one' for key 1, got %v", val)
	}

	if val, found := tm.Get(2); !found || val != testValueTwo {
		t.Errorf("Expected 'two' for key 2, got %v", val)
	}

	// Test non-existent key
	if _, found := tm.Get(4); found {
		t.Error("Should not find value for non-existent key 4")
	}

	// Test size
	if size := tm.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test updating existing key
	tm.Put(1, "one updated")
	if val, _ := tm.Get(1); val != "one updated" {
		t.Errorf("Expected 'one updated' after update, got %v", val)
	}
}

func TestTreeMapRemove(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Setup
	tm.Put(1, testValueOne)
	tm.Put(2, testValueTwo)
	tm.Put(3, testValueThree)

	// Test removing existing key
	if removed := tm.Remove(2); !removed {
		t.Error("Remove should return true for existing key")
	}
	if _, found := tm.Get(2); found {
		t.Error("Should not find value after removal")
	}
	if size := tm.Size(); size != 2 {
		t.Errorf("Expected size 2 after removal, got %d", size)
	}

	// Test removing non-existent key
	if removed := tm.Remove(4); removed {
		t.Error("Remove should return false for non-existent key")
	}

	// Test removing remaining keys
	tm.Remove(1)
	tm.Remove(3)
	if !tm.IsEmpty() {
		t.Error("TreeMap should be empty after removing all keys")
	}
}

func TestTreeMapFromMap(t *testing.T) {
	source := map[int]string{
		1: testValueOne,
		2: testValueTwo,
		3: testValueThree,
	}

	tm := NewTreeMapFromMap(source, lessInt)

	if tm.Size() != len(source) {
		t.Errorf("Expected size %d, got %d", len(source), tm.Size())
	}

	for k, v := range source {
		if val, found := tm.Get(k); !found || val != v {
			t.Errorf("Expected %v for key %v, got %v", v, k, val)
		}
	}
}

func TestTreeMapRangeOperations(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Setup
	values := map[int]string{
		1: testValueOne,
		3: testValueThree,
		5: testValueFive,
		7: testValueSeven,
		9: testValueNine,
	}
	for k, v := range values {
		tm.Put(k, v)
	}

	// Test Floor
	if k, v, found := tm.Floor(4); !found || k != 3 || v != testValueThree {
		t.Errorf("Floor(4) should return (3, 'three'), got (%v, %v)", k, v)
	}
	if k, v, found := tm.Floor(5); !found || k != 5 || v != testValueFive {
		t.Errorf("Floor(5) should return (5, 'five'), got (%v, %v)", k, v)
	}
	if _, _, found := tm.Floor(0); found {
		t.Error("Floor(0) should return not found")
	}

	// Test Ceiling
	if k, v, found := tm.Ceiling(4); !found || k != 5 || v != testValueFive {
		t.Errorf("Ceiling(4) should return (5, 'five'), got (%v, %v)", k, v)
	}
	if k, v, found := tm.Ceiling(5); !found || k != 5 || v != testValueFive {
		t.Errorf("Ceiling(5) should return (5, 'five'), got (%v, %v)", k, v)
	}
	if _, _, found := tm.Ceiling(10); found {
		t.Error("Ceiling(10) should return not found")
	}

	// Test Lower
	if k, v, found := tm.Lower(5); !found || k != 3 || v != testValueThree {
		t.Errorf("Lower(5) should return (3, 'three'), got (%v, %v)", k, v)
	}
	if _, _, found := tm.Lower(1); found {
		t.Error("Lower(1) should return not found")
	}

	// Test Higher
	if k, v, found := tm.Higher(5); !found || k != 7 || v != testValueSeven {
		t.Errorf("Higher(5) should return (7, 'seven'), got (%v, %v)", k, v)
	}
	if _, _, found := tm.Higher(9); found {
		t.Error("Higher(9) should return not found")
	}
}

func TestTreeMapStatisticalOperations(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Setup ordered key-value pairs
	pairs := []struct {
		key   int
		value string
	}{
		{1, testValueOne},
		{3, testValueThree},
		{5, testValueFive},
		{7, testValueSeven},
		{9, testValueNine},
	}

	for _, pair := range pairs {
		tm.Put(pair.key, pair.value)
	}

	// Test Rank
	if rank := tm.Rank(3); rank != 1 {
		t.Errorf("Expected rank 1 for key 3, got %d", rank)
	}
	if rank := tm.Rank(6); rank != 3 {
		t.Errorf("Expected rank 3 for key 6 (between 5 and 7), got %d", rank)
	}

	// Test Select
	if k, v, found := tm.Select(0); !found || k != 1 || v != testValueOne {
		t.Errorf("Select(0) should return (1, 'one'), got (%v, %v)", k, v)
	}
	if k, v, found := tm.Select(2); !found || k != 5 || v != testValueFive {
		t.Errorf("Select(2) should return (5, 'five'), got (%v, %v)", k, v)
	}
	if _, _, found := tm.Select(5); found {
		t.Error("Select(5) should return not found")
	}

	// Test Min/Max
	if k, v, found := tm.Min(); !found || k != 1 || v != testValueOne {
		t.Errorf("Min() should return (1, 'one'), got (%v, %v)", k, v)
	}
	if k, v, found := tm.Max(); !found || k != 9 || v != testValueNine {
		t.Errorf("Max() should return (9, 'nine'), got (%v, %v)", k, v)
	}
}

func TestTreeMapCollectionOperations(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Setup
	pairs := map[int]string{
		1: testValueOne,
		2: testValueTwo,
		3: testValueThree,
	}
	for k, v := range pairs {
		tm.Put(k, v)
	}

	// Test Keys
	keys := tm.Keys()
	if len(keys) != len(pairs) {
		t.Errorf("Expected %d keys, got %d", len(pairs), len(keys))
	}
	for i, k := range keys {
		if i > 0 && !lessInt(keys[i-1], k) {
			t.Error("Keys should be in sorted order")
		}
	}

	// Test Values
	values := tm.Values()
	if len(values) != len(pairs) {
		t.Errorf("Expected %d values, got %d", len(pairs), len(values))
	}

	// Test Entries
	entries := tm.Entries()
	if len(entries) != len(pairs) {
		t.Errorf("Expected %d entries, got %d", len(pairs), len(entries))
	}
	for i, entry := range entries {
		if i > 0 && !lessInt(entries[i-1].Key, entry.Key) {
			t.Error("Entries should be in sorted order by key")
		}
		if val, found := pairs[entry.Key]; !found || val != entry.Value {
			t.Errorf("Entry {%v: %v} doesn't match original data", entry.Key, entry.Value)
		}
	}
}

func TestTreeMapAdvancedOperations(t *testing.T) {
	tm := NewTreeMap[int, string](lessInt)

	// Setup
	pairs := map[int]string{
		1: testValueOne,
		2: testValueTwo,
		3: testValueThree,
		4: testValueFour,
		5: testValueFive,
	}
	for k, v := range pairs {
		tm.Put(k, v)
	}

	// Test Filter
	evenOnly := tm.Filter(func(k int, v string) bool {
		return k%2 == 0
	})
	if evenOnly.Size() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", evenOnly.Size())
	}
	if val, found := evenOnly.Get(2); !found || val != testValueTwo {
		t.Error("Filtered map should contain 2: 'two'")
	}
	if val, found := evenOnly.Get(4); !found || val != testValueFour {
		t.Error("Filtered map should contain 4: 'four'")
	}

	// Test Range
	rangeEntries := tm.Range(2, 4)
	if len(rangeEntries) != 3 {
		t.Errorf("Expected 3 entries in range [2,4], got %d", len(rangeEntries))
	}
	for _, entry := range rangeEntries {
		if entry.Key < 2 || entry.Key > 4 {
			t.Errorf("Entry key %d outside range [2,4]", entry.Key)
		}
	}

	// Test Clone
	clone := tm.Clone()
	if !tm.Equals(clone) {
		t.Error("Clone should be equal to original map")
	}
	clone.Put(6, "six")
	if tm.Equals(clone) {
		t.Error("Modified clone should not be equal to original map")
	}

	// Test Height
	if height := tm.Height(); height < 2 {
		t.Errorf("Expected height >= 2 for 5 nodes, got %d", height)
	}
	// Note: Tree may not be balanced after sequential insertions as this is not a self-balancing BST
}
