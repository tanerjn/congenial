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
		// If all cities have been visited, return 0 since we do not need to return to the starting city
		return 0
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
func printPath(startCity int, path [][]int, n int) {
	mask := 1 << startCity // Start with the starting city visited
	pos := startCity
	visited := make(map[int]bool) // Track visited cities

	// Print the starting city
	fmt.Printf("Tour starting from %c: %c", 'A'+startCity, 'A'+startCity)
	visited[startCity] = true

	// Iterate through the path
	for i := 0; i < n-1; i++ {
		next := path[mask][pos]

		// Check if next city is valid
		if next == -1 {
			fmt.Println("Error: Path contains invalid city.")
			return
		}

		fmt.Printf(" -> %c", 'A'+next)
		mask |= (1 << next) // Update the bitmask
		pos = next          // Move to the next city
		visited[next] = true
	}

	// Print the final city (end of the path)
	fmt.Println()
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

	// Iterate through each city as the starting city
	for startCity := 0; startCity < n; startCity++ {
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

		// Solve TSP starting from the current city with only that city visited
		minCost := tsp(1<<startCity, startCity, dist, dp, path, n)

		// Print the minimum cost and path for the current starting city
		fmt.Printf("Minimum cost for tour starting from %c: %d\n", 'A'+startCity, minCost)
		printPath(startCity, path, n)
		fmt.Println()
	}
}
