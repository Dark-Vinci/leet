package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var dfs func(r1, r2 *TreeNode) *TreeNode

	dfs = func(r1, r2 *TreeNode) *TreeNode {
		if r1 == nil && r2 == nil {
			return nil
		}

		var (
			current = &TreeNode{}
			left    *TreeNode
			right   *TreeNode
		)

		if r1 == nil {
			current.Val = r2.Val
			left, right = dfs(nil, r2.Left), dfs(nil, r2.Right)
		} else if r2 == nil {
			current.Val = r1.Val
			left, right = dfs(r1.Left, nil), dfs(r1.Right, nil)
		} else {
			current.Val = r1.Val + r2.Val
			left, right = dfs(r1.Left, r2.Left), dfs(r1.Right, r2.Right)
		}

		current.Left = left
		current.Right = right

		return current
	}

	return dfs(root1, root2)
}
