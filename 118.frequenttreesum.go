package main

import "math"

func findFrequentTreeSum(root *TreeNode) []int {
	db := make(map[int]int)

	var dfs func(*TreeNode) int

	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		left, right := dfs(r.Left), dfs(r.Right)

		val := right + left + r.Val

		db[val] += 1

		return val
	}

	dfs(root)

	maxInt := math.MinInt

	for _, v := range db {
		if v > maxInt {
			maxInt = v
		}
	}

	result := make([]int, 0)

	for k, v := range db {
		if v == maxInt {
			result = append(result, k)
		}
	}

	return result
}
