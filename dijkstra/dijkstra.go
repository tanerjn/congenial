package main

import (
	"container/heap"
	"fmt"
	"math"
)

// A struct to represent an edge in the graph
type Edge struct {
	to, weight int
}

// A struct to represent the graph
type Graph struct {
	adjacencyList map[int][]Edge
	nodes         map[int]bool
}

// Initialize a new graph
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]Edge),
		nodes:         make(map[int]bool),
	}
}

// Add an edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.adjacencyList[from] = append(g.adjacencyList[from], Edge{to, weight})
	g.adjacencyList[to] = append(g.adjacencyList[to], Edge{from, weight}) // For undirected graph
	g.nodes[from] = true
	g.nodes[to] = true
}

// A priority queue implementation (Min-Heap) to track the node with the shortest known distance
type Item struct {
	node, distance int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance // The node with smaller distance has higher priority
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra's algorithm to find shortest paths from a given source node
func Dijkstra(g *Graph, start int) map[int]int {
	distances := make(map[int]int) // Stores the shortest distances to each node
	for node := range g.adjacencyList {
		distances[node] = math.MaxInt64 // Initialize distances as infinity
	}
	distances[start] = 0

	// Priority queue to hold the nodes to explore
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{start, 0})

	// Dijkstra's core logic
	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(Item)
		currentNode := currentItem.node
		currentDistance := currentItem.distance

		// Explore neighbors of the current node
		for _, edge := range g.adjacencyList[currentNode] {
			newDistance := currentDistance + edge.weight
			if newDistance < distances[edge.to] {
				distances[edge.to] = newDistance
				heap.Push(pq, Item{edge.to, newDistance})
			}
		}
	}

	return distances
}

// Function to print a text-based dynamic visualization of the graph
func (g *Graph) PrintGraph() {
	fmt.Println("Graph Visualization:")

	// Collect all edges in the form of string lines
	lines := []string{}
	for from, edges := range g.adjacencyList {
		for _, edge := range edges {
			if from < edge.to { // Prevent duplicating edges (since it's undirected)
				lines = append(lines, fmt.Sprintf("Node %d --(%d)-- Node %d", from, edge.weight, edge.to))
			}
		}
	}

	// Print the edges
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges between nodes (from, to, weight)
	graph.AddEdge(1, 2, 4)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(3, 2, 2)
	graph.AddEdge(3, 4, 5)
	graph.AddEdge(2, 4, 1)
	graph.AddEdge(4, 1, 4)
	graph.AddEdge(5, 3, 2)

	// Dynamically print the graph
	graph.PrintGraph()

	// Run Dijkstra's algorithm for each node in the graph
	for startNode := range graph.nodes {
		shortestDistances := Dijkstra(graph, startNode)

		// Print the shortest distances from the current start node
		fmt.Printf("\nShortest distances from node %d:\n", startNode)
		for node, distance := range shortestDistances {
			if distance == math.MaxInt64 {
				fmt.Printf("Node %d: unreachable\n", node)
			} else {
				fmt.Printf("Node %d: %d\n", node, distance)
			}
		}
	}
}
