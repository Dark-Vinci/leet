package main

import "slices"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(arr []int) *TreeNode

	dfs = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}

		m := slices.Max(arr)
		index := slices.Index(arr, m)

		l, r := dfs(arr[:index]), dfs(arr[index+1:])

		return &TreeNode{
			Left:  l,
			Right: r,
			Val:   m,
		}
	}

	return dfs(nums)
}
