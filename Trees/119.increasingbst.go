package Trees

func increasingBST(root *TreeNode) *TreeNode {
	var dfs func(r, t *TreeNode) *TreeNode

	dfs = func(r, t *TreeNode) *TreeNode {
		if r == nil {
			return t
		}

		res := dfs(r.Left, r)
		r.Left = nil
		r.Right = dfs(r.Right, t)

		return res
	}

	return dfs(root, nil)
}
