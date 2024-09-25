package main

import (
	"fmt"
	"math"
)

// eggDrop function to compute the minimum number of drops required with visualization
func eggDrop(eggs, floors int) int {
	// Create a DP table to store the minimum number of drops
	dp := make([][]int, eggs+1)
	for i := range dp {
		dp[i] = make([]int, floors+1)
	}

	// Base cases:
	// 1. With 0 floors, no drops are needed
	// 2. With 1 floor, only 1 drop is needed
	for i := 1; i <= eggs; i++ {
		dp[i][0] = 0 // 0 floors, 0 drops
		dp[i][1] = 1 // 1 floor, 1 drop
	}
	for j := 1; j <= floors; j++ {
		dp[1][j] = j // 1 egg, j drops (linear search)
	}

	// Visualize initial base cases
	fmt.Println("Initial DP Table (Base Case):")
	printDPTable(dp, eggs, floors)

	// Fill the rest of the dp table
	for i := 2; i <= eggs; i++ {
		for j := 2; j <= floors; j++ {
			dp[i][j] = math.MaxInt32
			for x := 1; x <= j; x++ {
				// Max of two scenarios: egg breaks or doesn't break
				res := 1 + max(dp[i-1][x-1], dp[i][j-x])
				if res < dp[i][j] {
					dp[i][j] = res
				}
			}
			// Print the DP table after updating each cell
			fmt.Printf("\nAfter filling dp[%d][%d]:\n", i, j)
			printDPTable(dp, eggs, floors)
		}
	}

	// The answer will be in dp[eggs][floors]
	return dp[eggs][floors]
}

// max utility function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Utility function to print the DP table
func printDPTable(dp [][]int, eggs, floors int) {
	fmt.Println("Current DP Table:")
	for i := 0; i <= eggs; i++ {
		for j := 0; j <= floors; j++ {
			// Display "inf" for uncalculated entries (MaxInt32)
			if dp[i][j] == math.MaxInt32 {
				fmt.Printf("%4s ", "inf")
			} else {
				fmt.Printf("%4d ", dp[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	eggs := 2   // Number of eggs
	floors := 5 // Number of floors

	// Output the minimum number of drops required and visualize the process
	result := eggDrop(eggs, floors)
	fmt.Printf("\nMinimum number of drops required for %d eggs and %d floors is: %d\n", eggs, floors, result)
}
