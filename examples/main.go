package main

import (
	"fmt"
	"strings"

	"github.com/dev-sujan/go-stl/stl"
)

func main() {
	fmt.Println("=== Go STL Examples ===")
	fmt.Println()

	// 1. Set Operations
	exampleSet()

	// 2. MultiSet Operations
	exampleMultiSet()

	// 3. MultiMap Operations
	exampleMultiMap()

	// 4. Stack Operations
	exampleStack()

	// 5. Queue Operations
	exampleQueue()

	// 6. Deque Operations
	exampleDeque()

	// 7. Priority Queue Operations
	examplePriorityQueue()

	// 8. Binary Search Tree Operations
	exampleBST()

	// 9. Trie Operations
	exampleTrie()

	// 10. TreeMap Operations
	exampleTreeMap()

	// 11. Graph Operations
	exampleGraph()

	// 12. Advanced Features
	exampleAdvancedFeatures()

	fmt.Println("=== All Examples Completed ===")
}

func exampleSet() {
	fmt.Println("1. Set Operations:")
	fmt.Println("------------------")

	// Create sets
	set1 := stl.NewSet[int]()
	set2 := stl.NewSet[int]()

	// Add elements
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	fmt.Printf("Set 1: %v\n", set1)
	fmt.Printf("Set 2: %v\n", set2)

	// Set operations
	union := set1.Union(set2)
	intersection := set1.Intersection(set2)
	difference := set1.Difference(set2)
	symmetricDiff := set1.SymmetricDifference(set2)

	fmt.Printf("Union: %v\n", union)
	fmt.Printf("Intersection: %v\n", intersection)
	fmt.Printf("Difference (set1 - set2): %v\n", difference)
	fmt.Printf("Symmetric Difference: %v\n", symmetricDiff)

	// Functional operations
	evenNumbers := set1.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers in set1: %v\n", evenNumbers)

	hasEven := set1.Any(func(x int) bool { return x%2 == 0 })
	allPositive := set1.All(func(x int) bool { return x > 0 })
	fmt.Printf("Has even numbers: %v\n", hasEven)
	fmt.Printf("All positive: %v\n", allPositive)

	fmt.Println()
}

func exampleMultiSet() {
	fmt.Println("2. MultiSet Operations:")
	fmt.Println("----------------------")

	ms := stl.NewMultiSet[string]()

	// Add elements (including duplicates)
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	for _, word := range words {
		ms.Add(word)
	}

	fmt.Printf("MultiSet: %v\n", ms)
	fmt.Printf("Count of 'apple': %d\n", ms.Count("apple"))
	fmt.Printf("Count of 'banana': %d\n", ms.Count("banana"))
	fmt.Printf("Count of 'cherry': %d\n", ms.Count("cherry"))

	// Most common elements
	mostCommon := ms.MostCommon(2)
	fmt.Printf("Most common elements: %v\n", mostCommon)

	// Remove elements
	ms.Remove("apple") // removes one occurrence
	fmt.Printf("After removing one 'apple': %v\n", ms)

	ms.RemoveAll("banana") // removes all occurrences
	fmt.Printf("After removing all 'banana': %v\n", ms)

	// Functional operations
	fruits := ms.Filter(func(word string) bool { return len(word) > 5 })
	fmt.Printf("Words with length > 5: %v\n", fruits)

	fmt.Println()
}

func exampleMultiMap() {
	fmt.Println("3. MultiMap Operations:")
	fmt.Println("----------------------")

	mm := stl.NewMultiMap[string, int]()

	// Add key-value pairs
	mm.Put("fruit", 1)
	mm.Put("fruit", 2)
	mm.Put("vegetable", 3)
	mm.Put("fruit", 4)
	mm.Put("meat", 5)

	fmt.Printf("MultiMap: %v\n", mm)
	fmt.Printf("Values for 'fruit': %v\n", mm.Get("fruit"))
	fmt.Printf("Values for 'vegetable': %v\n", mm.Get("vegetable"))

	// Get first and last values
	if first, exists := mm.GetFirst("fruit"); exists {
		fmt.Printf("First value for 'fruit': %d\n", first)
	}
	if last, exists := mm.GetLast("fruit"); exists {
		fmt.Printf("Last value for 'fruit': %d\n", last)
	}

	// Remove specific value
	mm.Remove("fruit", 2)
	fmt.Printf("After removing fruit->2: %v\n", mm)

	// Remove all values for a key
	mm.RemoveAll("meat")
	fmt.Printf("After removing all 'meat': %v\n", mm)

	// Get all keys and values
	fmt.Printf("All keys: %v\n", mm.Keys())
	fmt.Printf("All values: %v\n", mm.Values())

	// Filter operations
	largeValues := mm.Filter(func(key string, value int) bool { return value > 2 })
	fmt.Printf("Entries with value > 2: %v\n", largeValues)

	fmt.Println()
}

