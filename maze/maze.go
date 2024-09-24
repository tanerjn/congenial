package main

import (
	"container/list"
	"fmt"
)

// Directions for movement: up, down, left, right
var directions = [][2]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

// BFS function to find the shortest path in a maze
func bfs(maze [][]int, start, end [2]int) ([][2]int, bool) {
	rows := len(maze)
	cols := len(maze[0])
	queue := list.New()
	visited := make(map[[2]int]bool)
	path := make(map[[2]int][2]int)

	// Start BFS from the starting position
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).([2]int)

		// Check if we've reached the end
		if current == end {
			return reconstructPath(path, start, end), true
		}

		// Explore the neighboring cells
		for _, direction := range directions {
			nextRow := current[0] + direction[0]
			nextCol := current[1] + direction[1]
			nextCell := [2]int{nextRow, nextCol}

			// Check if the next cell is within bounds, not a wall, and not visited
			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && maze[nextRow][nextCol] == 0 && !visited[nextCell] {
				visited[nextCell] = true
				queue.PushBack(nextCell)
				path[nextCell] = current // Keep track of the path
			}
		}
	}
	return nil, false // No path found
}

// Function to reconstruct the path from start to end
func reconstructPath(path map[[2]int][2]int, start, end [2]int) [][2]int {
	var result [][2]int
	for at := end; at != start; at = path[at] {
		result = append(result, at)
	}
	result = append(result, start)
	// Reverse the path to get it from start to end
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Function to visualize the maze with the path
func visualizeMaze(maze [][]int, path [][2]int) {
	// Create a copy of the maze to mark the path
	mazeWithPath := make([][]int, len(maze))
	for i := range maze {
		mazeWithPath[i] = make([]int, len(maze[i]))
		copy(mazeWithPath[i], maze[i])
	}

	// Mark the path with a special character (2 in this case)
	for _, p := range path {
		mazeWithPath[p[0]][p[1]] = 2
	}

	// Print the maze with the path
	fmt.Println("Maze with the shortest path (0: path, 1: wall, 2: path taken):")
	for _, row := range mazeWithPath {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print("  ") // Print space for path
			} else if cell == 1 {
				fmt.Print("█ ") // Print wall
			} else if cell == 2 {
				fmt.Print("* ") // Print path taken
			}
		}
		fmt.Println()
	}
}

func main() {
	// Example maze
	maze := [][]int{
		{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	}

	start := [2]int{0, 0} // Start at top-left corner
	end := [2]int{5, 13}  // End at bottom-right corner

	path, found := bfs(maze, start, end)
	if found {
		fmt.Println("Shortest path found:")
		for _, p := range path {
			fmt.Printf("(%d, %d) ", p[0], p[1])
		}
		fmt.Println()

		// Visualize the maze with the path
		visualizeMaze(maze, path)
	} else {
		fmt.Println("No path found.")
	}
}
