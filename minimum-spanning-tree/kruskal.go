package main

import (
	"fmt"
	"sort"
)

// Edge represents an edge with two vertices and a weight
type Edge struct {
	from, to, weight int
}

// A Graph represents a collection of edges
type Graph struct {
	edges []Edge
	V     int // Number of vertices
}

// Disjoint Set (Union-Find) structure to detect cycles
type UnionFind struct {
	parent, rank []int
}

// Initialize UnionFind with n elements (0 to n-1)
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

// Find finds the root of the set containing x, with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges two sets if they are not already in the same set
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// Union by rank to keep the tree flat
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		return true
	}
	return false
}

// Kruskal's Algorithm to find MST
func kruskalMST(graph Graph) int {
	// Sort all the edges in increasing order by weight
	sort.Slice(graph.edges, func(i, j int) bool {
		return graph.edges[i].weight < graph.edges[j].weight
	})

	// Create a Union-Find structure to manage sets of vertices
	uf := NewUnionFind(graph.V)

	totalWeight := 0
	mstEdges := 0

	fmt.Println("Kruskal's Algorithm Simulation")
	fmt.Println("=============================")

	// Iterate through sorted edges
	for _, edge := range graph.edges {
		// Check if the current edge forms a cycle
		if uf.Union(edge.from, edge.to) {
			// Edge is included in the MST
			fmt.Printf("Selected Edge: (%d -> %d) with weight %d\n", edge.from, edge.to, edge.weight)
			totalWeight += edge.weight
			mstEdges++

			// If we already have V-1 edges in the MST, we are done
			if mstEdges == graph.V-1 {
				break
			}
		} else {
			// Edge would form a cycle, so it's ignored
			fmt.Printf("Ignored Edge: (%d -> %d) with weight %d (would form cycle)\n", edge.from, edge.to, edge.weight)
		}
	}

	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", totalWeight)
	return totalWeight
}

func main() {
	// Define a graph with 5 vertices and edges with weights
	graph := Graph{
		V: 5,
		edges: []Edge{
			{from: 0, to: 1, weight: 2},
			{from: 0, to: 3, weight: 6},
			{from: 1, to: 2, weight: 3},
			{from: 1, to: 3, weight: 8},
			{from: 1, to: 4, weight: 5},
			{from: 2, to: 4, weight: 7},
			{from: 3, to: 4, weight: 9},
		},
	}

	// Run Kruskal's Algorithm and simulate the process
	kruskalMST(graph)
}
