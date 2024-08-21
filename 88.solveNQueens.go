package main

import (
	"bytes"
	"fmt"
)

func solveNQueens(n int) [][]string {
	var (
		result       = make([][]string, 0)
		columns      = make([]int, n)
		diagonals    = make([]int, n*2)
		oppDiagonals = make([]int, n*2)
		board        = make([][]byte, n)

		dfs func(int)
	)

	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	dfs = func(i int) {
		if i == n {
			curr := make([]string, n)

			for j := range curr {
				curr[j] = string(board[j])
			}

			result = append(result, curr)
		}

		for j := 0; j < n; j++ {
			if columns[j]+diagonals[i+j]+oppDiagonals[n-i+j] == 0 {
				columns[j], diagonals[i+j], oppDiagonals[n-i+j] = 1, 1, 1
				board[i][j] = 'Q'
				dfs(i + 1)
				board[i][j] = '.'
				columns[j], diagonals[i+j], oppDiagonals[n-i+j] = 0, 0, 0
			}
		}
	}

	dfs(0)

	return result
}

func myNQueens(n int) [][]string {
	result := make([][]string, 0)

	var solution func(board [][]byte, x, y int)

	m := make(map[string]struct{})

	solution = func(board [][]byte, X, Y int) {
		boardStr := ""
		for i := 0; i < n; i++ {
			boardStr += string(board[i])
		}

		if _, ok := m[fmt.Sprintf("%v-%v-%v", X, Y, boardStr)]; ok {
			return
		} else {
			m[fmt.Sprintf("%v-%v-%v", X, Y, boardStr)] = struct{}{}
		}

		// GET A CLONE OF THE BOARD
		cpBoard := make([][]byte, n)
		for i := 0; i < n; i++ {
			cpBoard[i] = bytes.Clone(board[i])
		}

		// SPAWN FOR ALL COLUMN IN SAME ROW
		for i := X + 1; i < n; i++ {
			solution(board, i, Y)
		}

		if cpBoard[Y][X] == '.' {
			return
		}

		// STAR THE ROW
		for i := 0; i < n; i++ {
			cpBoard[Y][i] = '.'
		}

		//STAR THE COLUMN
		for i := 0; i < n; i++ {
			cpBoard[i][X] = '.'
		}

		// UP, RIGHT
		for i, j := X+1, Y-1; i < n && j >= 0; i, j = i+1, j-1 {
			cpBoard[j][i] = '.'
		}

		// UP, LEFT
		for i, j := X-1, Y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			cpBoard[j][i] = '.'
		}

		// DOWN, LEFT
		for i, j := X-1, Y+1; i >= 0 && j < n; i, j = i-1, j+1 {
			cpBoard[j][i] = '.'
		}

		// DOWN, RIGHT
		for i, j := X+1, Y+1; i < n && j < n; i, j = i+1, j+1 {
			cpBoard[j][i] = '.'
		}

		cpBoard[Y][X] = 'Q'

		if Y == n-1 && countQueen(cpBoard) == n {
			res := make([]string, 0)

			for i := 0; i < n; i++ {
				res = append(res, string(cpBoard[i]))
			}

			result = append(result, res)
			return
		}

		// SOLVE FOR ALL OTHER ROW
		for i := Y + 1; i < n; i++ {
			solution(cpBoard, 0, i)
		}
	}

	board := make([][]byte, n)

	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '_'
		}
	}

	solution(board, 0, 0)

	return result
}

func countQueen(b [][]byte) int {
	res := 0

Loop:
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == 'Q' {
				res += 1
				continue Loop
			}
		}
	}

	return res
}
