package main

import (
	"fmt"
	"math"
)

// Function to solve the Maximum Subarray Problem using Kadane's Algorithm
func maxSubArray(nums []int) (int, []int) {
	// Initialize current sum and max sum
	currentSum := 0
	maxSum := math.MinInt32

	// Variables to track the start and end of the max subarray
	start, end, tempStart := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		// Add current element to current sum
		currentSum += nums[i]

		// Update max sum if current sum is greater
		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart // Set the start index of the max subarray
			end = i           // Set the end index of the max subarray
		}

		// If current sum becomes negative, reset it and move the temporary start pointer
		if currentSum < 0 {
			currentSum = 0
			tempStart = i + 1
		}
	}

	// Return the max sum and the subarray that gives the max sum
	return maxSum, nums[start : end+1]
}

func main() {
	// Example array
	arr := []int{-2, 1, -3, -11, -5, 3, 2, -1, 4, -1, 2, 1, 0, 0, 1, -5, 4}

	// Get the maximum sum and the subarray
	maxSum, subArray := maxSubArray(arr)

	// Print the results
	fmt.Printf("Maximum Subarray Sum: %d\n", maxSum)
	fmt.Printf("Maximum Subarray: %v\n", subArray)
}
