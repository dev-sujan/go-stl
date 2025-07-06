package stl

import (
	"testing"
)

func TestTrieBasicOperations(t *testing.T) {
	trie := NewTrie()

	// Test Insert and Search
	words := []string{"apple", "application", "app", "banana", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test Size
	if trie.Size() != len(words) {
		t.Errorf("Expected size %d, got %d", len(words), trie.Size())
	}

	// Test Search
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Trie should contain %s", word)
		}
	}

	// Test non-existent word
	if trie.Search("orange") {
		t.Error("Trie should not contain 'orange'")
	}

	// Test StartsWith
	if !trie.StartsWith("app") {
		t.Error("Trie should have words starting with 'app'")
	}
	if !trie.StartsWith("ba") {
		t.Error("Trie should have words starting with 'ba'")
	}
	if trie.StartsWith("ora") {
		t.Error("Trie should not have words starting with 'ora'")
	}
}

func TestTrieDelete(t *testing.T) {
	trie := NewTrie()

	words := []string{"apple", "application", "app", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Delete existing word
	trie.Delete("app")
	if trie.Search("app") {
		t.Error("Trie should not contain 'app' after deletion")
	}

	// Other words with same prefix should still exist
	if !trie.Search("apple") {
		t.Error("Trie should still contain 'apple' after deleting 'app'")
	}
	if !trie.Search("application") {
		t.Error("Trie should still contain 'application' after deleting 'app'")
	}

	// Delete non-existent word
	trie.Delete("orange")
	if trie.Size() != len(words)-1 {
		t.Errorf("Expected size %d after deletion, got %d", len(words)-1, trie.Size())
	}

	// Delete remaining words
	trie.Delete("apple")
	trie.Delete("application")
	trie.Delete("banana")

	if trie.Size() != 0 {
		t.Errorf("Expected size 0 after deleting all words, got %d", trie.Size())
	}
}

// TestTrieWithValues is skipped as the GetValue method is not implemented.
func TestTrieWithValues(t *testing.T) {
	trie := NewTrie()

	// Insert words with associated values
	type WordData struct {
		Language string
		Category string
	}

	trie.InsertWithValue("apple", WordData{Language: "English", Category: "Fruit"})
	trie.InsertWithValue("banana", WordData{Language: "English", Category: "Fruit"})

	// Get value for existing word
	value, found := trie.SearchWithValue("apple")
	if !found {
		t.Error("Should find value for 'apple'")
	}
	if data, ok := value.(WordData); !ok || data.Category != "Fruit" {
		t.Errorf("Expected Category 'Fruit', got %v", data.Category)
	}

	// Get value for non-existent word
	_, found = trie.SearchWithValue("orange")
	if found {
		t.Error("Should not find value for non-existent word 'orange'")
	}

	// Update value
	trie.InsertWithValue("apple", WordData{Language: "English", Category: "Food"})
	value, _ = trie.SearchWithValue("apple")
	if data, ok := value.(WordData); !ok || data.Category != "Food" {
		t.Errorf("Expected updated Category 'Food', got %v", data.Category)
	}
}

func TestTrieFromSlice(t *testing.T) {
	words := []string{"apple", "application", "app", "banana", "ball"}

	trie := NewTrieFromSlice(words)

	if trie.Size() != len(words) {
		t.Errorf("Expected size %d, got %d", len(words), trie.Size())
	}

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Trie should contain %s", word)
		}
	}
}

func TestTrieIsEmpty(t *testing.T) {
	trie := NewTrie()

	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}

	trie.Insert("apple")
	if trie.IsEmpty() {
		t.Error("Trie with words should not be empty")
	}

	trie.Clear()
	if !trie.IsEmpty() {
		t.Error("Cleared trie should be empty")
	}
}

