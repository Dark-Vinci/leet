package Graphs

func numberOfRightTriangles(grid [][]int) int64 {
	var (
		result int64
		l1, l2 = len(grid), len(grid[0])
	)

	rows, cols := make([]int, l1), make([]int, l2)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 {
				rows[i] += 1
				cols[j] += 1
			}
		}
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 {
				r, c := rows[i]-1, cols[j]-1

				if r >= 0 && c >= 0 {
					result += int64(r * c)
				}
			}
		}
	}

	return result
}
