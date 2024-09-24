package main

import (
	"fmt"
	"time"
)

// Function to solve 0/1 Knapsack Problem with DP Table Simulation
func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)

	// Create a 2D slice to store the maximum value at each n, W
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Build the DP table
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				// If the current item can fit in the knapsack, we take the max of:
				// 1. Including the current item
				// 2. Excluding the current item
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				// Otherwise, we just exclude the item
				dp[i][w] = dp[i-1][w]
			}
			// Simulate the DP table at each step
			fmt.Printf("DP Table after including item %d with weight %d and value %d:\n", i, weights[i-1], values[i-1])
			printDPTable(dp, i, capacity)
			time.Sleep(1 * time.Second) // Adding a delay to simulate step-by-step process
		}
	}

	// The bottom-right corner of the table contains the maximum value we can achieve
	return dp[n][capacity]
}

// Utility function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to print the DP table up to the current row
func printDPTable(dp [][]int, currentItem int, capacity int) {
	for i := 0; i <= currentItem; i++ {
		for w := 0; w <= capacity; w++ {
			fmt.Printf("%2d ", dp[i][w])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// Example set of weights and values
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5

	// Solving the knapsack problem with visualization
	maxValue := knapsack(weights, values, capacity)

	fmt.Printf("Maximum value in Knapsack: %d\n", maxValue)
}
