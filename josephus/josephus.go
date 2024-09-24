package main

import (
	"fmt"
	"time"
)

// Function to find the position of the last remaining person iteratively and display steps
func josephusIterative(n, k int) int {
	people := make([]int, n)
	for i := 0; i < n; i++ {
		people[i] = i + 1 // Populate list with people numbered 1 to n
	}

	index := 0 // Start at the first person
	fmt.Println("Initial circle:", people)

	for len(people) > 1 {
		// Find the index of the next person to eliminate
		index = (index + k - 1) % len(people)

		// Remove the eliminated person from the list
		people = append(people[:index], people[index+1:]...)

		// // Show the new list after elimination
		fmt.Println("Remaining people:", people)
	}

	return people[0]
}

// Function to eliminate people recursively and display steps
func josephusRecursive(people []int, k int, index int) (int, []int) {
	// Base case: If only one person remains
	if len(people) == 1 {
		return people[0], people
	}

	// Find the index of the next person to eliminate
	index = (index + k - 1) % len(people)

	// Remove the eliminated person from the list
	people = append(people[:index], people[index+1:]...)

	// // Show the remaining people
	fmt.Println("Remaining people:", people)

	// Recursive call
	return josephusRecursive(people, k, index)
}

func main() {
	n := 40 // Number of people
	k := 2  // Every k-th person is eliminated

	// Measure the time for the iterative solution
	startIterative := time.Now()
	survivorIterative := josephusIterative(n, k)
	durationIterative := time.Since(startIterative)

	// Output for iterative solution
	fmt.Printf("The last person standing (Iterative) is at position %d\n", survivorIterative)
	fmt.Printf("Iterative solution took: %v\n", durationIterative)

	// Create a slice of people for the recursive solution
	people := make([]int, n)
	for i := 0; i < n; i++ {
		people[i] = i + 1 // Populate list with people numbered 1 to n
	}

	// Measure the time for the recursive solution
	startRecursive := time.Now()
	survivorRecursive, _ := josephusRecursive(people, k, 0)
	durationRecursive := time.Since(startRecursive)

	// Output for recursive solution
	fmt.Printf("The last person standing (Recursive) is at position %d\n", survivorRecursive)
	fmt.Printf("Recursive solution took: %v\n", durationRecursive)
}
