package stl

import (
	"fmt"
	"strings"
)

// TrieNode represents a node in a trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	value    interface{} // Optional value associated with the word
}

// Trie represents a prefix tree
type Trie struct {
	root *TrieNode
	size int
}

// NewTrie creates a new empty trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
		size: 0,
	}
}

// NewTrieFromSlice creates a trie from a slice of strings
func NewTrieFromSlice(words []string) *Trie {
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	return trie
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	t.InsertWithValue(word, nil)
}

// InsertWithValue adds a word with an associated value to the trie
func (t *Trie) InsertWithValue(word string, value interface{}) {
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}

	if !current.isEnd {
		t.size++
	}
	current.isEnd = true
	current.value = value
}

// Search checks if a word exists in the trie
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

// SearchWithValue returns the value associated with a word
func (t *Trie) SearchWithValue(word string) (interface{}, bool) {
	node := t.searchNode(word)
	if node != nil && node.isEnd {
		return node.value, true
	}
	return nil, false
}

// searchNode is a helper function that returns the node at the end of the word path
func (t *Trie) searchNode(word string) *TrieNode {
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			return nil
		}
		current = current.children[char]
	}

	return current
}

// StartsWith checks if any word in the trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

// Delete removes a word from the trie
func (t *Trie) Delete(word string) bool {
	return t.deleteRecursive(t.root, word, 0)
}

