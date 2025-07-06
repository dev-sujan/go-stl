package stl

import (
	"fmt"
	"math"
)

// BSTNode represents a node in a binary search tree.
type BSTNode[T comparable] struct {
	Value T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

// BST represents a binary search tree.
type BST[T comparable] struct {
	Root *BSTNode[T]
	Less func(T, T) bool
	Size int
}

// NewBST creates a new empty binary search tree with a comparator function.
func NewBST[T comparable](less func(T, T) bool) *BST[T] {
	return &BST[T]{
		Root: nil,
		Size: 0,
		Less: less,
	}
}

// NewBSTFromSlice creates a BST from a slice.
func NewBSTFromSlice[T comparable](slice []T, less func(T, T) bool) *BST[T] {
	bst := NewBST[T](less)
	for _, item := range slice {
		bst.Insert(item)
	}
	return bst
}

// Insert adds a value to the BST.
func (bst *BST[T]) Insert(value T) {
	bst.Root = bst.insertRecursive(bst.Root, value)
}

// insertRecursive is the recursive helper for Insert.
func (bst *BST[T]) insertRecursive(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil {
		bst.Size++
		return &BSTNode[T]{Value: value}
	}

	if bst.Less(value, node.Value) {
		node.Left = bst.insertRecursive(node.Left, value)
	} else if bst.Less(node.Value, value) {
		node.Right = bst.insertRecursive(node.Right, value)
	}
	// If value equals node.Value, do nothing (no duplicates)

	return node
}

// Search checks if a value exists in the BST.
func (bst *BST[T]) Search(value T) bool {
	return bst.searchRecursive(bst.Root, value) != nil
}

// searchRecursive is the recursive helper for Search.
func (bst *BST[T]) searchRecursive(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil || node.Value == value {
		return node
	}

	if bst.Less(value, node.Value) {
		return bst.searchRecursive(node.Left, value)
	}
	return bst.searchRecursive(node.Right, value)
}

// Delete removes a value from the BST.
func (bst *BST[T]) Delete(value T) bool {
	if bst.Search(value) {
		bst.Root = bst.deleteRecursive(bst.Root, value)
		bst.Size--
		return true
	}
	return false
}

// deleteRecursive is the recursive helper for Delete.
func (bst *BST[T]) deleteRecursive(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil {
		return nil
	}

	switch {
	case bst.Less(value, node.Value):
		node.Left = bst.deleteRecursive(node.Left, value)
	case bst.Less(node.Value, value):
		node.Right = bst.deleteRecursive(node.Right, value)
	default:
		// Node to delete found
		switch {
		case node.Left == nil:
			return node.Right
		case node.Right == nil:
			return node.Left
		default:
			// Node has two children
			// Find the inorder successor (smallest value in right subtree)
			node.Value = bst.findMinValue(node.Right)
			node.Right = bst.deleteRecursive(node.Right, node.Value)
		}
	}

	return node
}

// findMinValue finds the minimum value in a subtree.
func (bst *BST[T]) findMinValue(node *BSTNode[T]) T {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// Min returns the minimum value in the BST.
func (bst *BST[T]) Min() (T, bool) {
	if bst.IsEmpty() {
		var zero T
		return zero, false
	}
	return bst.findMinValue(bst.Root), true
}

// Max returns the maximum value in the BST.
func (bst *BST[T]) Max() (T, bool) {
	if bst.IsEmpty() {
		var zero T
		return zero, false
	}
	return bst.findMaxValue(bst.Root), true
}

// findMaxValue finds the maximum value in a subtree.
func (bst *BST[T]) findMaxValue(node *BSTNode[T]) T {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current.Value
}

// Floor returns the largest value less than or equal to the given value.
func (bst *BST[T]) Floor(value T) (T, bool) {
	result := bst.floorRecursive(bst.Root, value)
	if result == nil {
		var zero T
		return zero, false
	}
	return result.Value, true
}

// floorRecursive is the recursive helper for Floor.
func (bst *BST[T]) floorRecursive(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if bst.Less(value, node.Value) {
		return bst.floorRecursive(node.Left, value)
	}

	// Value is greater than node.Value
	floor := bst.floorRecursive(node.Right, value)
	if floor != nil {
		return floor
	}
	return node
}

// Ceiling returns the smallest value greater than or equal to the given value.
func (bst *BST[T]) Ceiling(value T) (T, bool) {
	result := bst.ceilingRecursive(bst.Root, value)
	if result == nil {
		var zero T
		return zero, false
	}
	return result.Value, true
}

// ceilingRecursive is the recursive helper for Ceiling.
func (bst *BST[T]) ceilingRecursive(node *BSTNode[T], value T) *BSTNode[T] {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if bst.Less(node.Value, value) {
		return bst.ceilingRecursive(node.Right, value)
	}

	// Value is less than node.Value
	ceiling := bst.ceilingRecursive(node.Left, value)
	if ceiling != nil {
		return ceiling
	}
	return node
}

// Rank returns the number of values less than the given value.
func (bst *BST[T]) Rank(value T) int {
	return bst.rankRecursive(bst.Root, value)
}

// rankRecursive is the recursive helper for Rank.
func (bst *BST[T]) rankRecursive(node *BSTNode[T], value T) int {
	if node == nil {
		return 0
	}

	switch {
	case bst.Less(value, node.Value):
		return bst.rankRecursive(node.Left, value)
	case bst.Less(node.Value, value):
		return 1 + bst.sizeOf(node.Left) + bst.rankRecursive(node.Right, value)
	default:
		return bst.sizeOf(node.Left)
	}
}

// Select returns the value with the given rank.
func (bst *BST[T]) Select(rank int) (T, bool) {
	if rank < 0 || rank >= bst.Size {
		var zero T
		return zero, false
	}
	result := bst.selectRecursive(bst.Root, rank)
	return result.Value, true
}

// selectRecursive is the recursive helper for Select.
func (bst *BST[T]) selectRecursive(node *BSTNode[T], rank int) *BSTNode[T] {
	if node == nil {
		return nil
	}

	leftSize := bst.sizeOf(node.Left)
	switch {
	case rank < leftSize:
		return bst.selectRecursive(node.Left, rank)
	case rank > leftSize:
		return bst.selectRecursive(node.Right, rank-leftSize-1)
	default:
		return node
	}
}

// sizeOf returns the size of a subtree.
func (bst *BST[T]) sizeOf(node *BSTNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + bst.sizeOf(node.Left) + bst.sizeOf(node.Right)
}

// IsEmpty checks if the BST is empty.
func (bst *BST[T]) IsEmpty() bool {
	return bst.Size == 0
}

// Clear removes all elements from the BST.
func (bst *BST[T]) Clear() {
	bst.Root = nil
	bst.Size = 0
}

// Height returns the height of the BST.
func (bst *BST[T]) Height() int {
	return bst.heightRecursive(bst.Root)
}

// heightRecursive is the recursive helper for Height.
func (bst *BST[T]) heightRecursive(node *BSTNode[T]) int {
	if node == nil {
		return -1
	}
	return 1 + int(math.Max(float64(bst.heightRecursive(node.Left)), float64(bst.heightRecursive(node.Right))))
}

// IsBalanced checks if the BST is balanced (height difference between left and right subtrees <= 1).
func (bst *BST[T]) IsBalanced() bool {
	return bst.isBalancedRecursive(bst.Root) != -1
}

// isBalancedRecursive is the recursive helper for IsBalanced.
func (bst *BST[T]) isBalancedRecursive(node *BSTNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := bst.isBalancedRecursive(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := bst.isBalancedRecursive(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// InOrder returns the BST elements in in-order traversal.
func (bst *BST[T]) InOrder() []T {
	var result []T
	bst.inOrderRecursive(bst.Root, &result)
	return result
}

// inOrderRecursive is the recursive helper for InOrder.
func (bst *BST[T]) inOrderRecursive(node *BSTNode[T], result *[]T) {
	if node != nil {
		bst.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrderRecursive(node.Right, result)
	}
}

// PreOrder returns the BST elements in pre-order traversal.
func (bst *BST[T]) PreOrder() []T {
	var result []T
	bst.preOrderRecursive(bst.Root, &result)
	return result
}

// preOrderRecursive is the recursive helper for PreOrder.
func (bst *BST[T]) preOrderRecursive(node *BSTNode[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Value)
		bst.preOrderRecursive(node.Left, result)
		bst.preOrderRecursive(node.Right, result)
	}
}

// PostOrder returns the BST elements in post-order traversal.
func (bst *BST[T]) PostOrder() []T {
	var result []T
	bst.postOrderRecursive(bst.Root, &result)
	return result
}

// postOrderRecursive is the recursive helper for PostOrder.
func (bst *BST[T]) postOrderRecursive(node *BSTNode[T], result *[]T) {
	if node != nil {
		bst.postOrderRecursive(node.Left, result)
		bst.postOrderRecursive(node.Right, result)
		*result = append(*result, node.Value)
	}
}

// LevelOrder returns the BST elements in level-order traversal (breadth-first).
func (bst *BST[T]) LevelOrder() []T {
	var result []T
	if bst.Root == nil {
		return result
	}

	queue := []*BSTNode[T]{bst.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

// String returns a string representation of the BST.
func (bst *BST[T]) String() string {
	return fmt.Sprintf("BST%v", bst.InOrder())
}

// ForEach applies a function to each element in in-order traversal.
func (bst *BST[T]) ForEach(fn func(T)) {
	bst.forEachRecursive(bst.Root, fn)
}

// forEachRecursive is the recursive helper for ForEach.
func (bst *BST[T]) forEachRecursive(node *BSTNode[T], fn func(T)) {
	if node != nil {
		bst.forEachRecursive(node.Left, fn)
		fn(node.Value)
		bst.forEachRecursive(node.Right, fn)
	}
}

// Filter returns a new BST containing elements that satisfy the predicate.
func (bst *BST[T]) Filter(predicate func(T) bool) *BST[T] {
	result := NewBST[T](bst.Less)
	bst.filterRecursive(bst.Root, predicate, result)
	return result
}

// filterRecursive is the recursive helper for Filter.
func (bst *BST[T]) filterRecursive(node *BSTNode[T], predicate func(T) bool, result *BST[T]) {
	if node != nil {
		bst.filterRecursive(node.Left, predicate, result)
		if predicate(node.Value) {
			result.Insert(node.Value)
		}
		bst.filterRecursive(node.Right, predicate, result)
	}
}

// Clone creates a deep copy of the BST.
func (bst *BST[T]) Clone() *BST[T] {
	result := NewBST[T](bst.Less)
	bst.cloneRecursive(bst.Root, result)
	return result
}

// cloneRecursive is the recursive helper for Clone.
func (bst *BST[T]) cloneRecursive(node *BSTNode[T], result *BST[T]) {
	if node != nil {
		bst.cloneRecursive(node.Left, result)
		result.Insert(node.Value)
		bst.cloneRecursive(node.Right, result)
	}
}

// Equals checks if two BSTs contain the same elements.
func (bst *BST[T]) Equals(other *BST[T]) bool {
	if bst.Size != other.Size {
		return false
	}

	values1 := bst.InOrder()
	values2 := other.InOrder()

	for i := 0; i < len(values1); i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true
}

// Range returns all values in the BST between min and max (inclusive).
func (bst *BST[T]) Range(min, max T) []T {
	var result []T
	bst.rangeRecursive(bst.Root, min, max, &result)
	return result
}

// rangeRecursive is the recursive helper for Range.
func (bst *BST[T]) rangeRecursive(node *BSTNode[T], min, max T, result *[]T) {
	if node == nil {
		return
	}

	// If current node is greater than min, recur for left subtree
	if bst.Less(min, node.Value) {
		bst.rangeRecursive(node.Left, min, max, result)
	}

	// If current node is in range, add it
	if !bst.Less(node.Value, min) && !bst.Less(max, node.Value) {
		*result = append(*result, node.Value)
	}

	// If current node is less than max, recur for right subtree
	if bst.Less(node.Value, max) {
		bst.rangeRecursive(node.Right, min, max, result)
	}
}

// Successor returns the successor of the given value.
func (bst *BST[T]) Successor(value T) (T, bool) {
	var successor *BSTNode[T]
	current := bst.Root

	for current != nil {
		switch {
		case bst.Less(current.Value, value):
			current = current.Right
		case bst.Less(value, current.Value):
			successor = current
			current = current.Left
		default:
			// Found the value
			if current.Right != nil {
				// Successor is the minimum value in right subtree
				return bst.findMinValue(current.Right), true
			}
			current = nil // Exit loop
		}
	}

	if successor != nil {
		return successor.Value, true
	}
	var zero T
	return zero, false
}

// Predecessor returns the predecessor of the given value.
func (bst *BST[T]) Predecessor(value T) (T, bool) {
	var predecessor *BSTNode[T]
	current := bst.Root

	for current != nil {
		switch {
		case bst.Less(value, current.Value):
			current = current.Left
		case bst.Less(current.Value, value):
			predecessor = current
			current = current.Right
		default:
			// Found the value
			if current.Left != nil {
				// Predecessor is the maximum value in left subtree
				return bst.findMaxValue(current.Left), true
			}
			current = nil // Exit loop
		}
	}

	if predecessor != nil {
		return predecessor.Value, true
	}
	var zero T
	return zero, false
}
