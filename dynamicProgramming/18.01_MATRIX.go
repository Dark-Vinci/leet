package dynamicProgramming

import "math"

func updateMatrix(mat [][]int) [][]int {
	var (
		l1, l2 = len(mat), len(mat[0])
		q      = [][2]int{}
		result = make([][]int, l1)
		dirs   = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	)

	for i := 0; i < l1; i++ {
		result[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			if mat[i][j] == 1 {
				result[i][j] = math.MaxInt
			} else {
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		l := len(q)

		for i := 0; i < l; i++ {
			curr := q[0]
			q = q[1:]

			x, y := curr[0], curr[1]
			d := result[x][y]

			for _, dir := range dirs {
				xx, yy := x+dir[0], y+dir[1]

				if xx >= 0 && yy >= 0 && xx < l1 && yy < l2 && result[xx][yy] > d+1 {
					q = append(q, [2]int{xx, yy})
					result[xx][yy] = d + 1
				}
			}
		}
	}

	return result
}
