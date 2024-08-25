package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(r, s *TreeNode) bool

	dfs = func(r, s *TreeNode) bool {
		if r == nil && s == nil {
			return true
		}

		if r == nil || s == nil {
			return false
		}

		if r.Val != s.Val {
			return false
		}

		if !dfs(r.Left, s.Left) {
			return false
		}

		return dfs(r.Right, s.Right)
	}

	if dfs(root, subRoot) {
		return true
	}

	if root != nil && isSubtree(root.Left, subRoot) {
		return true
	}

	return root != nil && isSubtree(root.Right, subRoot)
}
