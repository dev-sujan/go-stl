package stl

import (
	"fmt"
)

// Graph represents a graph using adjacency list representation
type Graph[T comparable] struct {
	adjacency map[T][]T
	directed  bool
}

// NewGraph creates a new empty graph
func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		adjacency: make(map[T][]T),
		directed:  directed,
	}
}

// NewGraphFromEdges creates a graph from a slice of edges
func NewGraphFromEdges[T comparable](edges [][2]T, directed bool) *Graph[T] {
	graph := NewGraph[T](directed)
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	return graph
}

// AddNode adds a node to the graph
func (g *Graph[T]) AddNode(node T) {
	if _, exists := g.adjacency[node]; !exists {
		g.adjacency[node] = []T{}
	}
}

// AddEdge adds an edge between two nodes
func (g *Graph[T]) AddEdge(from, to T) {
	g.AddNode(from)
	g.AddNode(to)

	g.adjacency[from] = append(g.adjacency[from], to)

	if !g.directed {
		g.adjacency[to] = append(g.adjacency[to], from)
	}
}

// RemoveNode removes a node and all its edges from the graph
func (g *Graph[T]) RemoveNode(node T) {
	// Remove all edges to this node
	for from := range g.adjacency {
		g.RemoveEdge(from, node)
	}

	// Remove the node itself
	delete(g.adjacency, node)
}

// RemoveEdge removes an edge between two nodes
func (g *Graph[T]) RemoveEdge(from, to T) {
	if neighbors, exists := g.adjacency[from]; exists {
		for i, neighbor := range neighbors {
			if neighbor == to {
				g.adjacency[from] = append(neighbors[:i], neighbors[i+1:]...)
				break
			}
		}
	}

	if !g.directed {
		if neighbors, exists := g.adjacency[to]; exists {
			for i, neighbor := range neighbors {
				if neighbor == from {
					g.adjacency[to] = append(neighbors[:i], neighbors[i+1:]...)
					break
				}
			}
		}
	}
}

// HasNode checks if a node exists in the graph
func (g *Graph[T]) HasNode(node T) bool {
	_, exists := g.adjacency[node]
	return exists
}

// HasEdge checks if an edge exists between two nodes
func (g *Graph[T]) HasEdge(from, to T) bool {
	if neighbors, exists := g.adjacency[from]; exists {
		for _, neighbor := range neighbors {
			if neighbor == to {
				return true
			}
		}
	}
	return false
}

// GetNeighbors returns all neighbors of a node
func (g *Graph[T]) GetNeighbors(node T) []T {
	if neighbors, exists := g.adjacency[node]; exists {
		result := make([]T, len(neighbors))
		copy(result, neighbors)
		return result
	}
	return []T{}
}

