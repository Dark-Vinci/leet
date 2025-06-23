package main

//Platform: Leetcode
//Link: https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/description/

func onesMinusZeros(grid [][]int) [][]int {
	var (
		result = make([][]int, len(grid))
		col0   = make(map[int]int)
		col1   = make(map[int]int)
		row1   = make(map[int]int)
		row0   = make(map[int]int)
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				row0[i] += 1
				col0[j] += 1
				continue
			}

			row1[i] += 1
			col1[j] += 1
		}
	}

	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[0]))

		for j := 0; j < len(grid[0]); j++ {
			result[i][j] = row1[i] + col1[j] - row0[i] - col0[j]
		}
	}

	return result
}
