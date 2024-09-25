package main

import (
	"fmt"
)

// subsetSum determines if there exists a subset of `set` with sum equal to `target` and prints the DP table
func subsetSum(set []int, target int) bool {
	n := len(set)

	// dp[i][j] will be true if a subset of the first i elements has a sum of j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	// Base case: A sum of 0 is always possible with an empty subset
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	// Fill the dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if set[i-1] > j {
				// If the current element is greater than the sum, we can't include it
				dp[i][j] = dp[i-1][j]
			} else {
				// We can either include the element or exclude it
				dp[i][j] = dp[i-1][j] || dp[i-1][j-set[i-1]]
			}
		}
	}

	// Print the DP table
	fmt.Println("DP Table:")
	for i := 0; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if dp[i][j] {
				fmt.Print("T ")
			} else {
				fmt.Print("F ")
			}
		}
		fmt.Println()
	}

	// Check if there's a solution
	if dp[n][target] {
		fmt.Println("\nSubset exists with the given sum!")
		printSubset(dp, set, n, target)
		return true
	} else {
		fmt.Println("No subset exists with the given sum.")
		return false
	}
}

// printSubset prints the actual subset that sums to the target value
func printSubset(dp [][]bool, set []int, n int, target int) {
	fmt.Println("The subset that sums to", target, "is:")

	subset := []int{}

	// Trace back the DP table to find which elements are part of the subset
	for i := n; i > 0 && target > 0; i-- {
		// If the value is the same as the row above, this element was not included in the subset
		if dp[i][target] && !dp[i-1][target] {
			subset = append(subset, set[i-1])
			target -= set[i-1] // Reduce the target
		}
	}

	// Print the subset
	for i := len(subset) - 1; i >= 0; i-- { // Print in reverse order to maintain original sequence
		fmt.Print(subset[i], " ")
	}
	fmt.Println()
}

func main() {
	// Example set and target
	set := []int{2, 4, 4}
	target := 10

	subsetSum(set, target)
}

//Row 0: With no elements, only the sum 0 is achievable.
//Row 1: With element 3, can achieve the sum 3.
//Row 2: With elements {3, 4}, can achieve sums 3, 4, and 7.
//Row 3: With elements {3, 4, 5}, can achieve sums 3, 4, 5, 7, 8, and 9.
