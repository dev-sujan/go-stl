# Go STL (Standard Template Library)

<p align="center">
  <img src="https://raw.githubusercontent.com/dev-sujan/go-stl/main/docs/logo.png" alt="Go STL Logo" width="320"/>
</p>

[![Go Reference](https://pkg.go.dev/badge/github.com/dev-sujan/go-stl.svg)](https://pkg.go.dev/github.com/dev-sujan/go-stl)
[![Go Report Card](https://goreportcard.com/badge/github.com/dev-sujan/go-stl)](https://goreportcard.com/report/github.com/dev-sujan/go-stl)
[![Build Status](https://github.com/dev-sujan/go-stl/actions/workflows/ci.yml/badge.svg)](https://github.com/dev-sujan/go-stl/actions)
[![Latest Release](https://img.shields.io/github/v/release/dev-sujan/go-stl)](https://github.com/dev-sujan/go-stl/releases/latest)
[![codecov](https://codecov.io/gh/dev-sujan/go-stl/branch/main/graph/badge.svg)](https://codecov.io/gh/dev-sujan/go-stl)

A comprehensive, robust, and user-friendly implementation of essential data structures and algorithms for Go, inspired by Java Collections and C++ STL. Designed for productivity, clarity, and performance, with a focus on generics, a consistent API, and advanced features.

---

## 👤 Ownership & Maintenance

This project is owned and maintained by [Sujan Maji](https://github.com/dev-sujan).

## 💬 Community & Support

- [Open an issue](https://github.com/dev-sujan/go-stl/issues)
- [Discussions](https://github.com/dev-sujan/go-stl/discussions)

---

## 📦 Overview

Go STL provides a full suite of data structures missing from Go's standard library, with a clean, consistent, and beginner-friendly API. All structures are generic, efficient, and come with advanced and functional programming support.

**Included structures:**
- **Set** (Unordered & Ordered)
- **MultiSet** (Bag)
- **MultiMap**
- **Deque** (Double-Ended Queue)
- **Binary Search Tree (BST)**
- **Trie** (Prefix Tree)
- **Graph**
- **TreeMap** (Ordered/Sorted Map)
- **Stack** (LIFO)
- **Queue** (FIFO)
- **PriorityQueue** (Heap-based)

---

## ✨ Features

- **Generics-first:** Type-safe, flexible, and future-proof
- **Consistent API:** Learn once, use everywhere
- **Functional support:** Filter, Map, ForEach, Any, All, and more
- **Advanced operations:** Sorting, searching, set/graph/trie algorithms
- **Optimized:** Fast, memory-efficient implementations
- **Well-documented:** Clear, example-driven docs
- **Battle-tested:** Comprehensive unit and benchmark tests

---

## 🚧 Roadmap & Planned Features

- AVL Tree (self-balancing BST)
- Red-Black Tree
- B-Tree (disk-friendly, large datasets)
- Skip List
- Bloom Filter
- Disjoint Set (Union-Find)
- Segment Tree
- Fenwick Tree (Binary Indexed Tree)
- Memory pooling for performance
- SIMD/vectorized operations
- Lock-free/concurrent data structures
- Compression/space-efficient representations
- More graph algorithms (Dijkstra, A*, etc.)
- More trie/pattern matching features
- Community suggestions welcome! Open an issue or discussion to propose features.

---

## 🚀 Installation

```bash
go get github.com/dev-sujan/go-stl
```

---

## 🏁 Quick Start

```go
package main

import (
    "fmt"
    "github.com/dev-sujan/go-stl/stl"
)

func main() {
    // Set: Unique values
    set := stl.NewSet[int]()
    set.Add(1)
    set.Add(2)
    fmt.Println("Set:", set) // Set[1 2]

    // MultiSet: Count duplicates
    ms := stl.NewMultiSet[string]()
    ms.Add("apple")
    ms.Add("apple")
    ms.Add("banana")
    fmt.Println("MultiSet:", ms) // MultiSetmap[apple:2 banana:1]

    // Deque: Fast front/back
    deque := stl.NewDeque[int](10)
    deque.PushBack(1)
    deque.PushFront(0)
    deque.PushBack(2)
    fmt.Println("Deque:", deque) // Deque[0 1 2]

    // BST: Ordered data
    bst := stl.NewBST[int](func(a, b int) bool { return a < b })
    bst.Insert(5)
    bst.Insert(3)
    bst.Insert(7)
    fmt.Println("BST InOrder:", bst.InOrder()) // [3 5 7]

    // Create a Trie
    trie := stl.NewTrie()
    trie.Insert("hello")
    trie.Insert("world")
    trie.Insert("help")
    fmt.Println("Words with prefix 'he':", trie.GetWordsWithPrefix("he")) // [hello help]

    // Create a Graph
    graph := stl.NewGraph[int](false) // undirected
    graph.AddEdge(1, 2)
    graph.AddEdge(2, 3)
    graph.AddEdge(1, 3)
    fmt.Println("Graph BFS from 1:", graph.BFS(1)) // [1 2 3]

    // Create a TreeMap
    treeMap := stl.NewTreeMap[string, int](func(a, b string) bool { return a < b })
    treeMap.Put("apple", 1)
    treeMap.Put("banana", 2)
    treeMap.Put("cherry", 3)
    fmt.Println("TreeMap Keys:", treeMap.Keys()) // [apple banana cherry]
}
```

---

## ❓ Why Go STL?

Go STL fills the gap in Go's standard library by providing a modern, consistent, and feature-rich set of data structures and algorithms, inspired by the best of Java and C++ but designed for Go's simplicity and performance. 

- **No more reinventing the wheel:** Use robust, tested, and production-ready collections out of the box.
- **Consistent, intuitive API:** Learn once, use everywhere—no surprises.
- **Generics-first:** Type-safe and future-proof, leveraging Go's latest features.
- **Advanced features:** Functional programming, sorting, searching, and more.
- **Performance-focused:** Optimized for speed and memory efficiency.
- **Comprehensive docs & tests:** Easy onboarding and reliable maintenance.
- **Community-driven:** Open to suggestions, improvements, and contributions.

---

## 📖 API Reference

- [GoDoc / pkg.go.dev documentation](https://pkg.go.dev/github.com/dev-sujan/go-stl)

---

## 📚 Data Structures & API

### Set
Unordered collection of unique elements with a full suite of set operations and utilities.
```go
set := stl.NewSet[int]()
set.Add(1)
set.Add(2)
set.Remove(1)
set.Contains(1)
set.Size()
set.IsEmpty()
set.Clear()
set.Clone()
set.Equals(otherSet)
set.ToSlice()
set.Union(otherSet)
set.Intersection(otherSet)
set.Difference(otherSet)
set.SymmetricDifference(otherSet)
set.IsSubset(otherSet)
set.IsSuperset(otherSet)
set.IsDisjoint(otherSet)
set.ForEach(func(x int) { fmt.Println(x) })
set.Filter(func(x int) bool { return x > 0 })
set.Map(func(x int) int { return x * 2 })
set.Any(func(x int) bool { return x%2 == 0 })
set.All(func(x int) bool { return x > 0 })
```
- **Time Complexity:** Add/Remove/Contains: O(1) avg; Set ops: O(n + m)

### MultiSet
Collection with duplicate tracking and all major multiset operations.
```go
ms := stl.NewMultiSet[string]()
ms.Add("apple")
ms.AddCount("apple", 3)
ms.Remove("apple")
ms.RemoveAll("apple")
ms.Contains("apple")
ms.Count("apple")
ms.Size()
ms.IsEmpty()
ms.Clear()
ms.Clone()
ms.Equals(otherMS)
ms.ToSlice()
ms.MostCommon(3)
ms.LeastCommon(2)
ms.Union(otherMS)
ms.Intersection(otherMS)
ms.Difference(otherMS)
ms.ForEach(func(x string, count int) { fmt.Println(x, count) })
ms.Filter(func(x string, count int) bool { return count > 1 })
ms.Map(func(x string, count int) string { return strings.ToUpper(x) })
ms.Any(func(x string, count int) bool { return count > 2 })
ms.All(func(x string, count int) bool { return count > 0 })
```
- **Time Complexity:** Add/Remove/Count: O(1) avg; MostCommon: O(n log n)

### MultiMap
Map with multiple values per key and a complete API for manipulation and queries.
```go
mm := stl.NewMultiMap[string, int]()
mm.Put("fruit", 1)
mm.PutAll("fruit", []int{2, 3})
mm.Get("fruit")
mm.GetFirst("fruit")
mm.GetLast("fruit")
mm.Remove("fruit", 1)
mm.RemoveAll("fruit")
mm.ContainsKey("fruit")
mm.ContainsValue(2)
mm.ContainsEntry("fruit", 2)
mm.Size()
mm.IsEmpty()
mm.Clear()
mm.Clone()
mm.Equals(otherMM)
mm.ToSlice()
mm.Invert()
mm.GetSortedKeys(func(a, b string) bool { return a < b })
mm.GetSortedValues("fruit", func(a, b int) bool { return a < b })
mm.ForEach(func(k string, v int) { fmt.Println(k, v) })
mm.Filter(func(k string, v int) bool { return v > 1 })
mm.Map(func(k string, v int) int { return v * 2 })
mm.Any(func(k string, v int) bool { return v == 2 })
mm.All(func(k string, v int) bool { return v > 0 })
```
- **Time Complexity:** Put/Get: O(1) avg; Remove: O(n)

### Deque
Double-ended queue with all core, random access, capacity, equality, functional, and utility methods.
```go
deque := stl.NewDeque[int](16)
deque.PushFront(1)
deque.PushBack(2)
deque.PopFront()
deque.PopBack()
deque.Front()
deque.Back()
deque.At(0)
deque.Set(0, 10)
deque.Insert(1, 99)
deque.Remove(0)
deque.RotateLeft(1)
deque.RotateRight(1)
deque.Reverse()
deque.Capacity()
deque.Reserve(100)
deque.TrimToSize()
deque.Clone()
deque.Equals(otherDeque)
deque.Filter(func(x int) bool { return x > 0 })
deque.Map(func(x int) int { return x * x })
deque.ForEach(func(x int) { fmt.Println(x) })
deque.ForEachReversed(func(x int) { fmt.Println(x) })
deque.ToSlice()
deque.IsEmpty()
deque.Size()
```
- **Time Complexity:** Push/Pop: O(1) amortized; Random access: O(1); Insert/Remove: O(n)

### Binary Search Tree (BST)
Ordered tree structure with a full set of search, traversal, and range operations.
```go
bst := stl.NewBST[int](func(a, b int) bool { return a < b })
bst.Insert(5)
bst.Search(5)
bst.Delete(5)
bst.Min()
bst.Max()
bst.Floor(4)
bst.Ceiling(6)
bst.Successor(5)
bst.Predecessor(5)
bst.Rank(5)
bst.Select(1)
bst.Range(3, 7)
bst.InOrder()
bst.PreOrder()
bst.PostOrder()
bst.LevelOrder()
bst.Size()
bst.IsEmpty()
bst.Clear()
bst.Clone()
bst.Equals(otherBST)
bst.ForEach(func(x int) { fmt.Println(x) })
```
- **Time Complexity:** Insert/Search/Delete: O(log n) avg, O(n) worst

### Trie
Prefix tree for string operations, pattern matching, and advanced queries.
```go
trie := stl.NewTrie()
trie.Insert("hello")
trie.Search("hello")
trie.Delete("hello")
trie.StartsWith("he")
trie.GetWordsWithPrefix("he")
trie.GetWordsWithPattern("h?llo")
trie.GetWordsByLength(5)
trie.GetWordsWithinDistance("hello", 2)
trie.EditDistance("hello", "helo")
trie.LongestCommonPrefix()
trie.Size()
trie.IsEmpty()
trie.Clear()
trie.Clone()
trie.Equals(otherTrie)
trie.ForEach(func(word string) { fmt.Println(word) })
trie.Filter(func(word string) bool { return len(word) > 3 })
trie.Map(func(word string) string { return strings.ToUpper(word) })
trie.Any(func(word string) bool { return word == "hello" })
trie.All(func(word string) bool { return len(word) > 0 })
```
- **Time Complexity:** Insert/Search: O(m); Prefix search: O(m + k); Pattern search: O(m + k); EditDistance: O(m^2)

### Graph
Adjacency list with all major graph algorithms and utilities.
```go
graph := stl.NewGraph[int](false)
graph.AddEdge(1, 2)
graph.AddEdge(2, 3)
graph.RemoveEdge(1, 2)
graph.AddNode(4)
graph.RemoveNode(3)
graph.BFS(1)
graph.DFS(1)
graph.DFSIterative(1)
graph.ShortestPath(1, 3)
graph.AllPaths(1, 3)
graph.ConnectedComponents()
graph.HasCycle()
graph.TopologicalSort()
graph.IsBipartite()
graph.Degree(2)
graph.InDegree(2)
graph.OutDegree(2)
graph.PrimMST()
graph.Filter(func(node, degree int) bool { return degree > 2 })
graph.Size()
graph.IsEmpty()
graph.Clear()
graph.Clone()
graph.Equals(otherGraph)
graph.ForEach(func(node int) { fmt.Println(node) })
```
- **Time Complexity:** AddEdge/RemoveEdge: O(1); BFS/DFS: O(V+E); ShortestPath: O(V+E); TopologicalSort: O(V+E); MST: O(E log V)

### TreeMap
Ordered map using BST with a complete set of map and range operations.
```go
treeMap := stl.NewTreeMap[string, int](func(a, b string) bool { return a < b })
treeMap.Put("apple", 1)
treeMap.Get("apple")
treeMap.Remove("apple")
treeMap.Min()
treeMap.Max()
treeMap.Floor("banana")
treeMap.Ceiling("banana")
treeMap.Lower("banana")
treeMap.Higher("banana")
treeMap.Rank("banana")
treeMap.Select(1)
treeMap.Range("apple", "cherry")
treeMap.Keys()
treeMap.Values()
treeMap.Entries()
treeMap.Size()
treeMap.IsEmpty()
treeMap.Clear()
treeMap.Clone()
treeMap.Equals(otherTreeMap)
treeMap.ForEach(func(k string, v int) { fmt.Println(k, v) })
```
- **Time Complexity:** Put/Get/Remove: O(log n) avg

### Stack
LIFO structure with a rich API, random access, capacity, and functional support.
```go
stack := stl.NewStack[int]()
stack.Push(1)
stack.PushAll([]int{2, 3})
stack.Pop() // 3
stack.Peek() // 2
stack.Size() // 2
stack.IsEmpty() // false
stack.Clear()
stack.GetAt(0)
stack.SetAt(0, 10)
stack.RemoveAt(0)
stack.InsertAt(0, 99)
stack.Contains(2)
stack.IndexOf(2)
stack.LastIndexOf(2)
stack.Remove(2)
stack.RemoveAll(2)
stack.Sort(func(a, b int) bool { return a < b })
stack.SortStable(func(a, b int) bool { return a < b })
stack.Shuffle()
stack.Reverse()
stack.Take(2)
stack.Drop(1)
stack.Capacity()
stack.Reserve(100)
stack.TrimToSize()
stack.Clone()
stack.Equals(otherStack)
stack.Filter(func(x int) bool { return x > 0 })
stack.Map(func(x int) int { return x * x })
stack.ForEach(func(x int) { fmt.Println(x) })
stack.ForEachReversed(func(x int) { fmt.Println(x) })
```
- **Time Complexity:** Push/Pop/Peek: O(1); Random access: O(1); Search: O(n); Sort: O(n log n)

### Queue
FIFO structure with a rich API, random access, capacity, and functional support.
```go
queue := stl.NewQueue[string]()
queue.Enqueue("first")
queue.EnqueueAll([]string{"second", "third"})
queue.Dequeue()
queue.Peek()
queue.PeekBack()
queue.Size()
queue.IsEmpty()
queue.Clear()
queue.GetAt(0)
queue.SetAt(0, "foo")
queue.RemoveAt(0)
queue.InsertAt(0, "bar")
queue.Contains("second")
queue.IndexOf("second")
queue.LastIndexOf("second")
queue.Remove("second")
queue.RemoveAll("second")
queue.Sort(func(a, b string) bool { return a < b })
queue.SortStable(func(a, b string) bool { return a < b })
queue.Shuffle()
queue.Reverse()
queue.Take(2)
queue.Drop(1)
queue.Capacity()
queue.Reserve(100)
queue.TrimToSize()
queue.Clone()
queue.Equals(otherQueue)
queue.Filter(func(s string) bool { return len(s) > 3 })
queue.Map(func(s string) string { return strings.ToUpper(s) })
queue.ForEach(func(s string) { fmt.Println(s) })
queue.ForEachReversed(func(s string) { fmt.Println(s) })
```
- **Time Complexity:** Enqueue/Dequeue/Peek: O(1); Random access: O(1); Search: O(n); Sort: O(n log n)

### PriorityQueue
Heap-based queue with custom ordering and all major queue operations.
```go
pq := stl.NewPriorityQueue[int](func(a, b int) bool { return a < b })
pq.Enqueue(5)
pq.Dequeue()
pq.Peek()
pq.Size()
pq.IsEmpty()
pq.Clear()
pq.Clone()
pq.Equals(otherPQ)
```
- **Time Complexity:** Enqueue/Dequeue/Peek: O(log n); Size/IsEmpty: O(1)

---

## ⚡ Performance & Complexity

Below is a comprehensive table of time complexities for all major operations, including advanced and special operations, for each data structure. Let n = number of elements, m = word length (Trie), V = vertices, E = edges.

| Operation                | Set         | MultiSet    | MultiMap    | Deque         | BST / TreeMap | Trie         | Graph         | Stack / Queue | PriorityQueue |
|--------------------------|-------------|-------------|-------------|---------------|---------------|--------------|---------------|---------------|---------------|
| Add/Insert/Put           | O(1) avg    | O(1) avg    | O(1) avg    | O(1) amortized| O(log n) avg  | O(m)         | O(1)          | O(1)          | O(log n)      |
| Remove/Delete            | O(1) avg    | O(1) avg    | O(n)        | O(1) amortized| O(log n) avg  | O(m)         | O(1)          | O(1)          | O(log n)      |
| Contains/Search/Get      | O(1) avg    | O(1) avg    | O(1) avg    | O(1)          | O(log n) avg  | O(m)         | O(1)          | O(n)          | O(1)          |
| Min/Max                  | O(n)        | O(n log n)  | O(n)        | O(1)          | O(log n)      | O(n)         | O(n)          | O(n)          | O(1)          |
| Random Access (At/GetAt) | -           | -           | -           | O(1)          | -             | -            | -             | O(1)          | -             |
| Set Ops (Union, etc.)    | O(n + m)    | O(n + m)    | -           | -             | -             | -            | -             | -             | -             |
| Range Query              | -           | -           | -           | -             | O(log n + k)  | O(m + k)     | -             | -             | -             |
| Traversal (ForEach, etc.)| O(n)        | O(n)        | O(n)        | O(n)          | O(n)          | O(n)         | O(V+E)        | O(n)          | O(n)          |
| Sort/StableSort          | -           | -           | -           | O(n log n)    | O(n log n)*   | -            | -             | O(n log n)    | -             |
| Reverse                  | O(n)        | O(n)        | O(n)        | O(n)          | O(n)          | O(n)         | O(V+E)        | O(n)          | -             |
| Clone/Copy               | O(n)        | O(n)        | O(n)        | O(n)          | O(n)          | O(n)         | O(V+E)        | O(n)          | O(n)          |
| Size/IsEmpty             | O(1)        | O(1)        | O(1)        | O(1)          | O(1)          | O(1)         | O(1)          | O(1)          | O(1)          |
| Clear                    | O(1)        | O(1)        | O(1)        | O(1)          | O(1)          | O(1)         | O(1)          | O(1)          | O(1)          |
| Advanced (see below)     | See notes   | See notes   | See notes   | See notes     | See notes     | See notes    | See notes     | See notes     | See notes     |

**Advanced/Special Operations Notes:**
- **Set/MultiSet:** MostCommon/LeastCommon: O(n log n). Equality: O(n). Filter/Map/Any/All: O(n).
- **MultiMap:** Remove(key, value): O(n) for values per key. Invert: O(n). SortedKeys/Values: O(n log n).
- **Deque:** Insert/Remove at index: O(n). Rotate: O(k). Reserve/Trim: O(n). ForEachReversed: O(n).
- **BST/TreeMap:** Successor/Predecessor/Rank/Select: O(log n) avg. Range: O(log n + k). Traversals: O(n).
- **Trie:** GetWordsWithPrefix: O(m + k). Pattern/Distance queries: O(n) worst. LongestCommonPrefix: O(m).
- **Graph:** BFS/DFS: O(V+E). ShortestPath: O(V+E) (unweighted). TopologicalSort: O(V+E). MST: O(E log V).
- **Stack/Queue:** Search/IndexOf: O(n). Sort/Shuffle/Reverse: O(n log n)/O(n)/O(n).
- **PriorityQueue:** Heapify: O(n). Update: O(log n) if supported.

*BST/TreeMap are always sorted; explicit sort is not needed. For unsorted structures, sort is O(n log n).

---

## 🧠 Best Practices

- **Choose the right structure**: Use Set for uniqueness, MultiSet for counting, MultiMap for one-to-many, Deque for double-ended ops, TreeMap for ordered keys, Trie for prefix/string ops, Graph for relationships, Stack/Queue for LIFO/FIFO, PriorityQueue for prioritized processing.
- **Use comparators**: For custom types, provide a comparator function.
- **Handle edge cases**: Check for empty collections, missing keys, cycles, etc.

---

## 🧪 Development and Testing

### Running Tests

To run all tests:

```bash
go test -v ./stl
```

To run tests with race detection and generate coverage:

```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./stl
```

### Running Examples

The examples directory contains sample usage of all data structures:

```bash
go run examples/main.go
```

### Continuous Integration

This project uses GitHub Actions for CI. The workflow includes:
- Testing on multiple platforms (Ubuntu, Windows, macOS) and Go versions
- Code linting with golangci-lint
- Security scanning with gosec
- Dependency scanning
- Code coverage reporting with Codecov

---

## 📪 Contact & Support

For questions, suggestions, or support, please use GitHub Issues or Discussions. We welcome feedback and contributions!

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new features
4. Ensure all tests pass
5. Submit a pull request

### Commit Message Guidelines

When contributing to this project, please follow these commit message conventions:

- **feat:** A new feature or enhancement (e.g., `feat: add Red-Black Tree implementation`)
- **fix:** A bug fix (e.g., `fix: resolve concurrency issue in Queue`)
- **docs:** Documentation changes (e.g., `docs: improve MultiSet usage examples`)
- **style:** Code style changes (formatting, missing semi-colons, etc.)
- **refactor:** Code refactoring without changing functionality
- **perf:** Performance improvements (e.g., `perf: optimize Set union operation`)
- **test:** Adding or updating tests
- **chore:** Changes to build process or auxiliary tools

Each commit should have a subject line (50 chars max) and optionally a body with more detailed explanation.

---

## 📄 License

MIT License. See LICENSE file for details.

---

## 🙏 Acknowledgments

- Inspired by Java Collections Framework
- Influenced by C++ Standard Template Library
- Built with Go's generics system

---

## ❓ FAQ / Troubleshooting

**Q: Is Go STL production-ready?**  
A: Yes! It is tested, benchmarked, and used in real projects. Please report any issues you find.

**Q: How do I use Go STL with my own types?**  
A: Just provide a comparator function for ordered structures (e.g., `func(a, b MyType) bool { ... }`).

**Q: Why not use Go's built-in maps/slices?**  
A: Go STL provides advanced features, consistent APIs, and data structures not available in the standard library (e.g., Set, MultiSet, TreeMap, Trie, Deque, etc.).

**Q: Is it fast?**  
A: Yes! All structures are optimized for performance and memory usage. See benchmarks in the repo.

**Q: How do I contribute?**  
A: See the Contributing section below. PRs, issues, and suggestions are welcome!

**Q: Where can I ask questions?**  
A: Use GitHub Discussions or open an issue.