package Trees

import (
	"cmp"
	"math"
	"slices"
)

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var dfs func(r, min, max *TreeNode)

	dfs = func(r, min, max *TreeNode) {
		if r == nil {
			return
		}

		if r.Val > max.Val {
			r.Val, max.Val = max.Val, r.Val
		}

		if r.Val < min.Val {
			r.Val, min.Val = min.Val, r.Val
		}

		dfs(r.Left, min, r)
		dfs(r.Right, r, max)
	}

	for !isValidBST(root) {
		dfs(root, &TreeNode{Val: math.MinInt}, &TreeNode{Val: math.MaxInt})
	}
}

func recoverTree1(root *TreeNode) {
	nodes := make([]*TreeNode, 0)
	var dfs func(r *TreeNode)

	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r.Left)

		nodes = append(nodes, r)

		dfs(r.Right)
	}

	dfs(root)

	nodeValues := make([]int, len(nodes))

	for k, v := range nodes {
		nodeValues[k] = v.Val
	}

	slices.SortFunc(nodeValues, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	for i, j := range nodeValues {
		nodes[i].Val = j
	}
}
