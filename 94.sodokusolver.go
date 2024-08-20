package main

import "fmt"

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(b [][]byte) bool {
	r, c := -1, -1

O:
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == '.' {
				r = i
				c = j
				break O
			}
		}
	}

	if r < 0 && c < 0 {
		return true
	}

	for i := 1; i <= 9; i++ {
		iByte := fmt.Sprintf("%v", i)[0]

		if isValidMove(b, iByte, r, c) {
			b[r][c] = iByte

			if solve(b) {
				return true
			}

			b[r][c] = '.'
		}
	}

	return false
}

func isValidMove(b [][]byte, value byte, row, column int) bool {
	for i := 0; i < len(b[row]); i++ {
		if b[row][i] == value {
			return false
		}
	}

	for i := 0; i < len(b); i++ {
		if b[i][column] == value {
			return false
		}
	}

	rowStart := row - row%3
	columnStart := column - column%3

	for i := rowStart; i < rowStart+3; i++ {
		for j := columnStart; j < columnStart+3; j++ {
			if b[i][j] == value {
				return false
			}
		}
	}

	return true
}
