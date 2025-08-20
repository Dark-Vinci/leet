package backtracking

import "slices"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(i int, db []int)

	dfs = func(i int, db []int) {
		if i >= len(nums) {
			result = append(result, db)
			return
		}

		cAdd := append(slices.Clone(db), nums[i])
		cNot := slices.Clone(db)

		dfs(i+1, cAdd)
		dfs(i+1, cNot)
	}

	dfs(0, []int{})

	return result
}
