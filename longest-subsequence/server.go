package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var (
	nums []int
	lock sync.Mutex
)

// Function to find the length of the longest increasing subsequence
func longestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := []int{}
	for _, num := range nums {
		pos := binarySearch(tail, num)
		if pos < len(tail) {
			tail[pos] = num
		} else {
			tail = append(tail, num)
		}
	}

	return len(tail)
}

// Binary search to find the position where `target` should be placed
func binarySearch(tail []int, target int) int {
	left, right := 0, len(tail)-1
	for left <= right {
		mid := left + (right-left)/2
		if tail[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// Function to handle client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid number received:", line)
			continue
		}

		lock.Lock()
		nums = append(nums, num)
		lock.Unlock()

		lock.Lock()
		fmt.Printf("Received: %d\n", num)
		fmt.Printf("Current Array: %v\n", nums)
		fmt.Printf("Length of the longest increasing subsequence: %d\n\n", longestIncreasingSubsequence(nums))
		lock.Unlock()
	}
}

// Main function to start the server
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
