package Trees

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var (
		dfs    func(r *TreeNode) bool
		result = make([]int, 0)
		index  int
	)

	dfs = func(r *TreeNode) bool {
		if r == nil {
			return true
		}

		if r.Val != voyage[index] {
			return false
		}

		index++

		if r.Left != nil && r.Left.Val != voyage[index] {
			result = append(result, r.Val)

			return dfs(r.Right) && dfs(r.Left)
		}

		return dfs(r.Left) && dfs(r.Right)
	}

	if !dfs(root) {
		return []int{-1}
	}

	return result
}
