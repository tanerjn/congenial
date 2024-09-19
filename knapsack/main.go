package main

import (
	"fmt"
)

// Knapsack function: returns the maximum value that can be carried in the knapsack
func knapsack(values, weights []int, W int) int {
	// Number of items
	n := len(values)

	// Create a 2D slice to store the maximum value for each weight limit
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	// Fill the dp table
	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if weights[i-1] <= w {
				// Maximize the value between including the current item or excluding it
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				// If the item can't be included, carry forward the previous value
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// The bottom-right corner of the dp table contains the maximum value
	return dp[n][W]
}

// Utility function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example input: values and weights of items
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	// Knapsack capacity

	for W := 10; W <= 300; W += 10 {
		maxValue := knapsack(values, weights, W)
		fmt.Printf("Maximum value for knapsack capacity %d = %d\n", W, maxValue)
	}
}
