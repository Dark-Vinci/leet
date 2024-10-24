package main

func pacificAtlantic(heights [][]int) [][]int {
	var (
		result = make([][]int, 0)
		l1, l2 = len(heights), len(heights[0])
		atl    = make(map[[2]int]struct{})
		pac    = make(map[[2]int]struct{})
		dfs    func(visited map[[2]int]struct{}, i, j, prev int)
	)

	dfs = func(visited map[[2]int]struct{}, i, j, prev int) {
		n := [...]int{i, j}

		if _, ok := visited[n]; ok || i == l2 || j == l1 || i < 0 || j < 0 || heights[j][i] < prev {
			return
		}

		visited[n] = struct{}{}
		curr := heights[j][i]

		dfs(visited, i+1, j, curr) // RIGHT
		dfs(visited, i-1, j, curr) // LEFT
		dfs(visited, i, j-1, curr) // UP
		dfs(visited, i, j+1, curr) // DOWN
	}

	for i := 0; i < l2; i++ {
		// TOP
		dfs(pac, i, 0, heights[0][i])
		// BOTTOM
		dfs(atl, i, l1-1, heights[l1-1][i])
	}

	for i := 0; i < l1; i++ {
		// LEFT
		dfs(pac, 0, i, heights[i][0])
		// RIGHT
		dfs(atl, l2-1, i, heights[i][l2-1])
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			n := [...]int{i, j}
			_, ok1 := atl[n]
			_, ok2 := pac[n]

			if ok1 && ok2 {
				result = append(result, []int{j, i})
			}
		}
	}

	return result
}
