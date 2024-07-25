package main

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs func(prev []int)

	dfs = func(prev []int) {
		if sum(prev) > target {
			return
		}

		if sum(prev) == target {
			flag := false

			for _, v := range result {
				if slices.Equal(v, prev) {
					flag = true
				}
			}

			if !flag {
				result = append(result, prev)
			}

			return
		}

		for _, v := range candidates {
			newP := slices.Clone(prev)
			newP = append(newP, v)

			slices.Sort(newP)

			dfs(newP)
		}
	}

	dfs([]int{})

	return result
}

func sum(a []int) int {
	result := 0

	for _, v := range a {
		result += v
	}

	return result
}
