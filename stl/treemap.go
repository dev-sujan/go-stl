package stl

import (
	"fmt"
)

// TreeMapNode represents a node in a TreeMap
type TreeMapNode[K comparable, V any] struct {
	Key   K
	Value V
	Left  *TreeMapNode[K, V]
	Right *TreeMapNode[K, V]
}

// TreeMap represents an ordered map using a binary search tree
type TreeMap[K comparable, V any] struct {
	root *TreeMapNode[K, V]
	size int
	less func(K, K) bool // Comparator function
}

// NewTreeMap creates a new empty TreeMap with a comparator function
func NewTreeMap[K comparable, V any](less func(K, K) bool) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		root: nil,
		size: 0,
		less: less,
	}
}

// NewTreeMapFromMap creates a TreeMap from a regular map
func NewTreeMapFromMap[K comparable, V any](m map[K]V, less func(K, K) bool) *TreeMap[K, V] {
	tm := NewTreeMap[K, V](less)
	for key, value := range m {
		tm.Put(key, value)
	}
	return tm
}

// Put adds or updates a key-value pair in the TreeMap
func (tm *TreeMap[K, V]) Put(key K, value V) {
	tm.root = tm.putRecursive(tm.root, key, value)
}

// putRecursive is the recursive helper for Put
func (tm *TreeMap[K, V]) putRecursive(node *TreeMapNode[K, V], key K, value V) *TreeMapNode[K, V] {
	if node == nil {
		tm.size++
		return &TreeMapNode[K, V]{
			Key:   key,
			Value: value,
		}
	}

	if tm.less(key, node.Key) {
		node.Left = tm.putRecursive(node.Left, key, value)
	} else if tm.less(node.Key, key) {
		node.Right = tm.putRecursive(node.Right, key, value)
	} else {
		// Key already exists, update value
		node.Value = value
	}

	return node
}

// Get returns the value associated with the given key
func (tm *TreeMap[K, V]) Get(key K) (V, bool) {
	node := tm.getNode(key)
	if node != nil {
		return node.Value, true
	}
	var zero V
	return zero, false
}

// getNode is a helper function that returns the node with the given key
func (tm *TreeMap[K, V]) getNode(key K) *TreeMapNode[K, V] {
	current := tm.root

	for current != nil {
		if tm.less(key, current.Key) {
			current = current.Left
		} else if tm.less(current.Key, key) {
			current = current.Right
		} else {
			return current
		}
	}

	return nil
}

// Remove removes a key-value pair from the TreeMap
func (tm *TreeMap[K, V]) Remove(key K) bool {
	if tm.ContainsKey(key) {
		tm.root = tm.removeRecursive(tm.root, key)
		tm.size--
		return true
	}
	return false
}

// removeRecursive is the recursive helper for Remove
func (tm *TreeMap[K, V]) removeRecursive(node *TreeMapNode[K, V], key K) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	if tm.less(key, node.Key) {
		node.Left = tm.removeRecursive(node.Left, key)
	} else if tm.less(node.Key, key) {
		node.Right = tm.removeRecursive(node.Right, key)
	} else {
		// Node to remove found
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node has two children
		// Find the inorder successor (smallest key in right subtree)
		successor := tm.minNode(node.Right)
		node.Key = successor.Key
		node.Value = successor.Value
		node.Right = tm.removeRecursive(node.Right, successor.Key)
	}

	return node
}

