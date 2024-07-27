package main

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	result := make([][]int, 0)

	var dfs func(prev []int, cands []int)

	dfs = func(prev []int, cands []int) {
		if sum(prev) > target {
			return
		}

		if sum(prev) == target {
			result = append(result, prev)
			return
		}

		for i := 0; i < len(cands); i++ {
			if i > 0 && cands[i] == cands[i-1] {
				continue
			}

			current := slices.Clone(prev)
			current = append(current, cands[i])

			dfs(current, cands[i+1:])
		}
	}

	dfs([]int{}, candidates)

	return result
}