func exampleStack() {
	fmt.Println("4. Stack Operations:")
	fmt.Println("---------------------")
	stack := stl.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Printf("Stack: %v\n", stack)
	fmt.Printf("Size: %d\n", stack.Size())

	if top, exists := stack.Peek(); exists {
		fmt.Printf("Top element: %d\n", top)
	}

	if popped, exists := stack.Pop(); exists {
		fmt.Printf("Popped: %d\n", popped)
	}
	fmt.Printf("After pop: %v\n", stack)

	stack.PushAll([]int{6, 7, 8})
	fmt.Printf("After pushing [6,7,8]: %v\n", stack)

	// Stack operations
	evenStack := stack.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evenStack)

	squaredStack := stack.Map(func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squaredStack)

	stack.Reverse()
	fmt.Printf("Reversed: %v\n", stack)

	fmt.Println()
}

func exampleQueue() {
	fmt.Println("5. Queue Operations:")
	fmt.Println("---------------------")
	queue := stl.NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	queue.Enqueue("fourth")
	fmt.Printf("Queue: %v\n", queue)
	fmt.Printf("Size: %d\n", queue.Size())

	if front, exists := queue.Peek(); exists {
		fmt.Printf("Front: %s\n", front)
	}
	if back, exists := queue.PeekBack(); exists {
		fmt.Printf("Back: %s\n", back)
	}

	if dequeued, exists := queue.Dequeue(); exists {
		fmt.Printf("Dequeued: %s\n", dequeued)
	}
	fmt.Printf("After dequeue: %v\n", queue)

	queue.EnqueueAll([]string{"fifth", "sixth"})
	fmt.Printf("After enqueuing [fifth,sixth]: %v\n", queue)

	// Queue operations
	longQueue := queue.Filter(func(s string) bool { return len(s) > 4 })
	fmt.Printf("Long words: %v\n", longQueue)

	upperQueue := queue.Map(func(s string) string { return strings.ToUpper(s) })
	fmt.Printf("Uppercase: %v\n", upperQueue)

	fmt.Println()
}

func exampleDeque() {
	fmt.Println("6. Deque Operations:")
	fmt.Println("-------------------")

	deque := stl.NewDeque[int](8)

	// Add elements from both ends
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)
	deque.PushBack(3)
	deque.PushFront(-1)

	fmt.Printf("Deque: %v\n", deque)
	fmt.Printf("Size: %d\n", deque.Size())

	// Access front and back
	if front, exists := deque.Front(); exists {
		fmt.Printf("Front: %d\n", front)
	}
	if back, exists := deque.Back(); exists {
		fmt.Printf("Back: %d\n", back)
	}

	// Remove elements
	if val, exists := deque.PopFront(); exists {
		fmt.Printf("Popped front: %d\n", val)
	}
	if val, exists := deque.PopBack(); exists {
		fmt.Printf("Popped back: %d\n", val)
	}

	fmt.Printf("After popping: %v\n", deque)

	// Random access
	if val, exists := deque.At(1); exists {
		fmt.Printf("Element at index 1: %d\n", val)
	}

	// Insert at specific position
	deque.Insert(1, 10)
	fmt.Printf("After inserting 10 at index 1: %v\n", deque)

	// Rotate operations
	deque.RotateLeft(1)
	fmt.Printf("After rotating left by 1: %v\n", deque)

	deque.RotateRight(1)
	fmt.Printf("After rotating right by 1: %v\n", deque)

	// Functional operations
	positiveNumbers := deque.Filter(func(x int) bool { return x > 0 })
	fmt.Printf("Positive numbers: %v\n", positiveNumbers)

	fmt.Println()
}