// minNode finds the node with the minimum key in a subtree
func (tm *TreeMap[K, V]) minNode(node *TreeMapNode[K, V]) *TreeMapNode[K, V] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// maxNode finds the node with the maximum key in a subtree
func (tm *TreeMap[K, V]) maxNode(node *TreeMapNode[K, V]) *TreeMapNode[K, V] {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// ContainsKey checks if a key exists in the TreeMap
func (tm *TreeMap[K, V]) ContainsKey(key K) bool {
	return tm.getNode(key) != nil
}

// ContainsValue checks if a value exists in the TreeMap
func (tm *TreeMap[K, V]) ContainsValue(value V) bool {
	return tm.containsValueRecursive(tm.root, value)
}

// containsValueRecursive is the recursive helper for ContainsValue
func (tm *TreeMap[K, V]) containsValueRecursive(node *TreeMapNode[K, V], value V) bool {
	if node == nil {
		return false
	}

	if fmt.Sprintf("%v", node.Value) == fmt.Sprintf("%v", value) {
		return true
	}

	return tm.containsValueRecursive(node.Left, value) || tm.containsValueRecursive(node.Right, value)
}

// Min returns the key-value pair with the minimum key
func (tm *TreeMap[K, V]) Min() (K, V, bool) {
	if tm.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	node := tm.minNode(tm.root)
	return node.Key, node.Value, true
}

// Max returns the key-value pair with the maximum key
func (tm *TreeMap[K, V]) Max() (K, V, bool) {
	if tm.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	node := tm.maxNode(tm.root)
	return node.Key, node.Value, true
}

// Floor returns the largest key less than or equal to the given key
func (tm *TreeMap[K, V]) Floor(key K) (K, V, bool) {
	result := tm.floorRecursive(tm.root, key)
	if result == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return result.Key, result.Value, true
}

// floorRecursive is the recursive helper for Floor
func (tm *TreeMap[K, V]) floorRecursive(node *TreeMapNode[K, V], key K) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	}

	if tm.less(key, node.Key) {
		return tm.floorRecursive(node.Left, key)
	}

	// Key is greater than node.Key
	floor := tm.floorRecursive(node.Right, key)
	if floor != nil {
		return floor
	}
	return node
}

// Ceiling returns the smallest key greater than or equal to the given key
func (tm *TreeMap[K, V]) Ceiling(key K) (K, V, bool) {
	result := tm.ceilingRecursive(tm.root, key)
	if result == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return result.Key, result.Value, true
}

// ceilingRecursive is the recursive helper for Ceiling
func (tm *TreeMap[K, V]) ceilingRecursive(node *TreeMapNode[K, V], key K) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	}

	if tm.less(node.Key, key) {
		return tm.ceilingRecursive(node.Right, key)
	}

	// Key is less than node.Key
	ceiling := tm.ceilingRecursive(node.Left, key)
	if ceiling != nil {
		return ceiling
	}
	return node
}

// Lower returns the largest key strictly less than the given key
func (tm *TreeMap[K, V]) Lower(key K) (K, V, bool) {
	result := tm.lowerRecursive(tm.root, key)
	if result == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return result.Key, result.Value, true
}

// lowerRecursive is the recursive helper for Lower
func (tm *TreeMap[K, V]) lowerRecursive(node *TreeMapNode[K, V], key K) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	if tm.less(node.Key, key) {
		// Current node is less than key, check right subtree
		lower := tm.lowerRecursive(node.Right, key)
		if lower != nil {
			return lower
		}
		return node
	}

	// Current node is greater than or equal to key, check left subtree
	return tm.lowerRecursive(node.Left, key)
}

// Higher returns the smallest key strictly greater than the given key
func (tm *TreeMap[K, V]) Higher(key K) (K, V, bool) {
	result := tm.higherRecursive(tm.root, key)
	if result == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	return result.Key, result.Value, true
}

// higherRecursive is the recursive helper for Higher
func (tm *TreeMap[K, V]) higherRecursive(node *TreeMapNode[K, V], key K) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	if tm.less(key, node.Key) {
		// Current node is greater than key, check left subtree
		higher := tm.higherRecursive(node.Left, key)
		if higher != nil {
			return higher
		}
		return node
	}

	// Current node is less than or equal to key, check right subtree
	return tm.higherRecursive(node.Right, key)
}

// Rank returns the number of keys less than the given key
func (tm *TreeMap[K, V]) Rank(key K) int {
	return tm.rankRecursive(tm.root, key)
}

