package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isEq func(a *TreeNode, v int) bool

	isEq = func(a *TreeNode, v int) bool {
		if a == nil {
			return true
		}

		if a.Val != v {
			return false
		}

		return a.Val == v && isEq(a.Left, v) && isEq(a.Right, v)
	}

	return isEq(root, root.Val)
}
