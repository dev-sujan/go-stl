package stl

import (
	"testing"
)

func TestGraphBasicOperations(t *testing.T) {
	// Test undirected graph
	graph := NewGraph[string](false)

	// Test AddNode
	nodes := []string{"A", "B", "C", "D"}
	for _, node := range nodes {
		graph.AddNode(node)
	}

	if graph.NodeCount() != len(nodes) {
		t.Errorf("Expected node count %d, got %d", len(nodes), graph.NodeCount())
	}

	// Test AddEdge
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	if graph.EdgeCount() != 4 {
		t.Errorf("Expected edge count 4, got %d", graph.EdgeCount())
	}

	// Test HasEdge
	if !graph.HasEdge("A", "B") {
		t.Error("Graph should have edge from A to B")
	}
	if !graph.HasEdge("B", "A") {
		t.Error("Undirected graph should have edge from B to A")
	}
	if graph.HasEdge("A", "C") {
		t.Error("Graph should not have edge from A to C")
	}

	// Test GetNeighbors
	aNeighbors := graph.GetNeighbors("A")
	if len(aNeighbors) != 2 || !containsNode(aNeighbors, "B") || !containsNode(aNeighbors, "D") {
		t.Errorf("Node A should have neighbors B and D, got %v", aNeighbors)
	}
}

func TestDirectedGraph(t *testing.T) {
	graph := NewGraph[string](true)

	nodes := []string{"A", "B", "C", "D"}
	for _, node := range nodes {
		graph.AddNode(node)
	}

	// Add edges
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	// Test directed edges
	if !graph.HasEdge("A", "B") {
		t.Error("Graph should have edge from A to B")
	}
	if graph.HasEdge("B", "A") {
		t.Error("Directed graph should not have edge from B to A")
	}

	// Test GetNeighbors
	aNeighbors := graph.GetNeighbors("A")
	if len(aNeighbors) != 1 || !containsNode(aNeighbors, "B") {
		t.Errorf("Node A should have neighbor B only, got %v", aNeighbors)
	}
}

func TestGraphFromEdges(t *testing.T) {
	edges := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"D", "A"},
	}

	// Undirected graph
	graph := NewGraphFromEdges(edges, false)

	if graph.NodeCount() != 4 {
		t.Errorf("Expected node count 4, got %d", graph.NodeCount())
	}

	if graph.EdgeCount() != 4 {
		t.Errorf("Expected edge count 4, got %d", graph.EdgeCount())
	}

	if !graph.HasEdge("A", "B") || !graph.HasEdge("B", "A") {
		t.Error("Undirected graph should have edges in both directions")
	}

	// Directed graph
	directedGraph := NewGraphFromEdges(edges, true)
	if directedGraph.HasEdge("B", "A") {
		t.Error("Directed graph should not have edge from B to A")
	}
}

func TestGraphRemoveNode(t *testing.T) {
	graph := NewGraph[string](false)

	// Add nodes and edges
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	// Remove node B
	graph.RemoveNode("B")

	if graph.NodeCount() != 3 {
		t.Errorf("Expected node count 3 after removal, got %d", graph.NodeCount())
	}

	if graph.HasNode("B") {
		t.Error("Graph should not have node B after removal")
	}

	if graph.HasEdge("A", "B") || graph.HasEdge("B", "C") {
		t.Error("Graph should not have edges to/from node B after removal")
	}
}

func TestGraphRemoveEdge(t *testing.T) {
	graph := NewGraph[string](false)

	// Add nodes and edges
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	// Remove edge
	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") || graph.HasEdge("B", "A") {
		t.Error("Graph should not have edge between A and B after removal")
	}

	if graph.EdgeCount() != 3 {
		t.Errorf("Expected edge count 3 after removal, got %d", graph.EdgeCount())
	}
}

// TestGraphBFS is skipped as the method is not implemented
func TestGraphBFS(t *testing.T) {
	t.Skip("BFS method not implemented yet")
}

// TestGraphDFS is skipped as the method is not implemented
func TestGraphDFS(t *testing.T) {
	t.Skip("DFS method not implemented yet")
}

func TestGraphClear(t *testing.T) {
	graph := NewGraph[string](false)

	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	graph.Clear()

	if graph.NodeCount() != 0 {
		t.Errorf("Graph should have 0 nodes after clear, got %d", graph.NodeCount())
	}

	if graph.EdgeCount() != 0 {
		t.Errorf("Graph should have 0 edges after clear, got %d", graph.EdgeCount())
	}
}

// Helper function for tests
func containsNode[T comparable](nodes []T, target T) bool {
	for _, node := range nodes {
		if node == target {
			return true
		}
	}
	return false
}
