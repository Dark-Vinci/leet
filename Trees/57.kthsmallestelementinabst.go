package Trees

import (
	"slices"
)

func kthSmallest(root *TreeNode, k int) int {
	db := make([]int, 0)

	var dfs func(r *TreeNode)

	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		db = append(db, r.Val)

		dfs(r.Left)
		dfs(r.Right)
	}

	dfs(root)

	slices.Sort(db)

	return db[k-1]
}
