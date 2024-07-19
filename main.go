package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var table [][]int
	rows, cols := 9, 9
	tries := 0
	attemptedTables := 0
	valid := false

	for !valid {
		attemptedTables++
		table = [][]int{
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		}
		valid = true
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				for !isCellValid(&table, i, j) {
					table[i][j] = rand.Intn(9) + 1
					tries++
					if tries > 200 {
						valid = false
						break
					}
				}
				tries = 0
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}
	}

	fmt.Println("generated: ", attemptedTables)
	showTable(&table, rows, cols)
}

func isCellValid(table *[][]int, row int, col int) bool {
	if (*table)[row][col] == -1 {
		return false
	}
	validRow, validCol, validBlock := isValidInRow(table, row, col), isValidInCol(table, row, col), isValidInBlock(table, row, col)
	return validRow && validCol && validBlock
}

func isValidInRow(table *[][]int, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}

		if (*table)[row][i] == (*table)[row][col] {
			return false
		}
	}

	return true
}

func isValidInCol(table *[][]int, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}

		if (*table)[i][col] == (*table)[row][col] {
			return false
		}
	}

	return true
}

func isValidInBlock(table *[][]int, row int, col int) bool {
	blockRow, blockCol := 0, 0

	if col >= 6 {
		blockCol = 6
	} else if col >= 3 {
		blockCol = 3
	} else {
		blockCol = 0
	}

	if row >= 6 {
		blockRow = 6
	} else if row >= 3 {
		blockRow = 3
	} else {
		blockRow = 0
	}

	for i := blockRow; i < blockRow+3; i++ {
		for j := blockCol; j < blockCol+3; j++ {
			if i == row && j == col {
				continue
			}

			if (*table)[i][j] == (*table)[row][col] {
				return false
			}
		}
	}

	return true
}

func showTable(table *[][]int, rows int, cols int) {
	fmt.Println("————————————————————————————")
	for i := 0; i < rows; i++ {
		fmt.Print("| ")
		for j := 0; j < cols; j++ {
			if (*table)[i][j] == -1 {
				fmt.Print("- ")
			} else {
				fmt.Printf("%v ", (*table)[i][j])
			}
			if (j+1)%3 == 0 {
				fmt.Print(" | ")
			}
		}
		fmt.Println("")
		if (i+1)%3 == 0 {
			fmt.Println("————————————————————————————")
		}
	}
}
