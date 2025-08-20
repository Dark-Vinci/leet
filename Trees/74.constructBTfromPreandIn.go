package Trees

import (
	"slices"
)

func buildTree1(preorder []int, inorder []int) *TreeNode {
	var dfs func(arr []int, pIndex *int) *TreeNode

	dfs = func(arr []int, pIndex *int) *TreeNode {
		if len(arr) == 0 || *pIndex >= len(preorder) {
			return nil
		}

		node := &TreeNode{
			Val: preorder[*pIndex],
		}

		index := slices.Index(arr, preorder[*pIndex]) + 1

		left := arr[:index-1]
		right := arr[index:]

		*pIndex++

		node.Left, node.Right = dfs(left, pIndex), dfs(right, pIndex)

		return node
	}

	index := 0

	return dfs(inorder, &index)
}
