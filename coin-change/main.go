package main

import (
	"fmt"
	"math"
)

// Function to find the minimum number of coins needed to make up the given amount
func coinChange(coins []int, amount int) int {
	// Create a DP array and initialize with a large number
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// Update DP array for each coin
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1 // Amount cannot be formed with given coins
	}
	return dp[amount]
}

// Helper function to find the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{10, 20, 50} // Coin denominations
	amount := 90               // Total amount to make up

	result := coinChange(coins, amount)
	if result != -1 {
		fmt.Printf("Minimum number of coins needed: %d\n", result)
	} else {
		fmt.Println("Amount cannot be formed with the given coins.")
	}
}