// GetNodes returns all nodes in the graph
func (g *Graph[T]) GetNodes() []T {
	nodes := make([]T, 0, len(g.adjacency))
	for node := range g.adjacency {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetEdges returns all edges in the graph
func (g *Graph[T]) GetEdges() [][2]T {
	var edges [][2]T
	visited := make(map[string]bool)

	for from, neighbors := range g.adjacency {
		for _, to := range neighbors {
			edgeKey := fmt.Sprintf("%v->%v", from, to)
			reverseKey := fmt.Sprintf("%v->%v", to, from)

			if !g.directed {
				if !visited[edgeKey] && !visited[reverseKey] {
					edges = append(edges, [2]T{from, to})
					visited[edgeKey] = true
					visited[reverseKey] = true
				}
			} else {
				if !visited[edgeKey] {
					edges = append(edges, [2]T{from, to})
					visited[edgeKey] = true
				}
			}
		}
	}

	return edges
}

// NodeCount returns the number of nodes in the graph
func (g *Graph[T]) NodeCount() int {
	return len(g.adjacency)
}

// EdgeCount returns the number of edges in the graph
func (g *Graph[T]) EdgeCount() int {
	total := 0
	for _, neighbors := range g.adjacency {
		total += len(neighbors)
	}

	if !g.directed {
		total /= 2
	}

	return total
}

// IsEmpty checks if the graph is empty
func (g *Graph[T]) IsEmpty() bool {
	return len(g.adjacency) == 0
}

// Clear removes all nodes and edges from the graph
func (g *Graph[T]) Clear() {
	g.adjacency = make(map[T][]T)
}

// IsDirected checks if the graph is directed
func (g *Graph[T]) IsDirected() bool {
	return g.directed
}

// Degree returns the degree of a node (number of edges)
func (g *Graph[T]) Degree(node T) int {
	if neighbors, exists := g.adjacency[node]; exists {
		return len(neighbors)
	}
	return 0
}

// InDegree returns the in-degree of a node (for directed graphs)
func (g *Graph[T]) InDegree(node T) int {
	if !g.directed {
		return g.Degree(node)
	}

	count := 0
	for _, neighbors := range g.adjacency {
		for _, neighbor := range neighbors {
			if neighbor == node {
				count++
			}
		}
	}
	return count
}

// OutDegree returns the out-degree of a node (for directed graphs)
func (g *Graph[T]) OutDegree(node T) int {
	return g.Degree(node)
}

// BFS performs breadth-first search starting from the given node
func (g *Graph[T]) BFS(start T) []T {
	var result []T
	visited := make(map[T]bool)
	queue := []T{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// DFS performs depth-first search starting from the given node
func (g *Graph[T]) DFS(start T) []T {
	var result []T
	visited := make(map[T]bool)
	g.dfsRecursive(start, visited, &result)
	return result
}

// dfsRecursive is the recursive helper for DFS
func (g *Graph[T]) dfsRecursive(node T, visited map[T]bool, result *[]T) {
	visited[node] = true
	*result = append(*result, node)

	for _, neighbor := range g.GetNeighbors(node) {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited, result)
		}
	}
}

// DFSIterative performs iterative depth-first search
func (g *Graph[T]) DFSIterative(start T) []T {
	var result []T
	visited := make(map[T]bool)
	stack := []T{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			result = append(result, node)

			// Add neighbors in reverse order to maintain DFS order
			neighbors := g.GetNeighbors(node)
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}

	return result
}

// ConnectedComponents returns all connected components in the graph
func (g *Graph[T]) ConnectedComponents() [][]T {
	var components [][]T
	visited := make(map[T]bool)

	for node := range g.adjacency {
		if !visited[node] {
			var component []T
			g.dfsRecursive(node, visited, &component)
			components = append(components, component)
		}
	}

	return components
}

// IsConnected checks if the graph is connected
func (g *Graph[T]) IsConnected() bool {
	if g.IsEmpty() {
		return true
	}

	components := g.ConnectedComponents()
	return len(components) == 1
}

// ShortestPath finds the shortest path between two nodes using BFS
func (g *Graph[T]) ShortestPath(start, end T) ([]T, bool) {
	if !g.HasNode(start) || !g.HasNode(end) {
		return nil, false
	}

	visited := make(map[T]bool)
	parent := make(map[T]T)
	queue := []T{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			// Reconstruct path
			var path []T
			current := end
			for current != start {
				path = append([]T{current}, path...)
				current = parent[current]
			}
			path = append([]T{start}, path...)
			return path, true
		}

		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, false
}

// AllPaths finds all paths between two nodes
func (g *Graph[T]) AllPaths(start, end T) [][]T {
	var paths [][]T
	visited := make(map[T]bool)
	g.findAllPaths(start, end, visited, []T{start}, &paths)
	return paths
}

// findAllPaths is the recursive helper for AllPaths
func (g *Graph[T]) findAllPaths(current, end T, visited map[T]bool, path []T, paths *[][]T) {
	if current == end {
		pathCopy := make([]T, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
		return
	}

	visited[current] = true

	for _, neighbor := range g.GetNeighbors(current) {
		if !visited[neighbor] {
			g.findAllPaths(neighbor, end, visited, append(path, neighbor), paths)
		}
	}

	visited[current] = false
}

// HasCycle checks if the graph has a cycle
func (g *Graph[T]) HasCycle() bool {
	visited := make(map[T]bool)
	recStack := make(map[T]bool)

	for node := range g.adjacency {
		if !visited[node] {
			if g.hasCycleDFS(node, visited, recStack) {
				return true
			}
		}
	}

	return false
}

// hasCycleDFS is the recursive helper for HasCycle
func (g *Graph[T]) hasCycleDFS(node T, visited, recStack map[T]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range g.GetNeighbors(node) {
		if !visited[neighbor] {
			if g.hasCycleDFS(neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[node] = false
	return false
}

// TopologicalSort performs topological sorting (for DAGs)
func (g *Graph[T]) TopologicalSort() ([]T, bool) {
	if !g.directed {
		return nil, false
	}

	if g.HasCycle() {
		return nil, false
	}

	var result []T
	visited := make(map[T]bool)

	for node := range g.adjacency {
		if !visited[node] {
			g.topologicalSortDFS(node, visited, &result)
		}
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, true
}

// topologicalSortDFS is the recursive helper for TopologicalSort
func (g *Graph[T]) topologicalSortDFS(node T, visited map[T]bool, result *[]T) {
	visited[node] = true

	for _, neighbor := range g.GetNeighbors(node) {
		if !visited[neighbor] {
			g.topologicalSortDFS(neighbor, visited, result)
		}
	}

	*result = append(*result, node)
}

// IsBipartite checks if the graph is bipartite
func (g *Graph[T]) IsBipartite() bool {
	if g.IsEmpty() {
		return true
	}

	color := make(map[T]int)

	for node := range g.adjacency {
		if color[node] == 0 {
			if !g.isBipartiteBFS(node, color) {
				return false
			}
		}
	}

	return true
}

// isBipartiteBFS is the BFS helper for IsBipartite
func (g *Graph[T]) isBipartiteBFS(start T, color map[T]int) bool {
	queue := []T{start}
	color[start] = 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.GetNeighbors(node) {
			if color[neighbor] == 0 {
				color[neighbor] = -color[node]
				queue = append(queue, neighbor)
			} else if color[neighbor] == color[node] {
				return false
			}
		}
	}

	return true
}

// Clone creates a deep copy of the graph
func (g *Graph[T]) Clone() *Graph[T] {
	result := NewGraph[T](g.directed)

	for node, neighbors := range g.adjacency {
		result.adjacency[node] = make([]T, len(neighbors))
		copy(result.adjacency[node], neighbors)
	}

	return result
}

// Equals checks if two graphs are equal
func (g *Graph[T]) Equals(other *Graph[T]) bool {
	if g.directed != other.directed {
		return false
	}

	if len(g.adjacency) != len(other.adjacency) {
		return false
	}

	for node, neighbors := range g.adjacency {
		otherNeighbors, exists := other.adjacency[node]
		if !exists || len(neighbors) != len(otherNeighbors) {
			return false
		}

		// Create sets for comparison
		set1 := NewSetFromSlice(neighbors)
		set2 := NewSetFromSlice(otherNeighbors)

		if !set1.Equals(set2) {
			return false
		}
	}

	return true
}

// String returns a string representation of the graph
func (g *Graph[T]) String() string {
	return fmt.Sprintf("Graph{Directed: %v, Nodes: %d, Edges: %d}", g.directed, g.NodeCount(), g.EdgeCount())
}

// ForEachNode applies a function to each node in the graph
func (g *Graph[T]) ForEachNode(fn func(T)) {
	for node := range g.adjacency {
		fn(node)
	}
}

// ForEachEdge applies a function to each edge in the graph
func (g *Graph[T]) ForEachEdge(fn func(T, T)) {
	visited := make(map[string]bool)

	for from, neighbors := range g.adjacency {
		for _, to := range neighbors {
			edgeKey := fmt.Sprintf("%v->%v", from, to)
			reverseKey := fmt.Sprintf("%v->%v", to, from)

			if !g.directed {
				if !visited[edgeKey] && !visited[reverseKey] {
					fn(from, to)
					visited[edgeKey] = true
					visited[reverseKey] = true
				}
			} else {
				if !visited[edgeKey] {
					fn(from, to)
					visited[edgeKey] = true
				}
			}
		}
	}
}

// FilterNodes returns a new graph containing only nodes that satisfy the predicate
func (g *Graph[T]) FilterNodes(predicate func(T) bool) *Graph[T] {
	result := NewGraph[T](g.directed)

	for node := range g.adjacency {
		if predicate(node) {
			result.AddNode(node)
		}
	}

	for from, neighbors := range g.adjacency {
		if result.HasNode(from) {
			for _, to := range neighbors {
				if result.HasNode(to) {
					result.AddEdge(from, to)
				}
			}
		}
	}

	return result
}

// Subgraph returns a subgraph containing only the specified nodes
func (g *Graph[T]) Subgraph(nodes []T) *Graph[T] {
	nodeSet := NewSetFromSlice(nodes)
	return g.FilterNodes(func(node T) bool {
		return nodeSet.Contains(node)
	})
}

// Complement returns the complement of the graph
func (g *Graph[T]) Complement() *Graph[T] {
	result := NewGraph[T](g.directed)

	// Add all nodes
	for node := range g.adjacency {
		result.AddNode(node)
	}

	// Add all missing edges
	for from := range g.adjacency {
		for to := range g.adjacency {
			if from != to && !g.HasEdge(from, to) {
				result.AddEdge(from, to)
			}
		}
	}

	return result
}

// Union returns the union of two graphs
func (g *Graph[T]) Union(other *Graph[T]) *Graph[T] {
	if g.directed != other.directed {
		return nil
	}

	result := g.Clone()

	for from, neighbors := range other.adjacency {
		result.AddNode(from)
		for _, to := range neighbors {
			result.AddNode(to)
			result.AddEdge(from, to)
		}
	}

	return result
}

// Intersection returns the intersection of two graphs
func (g *Graph[T]) Intersection(other *Graph[T]) *Graph[T] {
	if g.directed != other.directed {
		return nil
	}

	result := NewGraph[T](g.directed)

	for from, neighbors := range g.adjacency {
		if other.HasNode(from) {
			result.AddNode(from)
			for _, to := range neighbors {
				if other.HasEdge(from, to) {
					result.AddNode(to)
					result.AddEdge(from, to)
				}
			}
		}
	}

	return result
}

// PrimMST returns the edges of the Minimum Spanning Tree using Prim's algorithm (for weighted graphs)
// For unweighted graphs, it returns a spanning tree (not minimum)
func (g *Graph[T]) PrimMST(start T) [][2]T {
	visited := make(map[T]bool)
	var mst [][2]T
	queue := []T{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.GetNeighbors(node) {
			if !visited[neighbor] {
				mst = append(mst, [2]T{node, neighbor})
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return mst
}

// Filter returns a slice of nodes that satisfy the predicate, along with their degree
func (g *Graph[T]) Filter(predicate func(node T, degree int) bool) []T {
	var result []T
	for node := range g.adjacency {
		if predicate(node, g.Degree(node)) {
			result = append(result, node)
		}
	}
	return result
}
