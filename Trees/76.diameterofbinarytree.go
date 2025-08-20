package Trees

func diameterOfBinaryTree(root *TreeNode) int {
	var result int

	var dfs func(r *TreeNode) int

	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		left, right := dfs(r.Left), dfs(r.Right)

		if result < left+right {
			result = left + right
		}

		return 1 + max(left, right)
	}

	_ = dfs(root)

	return result
}
