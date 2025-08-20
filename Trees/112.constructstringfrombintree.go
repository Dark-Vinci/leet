package Trees

import (
	"fmt"
)

func tree2str(root *TreeNode) string {
	var dfs func(r *TreeNode) string

	dfs = func(r *TreeNode) string {
		if r == nil {
			return ""
		}

		res := fmt.Sprintf("%d", r.Val)

		left, right := dfs(r.Left), dfs(r.Right)

		if left != "" {
			res += fmt.Sprintf("(%v)", left)
		} else if right != "" {
			res += "()"
		}

		if right != "" {
			res += fmt.Sprintf("(%v)", right)
		}

		return res
	}

	return dfs(root)
}