func examplePriorityQueue() {
	fmt.Println("7. Priority Queue Operations:")
	fmt.Println("------------------------------")
	// Min-heap priority queue
	pq := stl.NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.Enqueue(5)
	pq.Enqueue(2)
	pq.Enqueue(8)
	pq.Enqueue(1)
	pq.Enqueue(9)
	pq.Enqueue(3)
	fmt.Printf("Priority Queue: %v\n", pq)

	fmt.Println("Dequeuing in priority order:")
	for !pq.IsEmpty() {
		if item, exists := pq.Dequeue(); exists {
			fmt.Printf("  %d", item)
		}
	}
	fmt.Println()

	// Max-heap priority queue
	maxPQ := stl.NewPriorityQueue[int](func(a, b int) bool { return a > b })
	maxPQ.Enqueue(5)
	maxPQ.Enqueue(2)
	maxPQ.Enqueue(8)
	maxPQ.Enqueue(1)
	maxPQ.Enqueue(9)
	maxPQ.Enqueue(3)
	fmt.Printf("Max Priority Queue: %v\n", maxPQ)

	fmt.Println("Dequeuing in priority order (max first):")
	for !maxPQ.IsEmpty() {
		if item, exists := maxPQ.Dequeue(); exists {
			fmt.Printf("  %d", item)
		}
	}
	fmt.Println()
}

func exampleBST() {
	fmt.Println("8. Binary Search Tree Operations:")
	fmt.Println("--------------------------------")

	// Create BST with integer comparator
	bst := stl.NewBST[int](func(a, b int) bool { return a < b })

	// Insert elements
	elements := []int{5, 3, 7, 1, 9, 4, 6, 2, 8}
	for _, elem := range elements {
		bst.Insert(elem)
	}

	fmt.Printf("BST InOrder: %v\n", bst.InOrder())
	fmt.Printf("BST PreOrder: %v\n", bst.PreOrder())
	fmt.Printf("BST PostOrder: %v\n", bst.PostOrder())
	fmt.Printf("BST LevelOrder: %v\n", bst.LevelOrder())

	// Min and Max
	if min, exists := bst.Min(); exists {
		fmt.Printf("Minimum: %d\n", min)
	}
	if max, exists := bst.Max(); exists {
		fmt.Printf("Maximum: %d\n", max)
	}

	// Floor and Ceiling
	if floor, exists := bst.Floor(4); exists {
		fmt.Printf("Floor of 4: %d\n", floor)
	}
	if ceiling, exists := bst.Ceiling(4); exists {
		fmt.Printf("Ceiling of 4: %d\n", ceiling)
	}

	// Search operations
	fmt.Printf("Contains 5: %v\n", bst.Search(5))
	fmt.Printf("Contains 10: %v\n", bst.Search(10))

	// Delete operation
	bst.Delete(5)
	fmt.Printf("After deleting 5: %v\n", bst.InOrder())

	// Range queries
	rangeResult := bst.Range(3, 7)
	fmt.Printf("Range [3, 7]: %v\n", rangeResult)

	// Successor and Predecessor
	if succ, exists := bst.Successor(4); exists {
		fmt.Printf("Successor of 4: %d\n", succ)
	}
	if pred, exists := bst.Predecessor(6); exists {
		fmt.Printf("Predecessor of 6: %d\n", pred)
	}

	// Functional operations
	evenNumbers := bst.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evenNumbers.InOrder())

	fmt.Println()
}

