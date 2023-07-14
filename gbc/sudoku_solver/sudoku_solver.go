package main

import (
	"fmt"
	_ "log"
	_ "os"
	_ "os/exec"
	"strconv"
)

// ... for quick debugging
var p = fmt.Println

/*
SUDOKO RULES:
1. 9 by 9 square
2. Each row and column must contain number 1-9
3. each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// When trying to solve problems,
// 1. start ot think about what funcs we're going to need.
// 2, Then test each new func to make sure it does what you need it to do.

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	// puzz_cols := len(puzz[0])
	// cycles through rows in the puzz
	for _, row := range puzz {
		// cycles through the columns of each row
		for _, col := range row {
			fmt.Print(strconv.Itoa(col) + " ")
		}
		fmt.Println()
	}
}

// return a valid row/col for an empty space
func getEmptySpace(puzz [][]int) (int, int) {
	// cycles through rows in the puzz
	for i, row := range puzz {
		// cycles through the columns of each row
		for j, col := range row {
			if col == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNumValid(puzz [][]int, guess, row, col int) bool {
	for index := range puzz {
		if puzz[row][index] == guess && col != index {
			return false
		}
	}
	return true
}

func main() {
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	displayBoard(puzz)
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is solved!")
	}
	
	fmt.Println(isNumValid(puzz,7,0,0)) 

}
