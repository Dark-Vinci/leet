package main

//Platform: Leetcode
//Link: https://leetcode.com/problems/special-positions-in-a-binary-matrix/description/

func numSpecial(mat [][]int) int {
	var (
		y      = len(mat)
		x      = len(mat[0])
		cols1  = make(map[int]int)
		rows1  = make(map[int]int)
		result = 0
	)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if mat[i][j] == 1 {
				cols1[j] += 1
				rows1[i] += 1
				continue
			}

			cols1[j] += 0
			rows1[i] += 0
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if cols1[j] == 1 && rows1[i] == 1 && mat[i][j] == 1 {
				result++
			}
		}
	}

	return result
}
