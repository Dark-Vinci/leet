package Trees

func convertBST(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode, int) int

	dfs = func(r *TreeNode, top int) int {
		if r == nil {
			return 0
		}

		right := dfs(r.Right, top)

		curr := r.Val

		r.Val += right + top

		left := dfs(r.Left, r.Val)

		return curr + left + right
	}

	dfs(root, 0)

	return root
}
