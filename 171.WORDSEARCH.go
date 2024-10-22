package main

import "cmp"

func exist(board [][]byte, word string) bool {
	type node struct {
		x, y int
	}

	var (
		l1, l2, lw = len(board), len(board[0]), len(word)
		visited    = make(map[node]struct{})
		dfs        func(i, j, index int) bool
	)

	dfs = func(i, j, index int) bool {
		n := node{x: i, y: j}

		if j >= l1 || i >= l2 || i < 0 || j < 0 {
			return false
		}

		if _, ok := visited[n]; ok {
			return false
		}

		if word[index] != board[j][i] {
			return false
		}

		if index == lw-1 {
			return true
		}

		index++

		visited[n] = struct{}{}

		// UP, DOWN, LEFT, RIGHT
		l, r := dfs(i-1, j, index), dfs(i+1, j, index)
		u, d := dfs(i, j-1, index), dfs(i, j+1, index)

		delete(visited, n)

		return cmp.Or(l, r, u, d)
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			if board[j][i] == word[0] {
				clear(visited)
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
