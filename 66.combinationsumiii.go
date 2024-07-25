package main

import "slices"

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	var dfs func(values []int)

	dfs = func(values []int) {
		if len(values) > k {
			return
		}

		if len(values) == k && sum(values) == n {
			flag := false

			for _, v := range result {
				if slices.Equal(v, values) {
					flag = true
				}
			}

			if !flag {
				result = append(result, values)
			}

			return
		}

		for i := 1; i <= 9; i++ {
			if !slices.Contains(values, i) {
				newValues := slices.Clone(values)
				newValues = append(newValues, i)

				slices.Sort(newValues)

				dfs(newValues)
			}
		}
	}

	dfs([]int{})

	return result
}
