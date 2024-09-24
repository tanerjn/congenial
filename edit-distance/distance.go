package main

import (
	"fmt"
)

func editDistance(s1, s2 string) (int, [][]int) {
	m, n := len(s1), len(s2)
	//Create slice to store the distances
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	//initialise the dp array
	for i := 0; i <= m; i++ {
		dp[i][0] = i // Cost of deleteting all from s1
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // Cost of deleting all from s2
	}

	printDP(dp)

	//fill the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] // no cost if chars are same
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], // Deletion
					dp[i][j-1],   // Insertion
					dp[i-1][j-1]) // Substitution
			}
			// Print the state of the dp array after each trial
			fmt.Printf("After comparing '%s' and '%s':\n", s1[:i], s2[:j])
			printDP(dp)
		}
	}
	return dp[m][n], dp
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func printDP(dp [][]int) {
	for _, row := range dp {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	s1 := "coasters"
	s2 := "bungeear"
	distance, _ := editDistance(s1, s2)
	fmt.Printf("Minimum Edit Distance between '%s' and '%s': %d\n", s1, s2, distance)

}
