package main

func orangesRotting(grid [][]int) int {
	var (
		q      = make([][]int, 0)
		dirs   = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		good   = 0
		result = 0
		height = len(grid)
		width  = len(grid[0])
	)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				good++
			} else if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 && good > 0 {
		qLen := len(q)
		result++

		for i := 0; i < qLen; i++ {
			current := q[0]
			q = q[1:]

			row, col := current[0], current[1]

			for _, dir := range dirs {
				r, c := dir[0]+row, dir[1]+col

				if r >= height || r < 0 || c >= width || c < 0 || grid[r][c] != 1 {
					continue
				}

				grid[r][c] = 2
				q = append(q, []int{r, c})
				good--
			}
		}
	}

	if good == 0 {
		return result
	}

	return -1
}