// rankRecursive is the recursive helper for Rank
func (tm *TreeMap[K, V]) rankRecursive(node *TreeMapNode[K, V], key K) int {
	if node == nil {
		return 0
	}

	if tm.less(key, node.Key) {
		return tm.rankRecursive(node.Left, key)
	} else if tm.less(node.Key, key) {
		return 1 + tm.sizeOf(node.Left) + tm.rankRecursive(node.Right, key)
	} else {
		return tm.sizeOf(node.Left)
	}
}

// Select returns the key-value pair with the given rank
func (tm *TreeMap[K, V]) Select(rank int) (K, V, bool) {
	if rank < 0 || rank >= tm.size {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}
	result := tm.selectRecursive(tm.root, rank)
	return result.Key, result.Value, true
}

// selectRecursive is the recursive helper for Select
func (tm *TreeMap[K, V]) selectRecursive(node *TreeMapNode[K, V], rank int) *TreeMapNode[K, V] {
	if node == nil {
		return nil
	}

	leftSize := tm.sizeOf(node.Left)
	if rank < leftSize {
		return tm.selectRecursive(node.Left, rank)
	} else if rank > leftSize {
		return tm.selectRecursive(node.Right, rank-leftSize-1)
	} else {
		return node
	}
}

// sizeOf returns the size of a subtree
func (tm *TreeMap[K, V]) sizeOf(node *TreeMapNode[K, V]) int {
	if node == nil {
		return 0
	}
	return 1 + tm.sizeOf(node.Left) + tm.sizeOf(node.Right)
}

// Size returns the number of key-value pairs in the TreeMap
func (tm *TreeMap[K, V]) Size() int {
	return tm.size
}

// IsEmpty checks if the TreeMap is empty
func (tm *TreeMap[K, V]) IsEmpty() bool {
	return tm.size == 0
}

// Clear removes all key-value pairs from the TreeMap
func (tm *TreeMap[K, V]) Clear() {
	tm.root = nil
	tm.size = 0
}

// Keys returns all keys in the TreeMap in sorted order
func (tm *TreeMap[K, V]) Keys() []K {
	var keys []K
	tm.inOrderTraversal(tm.root, func(key K, value V) {
		keys = append(keys, key)
	})
	return keys
}

// Values returns all values in the TreeMap in key order
func (tm *TreeMap[K, V]) Values() []V {
	var values []V
	tm.inOrderTraversal(tm.root, func(key K, value V) {
		values = append(values, value)
	})
	return values
}

// Entries returns all key-value pairs in the TreeMap in sorted order
func (tm *TreeMap[K, V]) Entries() []struct {
	Key   K
	Value V
} {
	var entries []struct {
		Key   K
		Value V
	}
	tm.inOrderTraversal(tm.root, func(key K, value V) {
		entries = append(entries, struct {
			Key   K
			Value V
		}{key, value})
	})
	return entries
}

// inOrderTraversal performs an in-order traversal of the tree
func (tm *TreeMap[K, V]) inOrderTraversal(node *TreeMapNode[K, V], fn func(K, V)) {
	if node != nil {
		tm.inOrderTraversal(node.Left, fn)
		fn(node.Key, node.Value)
		tm.inOrderTraversal(node.Right, fn)
	}
}

// ToMap converts the TreeMap to a regular map
func (tm *TreeMap[K, V]) ToMap() map[K]V {
	result := make(map[K]V)
	tm.inOrderTraversal(tm.root, func(key K, value V) {
		result[key] = value
	})
	return result
}

// String returns a string representation of the TreeMap
func (tm *TreeMap[K, V]) String() string {
	return fmt.Sprintf("TreeMap%v", tm.ToMap())
}

// ForEach applies a function to each key-value pair in sorted order
func (tm *TreeMap[K, V]) ForEach(fn func(K, V)) {
	tm.inOrderTraversal(tm.root, fn)
}

