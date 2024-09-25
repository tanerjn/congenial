package main

import (
	"fmt"
	"math"
	"os"
)

// Graph structure representing the flow network
type Graph struct {
	capacity map[int]map[int]int // edge capacities
	flow     map[int]map[int]int // current flow
	source   int                 // source node
	sink     int                 // sink node
}

// NewGraph initializes a new flow network
func NewGraph(source, sink int) *Graph {
	return &Graph{
		capacity: make(map[int]map[int]int),
		flow:     make(map[int]map[int]int),
		source:   source,
		sink:     sink,
	}
}

// AddEdge adds a directed edge with the given capacity
func (g *Graph) AddEdge(u, v, cap int) {
	if g.capacity[u] == nil {
		g.capacity[u] = make(map[int]int)
	}
	if g.flow[u] == nil {
		g.flow[u] = make(map[int]int)
	}
	if g.capacity[v] == nil {
		g.capacity[v] = make(map[int]int)
	}
	if g.flow[v] == nil {
		g.flow[v] = make(map[int]int)
	}
	g.capacity[u][v] += cap // Allow multiple edges
}

// BFS finds an augmenting path using BFS
func (g *Graph) BFS(parent map[int]int) bool {
	visited := make(map[int]bool)
	queue := []int{g.source}
	visited[g.source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v, cap := range g.capacity[u] {
			if !visited[v] && cap-g.flow[u][v] > 0 { // Check if not visited and has remaining capacity
				visited[v] = true
				parent[v] = u
				if v == g.sink {
					return true
				}
				queue = append(queue, v)
			}
		}
	}
	return false
}

// FordFulkerson implements the Ford-Fulkerson method using BFS
func (g *Graph) FordFulkerson() int {
	parent := make(map[int]int)
	maxFlow := 0

	for g.BFS(parent) {
		// Find the maximum flow through the path found by BFS
		pathFlow := math.MaxInt32
		for v := g.sink; v != g.source; v = parent[v] {
			u := parent[v]
			pathFlow = min(pathFlow, g.capacity[u][v]-g.flow[u][v])
		}

		// Update residual capacities of the edges and reverse edges
		for v := g.sink; v != g.source; v = parent[v] {
			u := parent[v]
			g.flow[u][v] += pathFlow
			g.flow[v][u] -= pathFlow // Update reverse flow
		}

		maxFlow += pathFlow
		// Print the flow network after each augmenting path
		fmt.Println("Current flow after augmenting path:")
		g.PrintFlow()
	}

	return maxFlow
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PrintFlow prints the current flow and capacities in the network
func (g *Graph) PrintFlow() {
	fmt.Printf("Current Flow Network:\n")
	for u, edges := range g.capacity {
		for v, cap := range edges {
			currentFlow := g.flow[u][v]
			fmt.Printf("Edge %d -> %d | Capacity: %d | Flow: %d | Remaining Capacity: %d\n",
				u, v, cap, currentFlow, cap-currentFlow)
		}
	}
	fmt.Println()
}

// ExportDOT generates a DOT file for visualizing the flow network using Graphviz
func (g *Graph) ExportDOT(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start writing the DOT format
	fmt.Fprintf(file, "digraph FlowNetwork {\n")
	fmt.Fprintf(file, "    rankdir=LR;\n") // Left to right orientation

	// Write each edge with its flow and capacity
	for u, edges := range g.capacity {
		for v, cap := range edges {
			currentFlow := g.flow[u][v]
			fmt.Fprintf(file, "    %d -> %d [label=\"%d/%d\", color=blue];\n", u, v, currentFlow, cap)
		}
	}

	// Mark source and sink nodes
	fmt.Fprintf(file, "    %d [shape=doublecircle, style=filled, fillcolor=green];\n", g.source)
	fmt.Fprintf(file, "    %d [shape=doublecircle, style=filled, fillcolor=red];\n", g.sink)
	fmt.Fprintf(file, "}\n")

	fmt.Printf("DOT file exported to %s\n", filename)
	return nil
}

func main() {
	// Create a new graph (flow network)
	g := NewGraph(0, 5)

	// Add edges with capacities
	g.AddEdge(0, 1, 16)
	g.AddEdge(0, 2, 13)
	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 12)
	g.AddEdge(2, 1, 4)
	g.AddEdge(2, 4, 14)
	g.AddEdge(3, 2, 9)
	g.AddEdge(3, 5, 20)
	g.AddEdge(4, 3, 7)
	g.AddEdge(4, 5, 4)

	// Calculate the maximum flow
	maxFlow := g.FordFulkerson()
	fmt.Printf("The maximum possible flow from source to sink is: %d\n", maxFlow)

	// Export the graph to a DOT file for visualization
	g.ExportDOT("final_flow.dot")
}
