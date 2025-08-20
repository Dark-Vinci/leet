package Trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth func(r *TreeNode, max int) int

	depth = func(r *TreeNode, max int) int {
		if r == nil {
			return max
		}

		max += 1

		a := depth(r.Left, max)
		b := depth(r.Right, max)

		if a > b {
			return a
		}

		return b
	}

	return depth(root, 0)
}
