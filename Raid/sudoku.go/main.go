package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var board [9][9]int
	for i, row := range args {
		for j, char := range row {
			if char == '.' {
				board[i][j] = 0
			} else if char >= '1' && char <= '9' {
				board[i][j] = int(char - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}
	solvedBoard, solved := solve(board)
    if !solved {
        fmt.Println("error")
        return
    }
	for _, row := range solvedBoard {
        fmt.Println(row)
    }
}
func findEmpty(board [9][9]int) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
func isValid(board [9][9]int, row, col, num int) bool {

	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	startRow := row - row%3
	startcol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startcol+j] == num {
				return false
			}
		}
	}
	return true
}
func solve(board [9][9]int) ([9][9]int, bool) {
	row, col, found := findEmpty(board)
	if !found {
		return board, true 
	}
	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num  
			newBoard, solved := solve(board)
			if solved {
				return newBoard, true
			}

			board[row][col] = 0 
		}
	}
	return board, false  
}