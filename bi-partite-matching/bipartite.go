package main

import (
	"fmt"
	"time"
)

// Function to check if the graph is bipartite using BFS
func isBipartiteBFS(graph [][]int) (bool, []int) {
	color := make([]int, len(graph))
	for i := range color {
		color[i] = -1 // Uncolored
	}

	for i := 0; i < len(graph); i++ {
		if color[i] == -1 { // Not colored
			if !bfs(i, graph, color) {
				return false, color
			}
		}
	}
	return true, color
}

// BFS to assign colors and check bipartiteness
func bfs(start int, graph [][]int, color []int) bool {
	queue := []int{start}
	color[start] = 0 // Start coloring with color 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if color[neighbor] == -1 { // If not colored
				color[neighbor] = 1 - color[node] // Alternate color
				queue = append(queue, neighbor)
			} else if color[neighbor] == color[node] { // Same color as current node
				return false
			}
		}
	}
	return true
}

// Function to check if the graph is bipartite using DFS
func isBipartiteDFS(graph [][]int) (bool, []int) {
	color := make([]int, len(graph))
	for i := range color {
		color[i] = -1 // Uncolored
	}

	for i := 0; i < len(graph); i++ {
		if color[i] == -1 { // Not colored
			if !dfs(i, graph, color, 0) {
				return false, color
			}
		}
	}
	return true, color
}

// DFS to assign colors and check bipartiteness
func dfs(node int, graph [][]int, color []int, currentColor int) bool {
	color[node] = currentColor

	for _, neighbor := range graph[node] {
		if color[neighbor] == -1 { // If not colored
			if !dfs(neighbor, graph, color, 1-currentColor) {
				return false
			}
		} else if color[neighbor] == color[node] { // Same color as current node
			return false
		}
	}
	return true
}

// Function to visualize the bipartite graph in the terminal
func visualizeGraphTerminal(graph [][]int, colors []int) {
	group0 := []int{}
	group1 := []int{}

	// Separate nodes into two groups based on their color
	for i, color := range colors {
		if color == 0 {
			group0 = append(group0, i)
		} else {
			group1 = append(group1, i)
		}
	}

	// Print the nodes in each group
	fmt.Println("Group 0 (Color 0):", group0)
	fmt.Println("Group 1 (Color 1):", group1)
	fmt.Println()

	// Print the edges between the two groups
	fmt.Println("Edges between Group 0 and Group 1:")
	for _, node0 := range group0 {
		for _, neighbor := range graph[node0] {
			if contains(group1, neighbor) {
				fmt.Printf("  %d -- %d\n", node0, neighbor)
			}
		}
	}
}

// Helper function to check if a slice contains an element
func contains(slice []int, elem int) bool {
	for _, val := range slice {
		if val == elem {
			return true
		}
	}
	return false
}

func main() {
	// Define the adjacency list of the graph
	graph := [][]int{
		{1, 3, 5}, // Node 0 is connected to nodes in the other set: 1, 3, 5
		{0, 2},    // Node 1 is connected to nodes 0 and 2
		{1, 3},    // Node 2 is connected to nodes 1 and 3
		{0, 2, 4}, // Node 3 is connected to nodes 0, 2, and 4
		{3, 5},    // Node 4 is connected to nodes 3 and 5
		{0, 4},    // Node 5 is connected to nodes 0 and 4
	}

	// Measure time for BFS
	startBFS := time.Now()
	bipartiteBFS, colorsBFS := isBipartiteBFS(graph)
	durationBFS := time.Since(startBFS)

	if bipartiteBFS {
		fmt.Println("The graph is bipartite (BFS).")
	} else {
		fmt.Println("The graph is not bipartite (BFS).")
	}
	fmt.Printf("BFS Time taken: %v\n", durationBFS)

	// Visualize the bipartite graph using BFS result
	visualizeGraphTerminal(graph, colorsBFS)

	// Measure time for DFS
	startDFS := time.Now()
	bipartiteDFS, colorsDFS := isBipartiteDFS(graph)
	durationDFS := time.Since(startDFS)

	if bipartiteDFS {
		fmt.Println("The graph is bipartite (DFS).")
	} else {
		fmt.Println("The graph is not bipartite (DFS).")
	}
	fmt.Printf("DFS Time taken: %v\n", durationDFS)

	// Visualize the bipartite graph using DFS result
	visualizeGraphTerminal(graph, colorsDFS)
}
