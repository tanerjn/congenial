package main

import (
	"fmt"
	"math"
)

// Helper function to find the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function to solve TSP with dynamic programming and bitmasking
func tsp(mask int, pos int, dist [][]int, dp [][]int, path [][]int, n int) int {
	if mask == (1<<n)-1 {
		// If all cities have been visited, return cost to go back to starting city
		return dist[pos][0]
	}

	// If this subproblem has been solved before, return the stored result
	if dp[mask][pos] != -1 {
		return dp[mask][pos]
	}

	// Initialize the answer to a large value
	ans := math.MaxInt32
	nextCity := -1

	// Try to go to every other city and take the minimum path
	for city := 0; city < n; city++ {
		// Check if the city has already been visited (bitmask check)
		if (mask & (1 << city)) == 0 {
			// Calculate new cost to visit this city
			newAns := dist[pos][city] + tsp(mask|(1<<city), city, dist, dp, path, n)
			if newAns < ans {
				ans = newAns
				nextCity = city // Keep track of the next city in the optimal path
			}
		}
	}

	// Store the result and the next city to visit in the path
	dp[mask][pos] = ans
	path[mask][pos] = nextCity

	return ans
}

// Function to print the path in the TSP tour
func printPath(path [][]int, n int) {
	mask := 1 // Start with city A visited (bitmask 00000001)
	pos := 0  // Start at city A (index 0)

	fmt.Print("Tour: A") // Start with city A
	for i := 0; i < n-1; i++ {
		next := path[mask][pos]
		fmt.Printf(" -> %c", 'A'+next) // Print the next city in alphabetical order
		mask |= (1 << next)            // Mark the next city as visited
		pos = next                     // Move to the next city
	}
	fmt.Println(" -> A") // Return to the starting city A
}

func main() {
	// Example distance matrix for 8 cities (A, B, C, D, E, F, G, H)
	dist := [][]int{
		{0, 10, 15, 20, 25, 30, 35, 40}, // Distances from A
		{10, 0, 35, 25, 30, 20, 50, 45}, // Distances from B
		{15, 35, 0, 30, 55, 65, 45, 50}, // Distances from C
		{20, 25, 30, 0, 30, 35, 55, 60}, // Distances from D
		{25, 30, 55, 30, 0, 50, 40, 70}, // Distances from E
		{30, 20, 65, 35, 50, 0, 75, 80}, // Distances from F
		{35, 50, 45, 55, 40, 75, 0, 90}, // Distances from G
		{40, 45, 50, 60, 70, 80, 90, 0}, // Distances from H
	}

	n := len(dist)

	// Initialize memoization and path tables
	dp := make([][]int, 1<<n)
	path := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		path[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
			path[i][j] = -1
		}
	}

	// Solve TSP starting from city 0 (A) with no other cities visited (mask = 00000001)
	minCost := tsp(1, 0, dist, dp, path, n)

	// Print the minimum cost
	fmt.Printf("Minimum cost: %d\n", minCost)

	// Print the path
	printPath(path, n)
}
