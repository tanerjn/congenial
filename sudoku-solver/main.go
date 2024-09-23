package main

import (
	"fmt"
	"math"
)

const N = 9

// Function to print the Sudoku grid
func printGrid(grid [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Function to check if a number can be placed in a given cell
func isSafe(grid [N][N]int, row, col, num int) bool {
	// Check the row and column
	for x := 0; x < N; x++ {
		if grid[row][x] == num || grid[x][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := row / int(math.Sqrt(N)) * int(math.Sqrt(N))
	startCol := col / int(math.Sqrt(N)) * int(math.Sqrt(N))

	for i := 0; i < int(math.Sqrt(N)); i++ {
		for j := 0; j < int(math.Sqrt(N)); j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Function to solve the 9x9 Sudoku using backtracking
func solveSudoku(grid [N][N]int) bool {
	row, col := -1, -1
	isEmpty := false

	// Find an empty cell (represented by 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				row = i
				col = j
				isEmpty = true
				break
			}
		}
		if isEmpty {
			break
		}
	}

	// No empty cell left, puzzle is solved
	if !isEmpty {
		return true
	}

	// Try numbers from 1 to 9
	for num := 1; num <= N; num++ {
		if isSafe(grid, row, col, num) {
			// Print grid before placing a number
			fmt.Printf("Trying %d at position (%d, %d):\n", num, row, col)
			grid[row][col] = num
			printGrid(grid)

			// Recursively solve the rest of the puzzle
			if solveSudoku(grid) {
				return true
			}

			// Backtrack if not successful
			grid[row][col] = 0
			fmt.Printf("Backtracking from (%d, %d):\n", row, col)
			printGrid(grid)
		}
	}

	return false
}

func main() {
	// Example 9x9 Sudoku grid (0 represents empty cells)
	grid := [N][N]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Initial Sudoku grid:")
	printGrid(grid)

	if solveSudoku(grid) {
		fmt.Println("Sudoku solved successfully!")
	} else {
		fmt.Println("No solution exists")
	}
}
