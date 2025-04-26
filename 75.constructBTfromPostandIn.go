package main

import "slices"

func buildTree(inorder []int, postorder []int) *TreeNode {
	var dfs func(in []int, postIndex *int) *TreeNode

	dfs = func(arr []int, pIndex *int) *TreeNode {
		if *pIndex < 0 || len(arr) == 0 {
			return nil
		}

		node := &TreeNode{
			Val: postorder[*pIndex],
		}

		index := slices.Index(arr, postorder[*pIndex]) + 1

		left := arr[:index-1]
		right := arr[index:]

		*pIndex--

		node.Right, node.Left = dfs(right, pIndex), dfs(left, pIndex)

		return node
	}

	length := len(postorder) - 1

	return dfs(inorder, &length)
}
