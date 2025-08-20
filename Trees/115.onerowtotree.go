package Trees

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	var dfs func(r *TreeNode, height int)

	dfs = func(r *TreeNode, height int) {
		if r == nil {
			return
		}

		if height == depth-1 {
			l, ri := r.Left, r.Right

			r.Left = &TreeNode{
				Val:  val,
				Left: l,
			}

			r.Right = &TreeNode{
				Val:   val,
				Right: ri,
			}

			return
		}

		height++

		dfs(r.Left, height)
		dfs(r.Right, height)
	}

	dfs(root, 1)

	return root
}
