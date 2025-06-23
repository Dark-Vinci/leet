package main

//platform: Leetcode
//Link: https://leetcode.com/problems/first-completely-painted-row-or-column/

func firstCompleteIndex(arr []int, mat [][]int) int {
	var (
		y    = len(mat)
		x    = len(mat[0])
		l    = len(arr)
		cols = make(map[int]int)
		rows = make(map[int]int)
		kv   = make(map[int][2]int)
	)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			kv[mat[i][j]] = [2]int{i, j}
		}
	}

	for i := 0; i < l; i++ {
		value := arr[i]

		coords, _ := kv[value]
		yy, xx := coords[0], coords[1]

		rows[yy] += 1
		cols[xx] += 1

		if cols[xx] == y || rows[yy] == x {
			return i
		}
	}

	return -1
}
