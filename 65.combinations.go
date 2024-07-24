package main

import "slices"

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var dfs func(prev []int, max int)

	dfs = func(prev []int, max int) {
		if len(prev) == k {
			result = append(result, prev)
			return
		}

		for i := 1; i <= n; i++ {
			if !slices.Contains(prev, i) && i >= max {
				newP := slices.Clone(prev)
				newP = append(newP, i)

				dfs(newP, i)
			}
		}
	}

	dfs([]int{}, 0)

	return result
}