func TestTrieClear(t *testing.T) {
	trie := NewTrie()

	words := []string{"apple", "application", "app", "banana", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	trie.Clear()
	if !trie.IsEmpty() {
		t.Error("Trie should be empty after Clear()")
	}
	if trie.Size() != 0 {
		t.Errorf("Trie size should be 0 after Clear(), got %d", trie.Size())
	}

	for _, word := range words {
		if trie.Search(word) {
			t.Errorf("Trie should not contain %s after Clear()", word)
		}
	}
}

func TestTrieGetWordsWithPrefix(t *testing.T) {
	trie := NewTrie()

	words := []string{"apple", "application", "app", "banana", "ball", "cat"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test with prefix that has multiple matches
	appleWords := trie.GetWordsWithPrefix("app")
	if len(appleWords) != 3 {
		t.Errorf("Expected 3 words with prefix 'app', got %d", len(appleWords))
	}

	// Test with prefix that has single match
	catWords := trie.GetWordsWithPrefix("cat")
	if len(catWords) != 1 || catWords[0] != "cat" {
		t.Errorf("Expected only 'cat' with prefix 'cat', got %v", catWords)
	}

	// Test with prefix that has no matches
	orangeWords := trie.GetWordsWithPrefix("orange")
	if len(orangeWords) != 0 {
		t.Errorf("Expected no words with prefix 'orange', got %v", orangeWords)
	}

	// Test with empty prefix (should return all words)
	allWords := trie.GetWordsWithPrefix("")
	if len(allWords) != len(words) {
		t.Errorf("Expected %d words with empty prefix, got %d", len(words), len(allWords))
	}
}

func TestTrieGetWordsWithPattern(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "cart", "dog", "dart"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test with ? pattern (single character wildcard)
	matches := trie.GetWordsWithPattern("c?t")
	if len(matches) != 1 || matches[0] != "cat" {
		t.Errorf("Expected [cat] for pattern 'c?t', got %v", matches)
	}

	// Test with * pattern (multiple character wildcard)
	matches = trie.GetWordsWithPattern("c*t")
	if len(matches) != 2 || !contains(matches, "cat") || !contains(matches, "cart") {
		t.Errorf("Expected [cat cart] for pattern 'c*t', got %v", matches)
	}

	// Test with combination of ? and *
	matches = trie.GetWordsWithPattern("?a*t")
	if len(matches) != 3 || !contains(matches, "cart") || !contains(matches, "dart") || !contains(matches, "cat") {
		t.Errorf("Expected [cart dart cat] for pattern '?a*t', got %v", matches)
	}

	// Test with no matches
	matches = trie.GetWordsWithPattern("x*")
	if len(matches) != 0 {
		t.Errorf("Expected no matches for pattern 'x*', got %v", matches)
	}
}

func TestTrieEditDistance(t *testing.T) {
	trie := NewTrie()

	// Test exact match (distance 0)
	if dist := trie.EditDistance("cat", "cat"); dist != 0 {
		t.Errorf("Edit distance between same words should be 0, got %d", dist)
	}

	// Test single character difference
	if dist := trie.EditDistance("cat", "bat"); dist != 1 {
		t.Errorf("Edit distance for single char difference should be 1, got %d", dist)
	}

	// Test length difference
	if dist := trie.EditDistance("cat", "cats"); dist != 1 {
		t.Errorf("Edit distance for length difference should be 1, got %d", dist)
	}

	// Test completely different words
	if dist := trie.EditDistance("cat", "dog"); dist != 3 {
		t.Errorf("Edit distance between 'cat' and 'dog' should be 3, got %d", dist)
	}
}

func TestTrieGetWordsWithPrefixLimit(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "application", "app", "banana", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test with limit less than total matches
	matches := trie.GetWordsWithPrefixLimit("app", 2)
	if len(matches) != 2 {
		t.Errorf("Expected exactly 2 matches with prefix 'app', got %d", len(matches))
	}

	// Test with limit equal to total matches
	matches = trie.GetWordsWithPrefixLimit("ba", 2)
	if len(matches) != 2 || !contains(matches, "banana") || !contains(matches, "ball") {
		t.Errorf("Expected [banana ball] with prefix 'ba', got %v", matches)
	}

	// Test with limit greater than total matches
	matches = trie.GetWordsWithPrefixLimit("ball", 5)
	if len(matches) != 1 || matches[0] != "ball" {
		t.Errorf("Expected [ball] with prefix 'ball', got %v", matches)
	}

	// Test with no matches
	matches = trie.GetWordsWithPrefixLimit("xyz", 10)
	if len(matches) != 0 {
		t.Errorf("Expected no matches with prefix 'xyz', got %v", matches)
	}
}

func TestTrieLongestCommonPrefix(t *testing.T) {
	trie := NewTrie()

	// Test empty trie
	if prefix := trie.LongestCommonPrefix(); prefix != "" {
		t.Errorf("Expected empty prefix for empty trie, got %q", prefix)
	}

	// Test single word
	trie.Insert("hello")
	if prefix := trie.LongestCommonPrefix(); prefix != "hello" {
		t.Errorf("Expected 'hello' as prefix for single word, got %q", prefix)
	}

	// Test multiple words with common prefix
	trie.Insert("help")
	trie.Insert("helmet")
	if prefix := trie.LongestCommonPrefix(); prefix != "hel" {
		t.Errorf("Expected 'hel' as common prefix, got %q", prefix)
	}

	// Test words with no common prefix
	trie.Clear()
	trie.Insert("cat")
	trie.Insert("dog")
	if prefix := trie.LongestCommonPrefix(); prefix != "" {
		t.Errorf("Expected empty prefix for words with no common prefix, got %q", prefix)
	}
}

// Helper function for testing string slice containment.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