// deleteRecursive is the recursive helper for Delete
func (t *Trie) deleteRecursive(node *TrieNode, word string, index int) bool {
	if node == nil {
		return false
	}

	if index == len(word) {
		if node.isEnd {
			node.isEnd = false
			node.value = nil
			t.size--
			return len(node.children) == 0
		}
		return false
	}

	char := rune(word[index])
	child := node.children[char]
	if child == nil {
		return false
	}

	shouldDeleteChild := t.deleteRecursive(child, word, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// Size returns the number of words in the trie
func (t *Trie) Size() int {
	return t.size
}

// IsEmpty checks if the trie is empty
func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

// Clear removes all words from the trie
func (t *Trie) Clear() {
	t.root = &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
	t.size = 0
}

// GetAllWords returns all words in the trie
func (t *Trie) GetAllWords() []string {
	var words []string
	t.collectWords(t.root, "", &words)
	return words
}

// collectWords is a helper function to collect all words from the trie
func (t *Trie) collectWords(node *TrieNode, prefix string, words *[]string) {
	if node == nil {
		return
	}

	if node.isEnd {
		*words = append(*words, prefix)
	}

	for char, child := range node.children {
		t.collectWords(child, prefix+string(char), words)
	}
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	var words []string
	node := t.searchNode(prefix)
	if node != nil {
		t.collectWords(node, prefix, &words)
	}
	return words
}

// GetWordsWithPrefixLimit returns up to n words that start with the given prefix
func (t *Trie) GetWordsWithPrefixLimit(prefix string, limit int) []string {
	var words []string
	node := t.searchNode(prefix)
	if node != nil {
		t.collectWordsLimit(node, prefix, &words, limit)
	}
	return words
}

// collectWordsLimit is a helper function to collect words with a limit
func (t *Trie) collectWordsLimit(node *TrieNode, prefix string, words *[]string, limit int) {
	if node == nil || len(*words) >= limit {
		return
	}

	if node.isEnd {
		*words = append(*words, prefix)
	}

	for char, child := range node.children {
		if len(*words) >= limit {
			break
		}
		t.collectWordsLimit(child, prefix+string(char), words, limit)
	}
}

// LongestCommonPrefix returns the longest common prefix of all words in the trie
func (t *Trie) LongestCommonPrefix() string {
	if t.IsEmpty() {
		return ""
	}

	var prefix strings.Builder
	current := t.root

	for len(current.children) == 1 && !current.isEnd {
		// Find the single child
		var char rune
		for c := range current.children {
			char = c
			break
		}
		prefix.WriteRune(char)
		current = current.children[char]
	}

	return prefix.String()
}

// GetWordsByLength returns all words with the specified length
func (t *Trie) GetWordsByLength(length int) []string {
	var words []string
	t.collectWordsByLength(t.root, "", 0, length, &words)
	return words
}

// collectWordsByLength is a helper function to collect words with a specific length
func (t *Trie) collectWordsByLength(node *TrieNode, prefix string, currentLength, targetLength int, words *[]string) {
	if node == nil {
		return
	}

	if currentLength == targetLength {
		if node.isEnd {
			*words = append(*words, prefix)
		}
		return
	}

	for char, child := range node.children {
		t.collectWordsByLength(child, prefix+string(char), currentLength+1, targetLength, words)
	}
}

// GetWordsWithPattern returns all words that match the given pattern
// '?' represents any single character, '*' represents any sequence of characters
func (t *Trie) GetWordsWithPattern(pattern string) []string {
	var words []string
	t.matchPattern(t.root, "", pattern, 0, &words)
	return words
}

// matchPattern is a helper function to match words against a pattern
func (t *Trie) matchPattern(node *TrieNode, prefix, pattern string, index int, words *[]string) {
	if node == nil {
		return
	}

	if index == len(pattern) {
		if node.isEnd {
			*words = append(*words, prefix)
		}
		return
	}

	char := rune(pattern[index])

	if char == '?' {
		// Match any single character
		for c, child := range node.children {
			t.matchPattern(child, prefix+string(c), pattern, index+1, words)
		}
	} else if char == '*' {
		// Match any sequence of characters
		t.matchPattern(node, prefix, pattern, index+1, words) // Skip current position
		for c, child := range node.children {
			t.matchPattern(child, prefix+string(c), pattern, index, words) // Continue matching
		}
	} else {
		// Match exact character
		if child := node.children[char]; child != nil {
			t.matchPattern(child, prefix+string(char), pattern, index+1, words)
		}
	}
}

// EditDistance returns the minimum number of operations to transform one word to another
func (t *Trie) EditDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}

	return dp[m][n]
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetWordsWithinDistance returns all words in the trie within a given edit distance
func (t *Trie) GetWordsWithinDistance(target string, maxDistance int) []string {
	var words []string
	t.getAllWords().ForEach(func(word string) {
		if t.EditDistance(word, target) <= maxDistance {
			words = append(words, word)
		}
	})
	return words
}

// getAllWords returns a set of all words in the trie
func (t *Trie) getAllWords() *Set[string] {
	words := NewSet[string]()
	t.collectWordsToSet(t.root, "", words)
	return words
}

// collectWordsToSet is a helper function to collect words into a set
func (t *Trie) collectWordsToSet(node *TrieNode, prefix string, words *Set[string]) {
	if node == nil {
		return
	}

	if node.isEnd {
		words.Add(prefix)
	}

	for char, child := range node.children {
		t.collectWordsToSet(child, prefix+string(char), words)
	}
}

// Height returns the height of the trie
func (t *Trie) Height() int {
	return t.heightRecursive(t.root)
}

// heightRecursive is the recursive helper for Height
func (t *Trie) heightRecursive(node *TrieNode) int {
	if node == nil {
		return 0
	}

	maxHeight := 0
	for _, child := range node.children {
		height := t.heightRecursive(child)
		if height > maxHeight {
			maxHeight = height
		}
	}

	return 1 + maxHeight
}

// String returns a string representation of the trie
func (t *Trie) String() string {
	words := t.GetAllWords()
	return fmt.Sprintf("Trie%v", words)
}

// ForEach applies a function to each word in the trie
func (t *Trie) ForEach(fn func(string)) {
	t.forEachRecursive(t.root, "", fn)
}

// forEachRecursive is the recursive helper for ForEach
func (t *Trie) forEachRecursive(node *TrieNode, prefix string, fn func(string)) {
	if node == nil {
		return
	}

	if node.isEnd {
		fn(prefix)
	}

	for char, child := range node.children {
		t.forEachRecursive(child, prefix+string(char), fn)
	}
}

// Filter returns a new trie containing words that satisfy the predicate
func (t *Trie) Filter(predicate func(string) bool) *Trie {
	result := NewTrie()
	t.filterRecursive(t.root, "", predicate, result)
	return result
}

// filterRecursive is the recursive helper for Filter
func (t *Trie) filterRecursive(node *TrieNode, prefix string, predicate func(string) bool, result *Trie) {
	if node == nil {
		return
	}

	if node.isEnd && predicate(prefix) {
		result.InsertWithValue(prefix, node.value)
	}

	for char, child := range node.children {
		t.filterRecursive(child, prefix+string(char), predicate, result)
	}
}

// Clone creates a deep copy of the trie
func (t *Trie) Clone() *Trie {
	result := NewTrie()
	t.cloneRecursive(t.root, "", result)
	return result
}

// cloneRecursive is the recursive helper for Clone
func (t *Trie) cloneRecursive(node *TrieNode, prefix string, result *Trie) {
	if node == nil {
		return
	}

	if node.isEnd {
		result.InsertWithValue(prefix, node.value)
	}

	for char, child := range node.children {
		t.cloneRecursive(child, prefix+string(char), result)
	}
}

// Equals checks if two tries contain the same words
func (t *Trie) Equals(other *Trie) bool {
	if t.size != other.size {
		return false
	}

	words1 := t.GetAllWords()
	words2 := other.GetAllWords()

	if len(words1) != len(words2) {
		return false
	}

	// Create sets for comparison
	set1 := NewSetFromSlice(words1)
	set2 := NewSetFromSlice(words2)

	return set1.Equals(set2)
}

// GetWordsWithSuffix returns all words that end with the given suffix
// Note: This is not efficient for large tries as it needs to check all words
func (t *Trie) GetWordsWithSuffix(suffix string) []string {
	var words []string
	t.forEachRecursive(t.root, "", func(word string) {
		if strings.HasSuffix(word, suffix) {
			words = append(words, word)
		}
	})
	return words
}

// GetWordsContaining returns all words that contain the given substring
// Note: This is not efficient for large tries as it needs to check all words
func (t *Trie) GetWordsContaining(substring string) []string {
	var words []string
	t.forEachRecursive(t.root, "", func(word string) {
		if strings.Contains(word, substring) {
			words = append(words, word)
		}
	})
	return words
}
