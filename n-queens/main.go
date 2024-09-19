package main

import (
	"fmt"
	"time"
)

// Function to print the chessboard
func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Function to check if it's safe to place a queen at (row, col)
func isSafe(board [][]int, row, col, n int) bool {
	// Check this column on upper side
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Check upper right diagonal
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

// Function to solve the N-Queens problem using backtracking
func solveNQueens(board [][]int, row, n int) bool {
	if row >= n {
		return true
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col, n) {
			board[row][col] = 1 // Place queen
			fmt.Printf("Placing queen at (%d, %d):\n", row, col)
			printBoard(board)
			time.Sleep(100 * time.Millisecond) // Pause to show the trial (adjust as needed)

			if solveNQueens(board, row+1, n) {
				return true
			}

			board[row][col] = 0 // Remove queen (backtrack)
			fmt.Printf("Backtracking from (%d, %d):\n", row, col)
			printBoard(board)
			time.Sleep(100 * time.Millisecond) // Pause to show the backtracking (adjust as needed)
		}
	}

	return false
}

func main() {
	n := 8 // Size of the chessboard (8 for an 8x8 board)

	// Initialize the board
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	fmt.Println("Starting N-Queens Solver...")
	if solveNQueens(board, 0, n) {
		fmt.Println("Solution found:")
		printBoard(board)
	} else {
		fmt.Println("No solution exists.")
	}
}
