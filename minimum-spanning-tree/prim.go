package main

import (
	"container/heap"
	"fmt"
)

// Edge represents an edge in the graph with a destination vertex and a weight
type Edge struct {
	to, weight int
}

// Item represents a vertex in the priority queue
type Item struct {
	node, priority int
	index          int
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority (smallest weight)
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Prim's algorithm to find Minimum Spanning Tree with simulation
func primMST(graph [][]Edge, n int) int {
	// Array to check if a vertex is already included in the MST
	inMST := make([]bool, n)

	// Priority queue to pick the smallest edge
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Start from the vertex 0
	heap.Push(pq, &Item{node: 0, priority: 0})

	totalWeight := 0

	fmt.Println("Starting Prim's Algorithm Simulation...")
	fmt.Println("========================================")
	fmt.Println("Adding edges to the priority queue:")

	for pq.Len() > 0 {
		// Get the edge with the smallest weight
		item := heap.Pop(pq).(*Item)
		node := item.node
		weight := item.priority

		// If the node is already in MST, continue
		if inMST[node] {
			continue
		}

		// Include the node in MST
		inMST[node] = true
		totalWeight += weight

		// Simulate edge inclusion
		if weight != 0 {
			fmt.Printf("Selected Edge: (Node %d) with weight %d\n", node, weight)
		}

		fmt.Printf("Current MST Weight: %d\n", totalWeight)
		fmt.Println("----------------------------------------")

		// Look at all edges from this node
		for _, edge := range graph[node] {
			if !inMST[edge.to] {
				fmt.Printf("Considering Edge: (%d -> %d) with weight %d\n", node, edge.to, edge.weight)
				heap.Push(pq, &Item{node: edge.to, priority: edge.weight})
			}
		}

		fmt.Println("========================================")
	}

	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", totalWeight)
	return totalWeight
}

func main() {
	// Graph represented as adjacency list
	// Each node has a list of edges with the destination node and weight
	graph := [][]Edge{
		0: {{to: 1, weight: 2}, {to: 3, weight: 6}},
		1: {{to: 0, weight: 2}, {to: 2, weight: 3}, {to: 3, weight: 8}, {to: 4, weight: 5}},
		2: {{to: 1, weight: 3}, {to: 4, weight: 7}},
		3: {{to: 0, weight: 6}, {to: 1, weight: 8}, {to: 4, weight: 9}},
		4: {{to: 1, weight: 5}, {to: 2, weight: 7}, {to: 3, weight: 9}},
	}

	n := len(graph) // Number of vertices
	totalWeight := primMST(graph, n)
	fmt.Printf("Total weight of the Minimum Spanning Tree: %d\n", totalWeight)
}