// Filter returns a new TreeMap containing entries that satisfy the predicate
func (tm *TreeMap[K, V]) Filter(predicate func(K, V) bool) *TreeMap[K, V] {
	result := NewTreeMap[K, V](tm.less)
	tm.filterRecursive(tm.root, predicate, result)
	return result
}

// filterRecursive is the recursive helper for Filter
func (tm *TreeMap[K, V]) filterRecursive(node *TreeMapNode[K, V], predicate func(K, V) bool, result *TreeMap[K, V]) {
	if node != nil {
		tm.filterRecursive(node.Left, predicate, result)
		if predicate(node.Key, node.Value) {
			result.Put(node.Key, node.Value)
		}
		tm.filterRecursive(node.Right, predicate, result)
	}
}

// Clone creates a deep copy of the TreeMap
func (tm *TreeMap[K, V]) Clone() *TreeMap[K, V] {
	result := NewTreeMap[K, V](tm.less)
	tm.cloneRecursive(tm.root, result)
	return result
}

// cloneRecursive is the recursive helper for Clone
func (tm *TreeMap[K, V]) cloneRecursive(node *TreeMapNode[K, V], result *TreeMap[K, V]) {
	if node != nil {
		tm.cloneRecursive(node.Left, result)
		result.Put(node.Key, node.Value)
		tm.cloneRecursive(node.Right, result)
	}
}

// Equals checks if two TreeMaps contain the same key-value pairs
func (tm *TreeMap[K, V]) Equals(other *TreeMap[K, V]) bool {
	if tm.size != other.size {
		return false
	}

	entries1 := tm.Entries()
	entries2 := other.Entries()

	for i := 0; i < len(entries1); i++ {
		if entries1[i].Key != entries2[i].Key || fmt.Sprintf("%v", entries1[i].Value) != fmt.Sprintf("%v", entries2[i].Value) {
			return false
		}
	}

	return true
}

// Range returns all key-value pairs in the TreeMap between min and max (inclusive)
func (tm *TreeMap[K, V]) Range(min, max K) []struct {
	Key   K
	Value V
} {
	var result []struct {
		Key   K
		Value V
	}
	tm.rangeRecursive(tm.root, min, max, &result)
	return result
}

// rangeRecursive is the recursive helper for Range
func (tm *TreeMap[K, V]) rangeRecursive(node *TreeMapNode[K, V], min, max K, result *[]struct {
	Key   K
	Value V
}) {
	if node == nil {
		return
	}

	// If current node is greater than min, recur for left subtree
	if tm.less(min, node.Key) {
		tm.rangeRecursive(node.Left, min, max, result)
	}

	// If current node is in range, add it
	if !tm.less(node.Key, min) && !tm.less(max, node.Key) {
		*result = append(*result, struct {
			Key   K
			Value V
		}{node.Key, node.Value})
	}

	// If current node is less than max, recur for right subtree
	if tm.less(node.Key, max) {
		tm.rangeRecursive(node.Right, min, max, result)
	}
}

// Height returns the height of the TreeMap
func (tm *TreeMap[K, V]) Height() int {
	return tm.heightRecursive(tm.root)
}

// heightRecursive is the recursive helper for Height
func (tm *TreeMap[K, V]) heightRecursive(node *TreeMapNode[K, V]) int {
	if node == nil {
		return -1
	}
	leftHeight := tm.heightRecursive(node.Left)
	rightHeight := tm.heightRecursive(node.Right)
	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}

// IsBalanced checks if the TreeMap is balanced
func (tm *TreeMap[K, V]) IsBalanced() bool {
	return tm.isBalancedRecursive(tm.root) != -1
}

// isBalancedRecursive is the recursive helper for IsBalanced
func (tm *TreeMap[K, V]) isBalancedRecursive(node *TreeMapNode[K, V]) int {
	if node == nil {
		return 0
	}

	leftHeight := tm.isBalancedRecursive(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := tm.isBalancedRecursive(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}
