package Grid

func countBattleships(board [][]byte) int {
	type node struct {
		x, y int
	}

	var (
		dfs             func(isHorizontal bool, i, j int)
		l1, l2          = len(board), len(board[0])
		result, visited = 0, make(map[node]struct{})
	)

	dfs = func(isHorizontal bool, i, j int) {
		n := node{x: i, y: j}

		if _, ok := visited[n]; ok {
			return
		}

		if isHorizontal {
			if i >= l2 || i < 0 {
				return
			}

			if board[j][i] == '.' {
				return
			}

			visited[n] = struct{}{}

			dfs(isHorizontal, i+1, j)
			dfs(isHorizontal, i-1, j)
		} else {
			if j >= l1 || j < 0 {
				return
			}

			if board[j][i] == '.' {
				return
			}

			visited[n] = struct{}{}

			dfs(isHorizontal, i, j+1)
			dfs(isHorizontal, i, j-1)
		}
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			n := node{x: i, y: j}
			if _, ok := visited[n]; !ok && board[j][i] == 'X' {
				isHor := true

				if i+1 < l2 {
					if board[j][i+1] == 'X' {
						isHor = true
					} else {
						isHor = false
					}
				} else if j+1 < l1 {
					if board[j+1][i] == 'X' {
						isHor = false
					} else {
						isHor = true
					}
				}

				dfs(isHor, i, j)
				result++
			}
		}
	}

	return result
}
