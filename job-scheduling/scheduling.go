package main

import (
	"fmt"
	"sort"
	"time"
)

// Job struct to represent each job with start time, end time, and profit
type Job struct {
	start  int
	end    int
	profit int
}

// Binary search function to find the latest job that doesn't overlap
func binarySearch(jobs []Job, index int) int {
	low, high := 0, index-1

	for low <= high {
		mid := (low + high) / 2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				low = mid + 1
			} else {
				return mid
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Function to find the maximum profit of non-overlapping jobs and visualize the process over time
func jobSchedulingWithTimeVisualization(jobs []Job) int {
	// Sort jobs by their end time
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	// Create DP array to store the maximum profit up to each job
	n := len(jobs)
	dp := make([]int, n)
	dp[0] = jobs[0].profit

	// To store the optimal job sequence
	scheduledJobs := make([]int, n)

	// Fill the DP array
	for i := 1; i < n; i++ {
		// Option 1: Include the current job and add profit from the last non-overlapping job
		profitIncludingCurrent := jobs[i].profit
		lastNonOverlapping := binarySearch(jobs, i)
		if lastNonOverlapping != -1 {
			profitIncludingCurrent += dp[lastNonOverlapping]
		}

		// Option 2: Exclude the current job
		profitExcludingCurrent := dp[i-1]

		// Choose the maximum of both options
		if profitIncludingCurrent > profitExcludingCurrent {
			dp[i] = profitIncludingCurrent
			scheduledJobs[i] = i
		} else {
			dp[i] = profitExcludingCurrent
			scheduledJobs[i] = scheduledJobs[i-1]
		}
	}

	// Reconstruct the optimal job schedule
	finalSchedule := []Job{}
	index := n - 1
	for index >= 0 {
		if scheduledJobs[index] == index {
			finalSchedule = append(finalSchedule, jobs[index])
			index = binarySearch(jobs, index)
		} else {
			index--
		}
	}

	// Reverse the finalSchedule to get the correct order of job execution
	for i, j := 0, len(finalSchedule)-1; i < j; i, j = i+1, j-1 {
		finalSchedule[i], finalSchedule[j] = finalSchedule[j], finalSchedule[i]
	}

	// Visualize time-based task execution
	fmt.Println("Starting time-based job visualization...")

	currentTime := 0
	for _, job := range finalSchedule {
		// Simulate idle time if there's a gap between jobs
		if currentTime < job.start {
			for idle := currentTime; idle < job.start; idle++ {
				fmt.Printf("Time %d: Idle\n", idle)
				time.Sleep(1 * time.Second) // Simulate 1 second per time unit
			}
		}
		// Simulate job execution
		for t := job.start; t < job.end; t++ {
			fmt.Printf("Time %d: Working on Job (Start: %d, End: %d, Profit: %d)\n", t, job.start, job.end, job.profit)
			time.Sleep(1 * time.Second) // Simulate 1 second per time unit
		}
		currentTime = job.end
	}

	fmt.Println("All jobs completed!")
	fmt.Println()

	// The maximum profit is in the last entry of the DP array
	return dp[n-1]
}

// Main function
func main() {
	// Example set of jobs: (start_time, end_time, profit)
	jobs := []Job{
		{1, 3, 50},
		{2, 5, 20},
		{4, 6, 70},
		{6, 7, 60},
		{5, 8, 30},
		{7, 9, 80},
	}

	// Output the maximum profit with visualization
	maxProfit := jobSchedulingWithTimeVisualization(jobs)
	fmt.Printf("Maximum profit: %d\n", maxProfit)
}
