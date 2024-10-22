package main

func islandPerimeter(grid [][]int) int {
	l1, l2 := len(grid), len(grid[0])
	result := 0

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			if grid[j][i] == 1 {
				val := 4

				if j+1 < l1 && grid[j+1][i] == 1 { //down
					val--
				}

				if j-1 >= 0 && grid[j-1][i] == 1 { //up
					val--
				}

				if i+1 < l2 && grid[j][i+1] == 1 { //right
					val--
				}

				if i-1 >= 0 && grid[j][i-1] == 1 { //left
					val--
				}

				result += val
			}
		}
	}

	return result
}