func exampleTrie() {
	fmt.Println("9. Trie Operations:")
	fmt.Println("------------------")

	trie := stl.NewTrie()

	// Insert words
	words := []string{"hello", "world", "help", "hero", "heroic", "heroism", "cat", "car", "card"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Printf("All words: %v\n", trie.GetAllWords())

	// Search operations
	fmt.Printf("Contains 'hello': %v\n", trie.Search("hello"))
	fmt.Printf("Contains 'help': %v\n", trie.Search("help"))
	fmt.Printf("Contains 'hero': %v\n", trie.Search("hero"))
	fmt.Printf("Contains 'heroic': %v\n", trie.Search("heroic"))

	// Prefix operations
	fmt.Printf("Starts with 'he': %v\n", trie.StartsWith("he"))
	fmt.Printf("Words with prefix 'he': %v\n", trie.GetWordsWithPrefix("he"))
	fmt.Printf("Words with prefix 'car': %v\n", trie.GetWordsWithPrefix("car"))

	// Pattern matching
	fmt.Printf("Words matching 'h?llo': %v\n", trie.GetWordsWithPattern("h?llo"))
	fmt.Printf("Words matching 'h*': %v\n", trie.GetWordsWithPattern("h*"))

	// Word length queries
	fmt.Printf("Words with length 4: %v\n", trie.GetWordsByLength(4))
	fmt.Printf("Words with length 5: %v\n", trie.GetWordsByLength(5))

	// Edit distance
	distance := trie.EditDistance("hello", "helo")
	fmt.Printf("Edit distance between 'hello' and 'helo': %d\n", distance)

	// Words within edit distance
	similar := trie.GetWordsWithinDistance("hello", 2)
	fmt.Printf("Words within edit distance 2 of 'hello': %v\n", similar)

	// Longest common prefix
	commonPrefix := trie.LongestCommonPrefix()
	fmt.Printf("Longest common prefix: %s\n", commonPrefix)

	// Delete operation
	trie.Delete("hero")
	fmt.Printf("After deleting 'hero': %v\n", trie.GetAllWords())

	// Functional operations
	longWords := trie.Filter(func(word string) bool { return len(word) > 4 })
	fmt.Printf("Words with length > 4: %v\n", longWords.GetAllWords())

	fmt.Println()
}

func exampleTreeMap() {
	fmt.Println("10. TreeMap Operations:")
	fmt.Println("---------------------")

	// Create TreeMap with string comparator
	treeMap := stl.NewTreeMap[string, int](func(a, b string) bool { return a < b })

	// Add key-value pairs
	pairs := map[string]int{
		"apple":      1,
		"banana":     2,
		"cherry":     3,
		"date":       4,
		"elderberry": 5,
	}

	for key, value := range pairs {
		treeMap.Put(key, value)
	}

	fmt.Printf("TreeMap: %v\n", treeMap)
	fmt.Printf("Size: %d\n", treeMap.Size())

	// Min and Max
	if minKey, minVal, exists := treeMap.Min(); exists {
		fmt.Printf("Minimum: %s -> %d\n", minKey, minVal)
	}
	if maxKey, maxVal, exists := treeMap.Max(); exists {
		fmt.Printf("Maximum: %s -> %d\n", maxKey, maxVal)
	}

	// Floor and Ceiling
	if floorKey, floorVal, exists := treeMap.Floor("coconut"); exists {
		fmt.Printf("Floor of 'coconut': %s -> %d\n", floorKey, floorVal)
	}
	if ceilingKey, ceilingVal, exists := treeMap.Ceiling("coconut"); exists {
		fmt.Printf("Ceiling of 'coconut': %s -> %d\n", ceilingKey, ceilingVal)
	}

	// Lower and Higher
	if lowerKey, lowerVal, exists := treeMap.Lower("cherry"); exists {
		fmt.Printf("Lower than 'cherry': %s -> %d\n", lowerKey, lowerVal)
	}
	if higherKey, higherVal, exists := treeMap.Higher("cherry"); exists {
		fmt.Printf("Higher than 'cherry': %s -> %d\n", higherKey, higherVal)
	}

	// Rank and Select
	rank := treeMap.Rank("cherry")
	fmt.Printf("Rank of 'cherry': %d\n", rank)

	if selectKey, selectVal, exists := treeMap.Select(2); exists {
		fmt.Printf("Element at rank 2: %s -> %d\n", selectKey, selectVal)
	}

	// Range queries
	rangeResult := treeMap.Range("banana", "date")
	fmt.Printf("Range [banana, date]: %v\n", rangeResult)

	// Get operations
	if val, exists := treeMap.Get("banana"); exists {
		fmt.Printf("Value for 'banana': %d\n", val)
	}

	// Remove operation
	treeMap.Remove("cherry")
	fmt.Printf("After removing 'cherry': %v\n", treeMap)

	// Keys and Values
	fmt.Printf("Keys: %v\n", treeMap.Keys())
	fmt.Printf("Values: %v\n", treeMap.Values())

	// Functional operations
	longKeys := treeMap.Filter(func(key string, value int) bool { return len(key) > 5 })
	fmt.Printf("Entries with key length > 5: %v\n", longKeys)

	fmt.Println()
}

func exampleGraph() {
	fmt.Println("11. Graph Operations:")
	fmt.Println("-------------------")

	// Create a new directed graph
	graph := stl.NewGraph[int](true)

	// Add edges
	edges := [][2]int{
		{1, 2}, {1, 3}, {2, 4}, {3, 4},
		{4, 5}, {3, 5}, {5, 6}, {6, 7},
		{5, 8}, {8, 9}, {7, 9},
	}

	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	fmt.Printf("Graph: %v\n", graph)

	// Get neighbors
	fmt.Printf("Neighbors of 5: %v\n", graph.GetNeighbors(5))

	// Degree of a node
	fmt.Printf("Degree of 5: %d\n", graph.Degree(5))

	// BFS and DFS
	fmt.Printf("BFS starting from 1: %v\n", graph.BFS(1))
	fmt.Printf("DFS starting from 1: %v\n", graph.DFS(1))

	// Shortest path
	if path, exists := graph.ShortestPath(1, 9); exists {
		fmt.Printf("Shortest path from 1 to 9: %v\n", path)
	}

	// Minimum spanning tree (MST) - Prim's algorithm
	fmt.Printf("Minimum Spanning Tree (Prim's) starting from 1: %v\n", graph.PrimMST(1))

	// Functional operations
	largeDegreeNodes := graph.Filter(func(node int, degree int) bool { return degree > 2 })
	fmt.Printf("Nodes with degree > 2: %v\n", largeDegreeNodes)

	fmt.Println()
}

func exampleAdvancedFeatures() {
	fmt.Println("12. Advanced Features:")
	fmt.Println("-------------------")

	// Example: Word frequency analysis using MultiSet
	fmt.Println("Word Frequency Analysis:")
	text := "the quick brown fox jumps over the lazy dog the fox is quick"
	words := strings.Fields(text)

	wordFreq := stl.NewMultiSet[string]()
	for _, word := range words {
		wordFreq.Add(word)
	}

	fmt.Printf("Word frequencies: %v\n", wordFreq.ToCountMap())
	fmt.Printf("Most common words: %v\n", wordFreq.MostCommon(3))

	// Example: Autocomplete using Trie
	fmt.Println("\nAutocomplete System:")
	autocomplete := stl.NewTrie()
	dictionary := []string{"algorithm", "alphabet", "alpine", "altitude", "always", "amazing", "ambition"}
	for _, word := range dictionary {
		autocomplete.Insert(word)
	}

	suggestions := autocomplete.GetWordsWithPrefixLimit("al", 5)
	fmt.Printf("Suggestions for 'al': %v\n", suggestions)

	// Example: Social network using Graph
	fmt.Println("\nSocial Network Graph:")
	socialGraph := stl.NewGraph[string](false)
	friendships := [][2]string{
		{"Alice", "Bob"}, {"Alice", "Charlie"}, {"Bob", "David"},
		{"Charlie", "Eve"}, {"David", "Eve"}, {"Eve", "Frank"},
	}

	for _, friendship := range friendships {
		socialGraph.AddEdge(friendship[0], friendship[1])
	}

	// Find mutual friends
	aliceFriends := stl.NewSetFromSlice(socialGraph.GetNeighbors("Alice"))
	bobFriends := stl.NewSetFromSlice(socialGraph.GetNeighbors("Bob"))
	mutualFriends := aliceFriends.Intersection(bobFriends)
	fmt.Printf("Mutual friends of Alice and Bob: %v\n", mutualFriends)

	// Example: Priority queue simulation using TreeMap
	fmt.Println("\nPriority Queue Simulation:")
	priorityQueue := stl.NewTreeMap[int, string](func(a, b int) bool { return a < b })
	priorityQueue.Put(3, "Low priority task")
	priorityQueue.Put(1, "High priority task")
	priorityQueue.Put(2, "Medium priority task")

	fmt.Println("Processing tasks in priority order:")
	for !priorityQueue.IsEmpty() {
		if key, value, exists := priorityQueue.Min(); exists {
			fmt.Printf("Processing priority %d: %s\n", key, value)
			priorityQueue.Remove(key)
		}
	}

	// Example: Sliding window using Deque
	fmt.Println("\nSliding Window Maximum:")
	numbers := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3 // window size

	deque := stl.NewDeque[int](len(numbers))
	result := make([]int, 0, len(numbers)-k+1)

	for i, num := range numbers {
		// Remove elements outside the window
		if !deque.IsEmpty() {
			if front, _ := deque.Front(); front <= i-k {
				deque.PopFront()
			}
		}

		// Remove smaller elements from the back
		for !deque.IsEmpty() {
			if back, _ := deque.Back(); numbers[back] < num {
				deque.PopBack()
			} else {
				break
			}
		}

		deque.PushBack(i)

		// Add maximum to result
		if i >= k-1 {
			if front, _ := deque.Front(); front != -1 {
				result = append(result, numbers[front])
			}
		}
	}

	fmt.Printf("Sliding window maximum for window size %d: %v\n", k, result)
}
