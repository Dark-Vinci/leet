package Trees

func generateTrees(n int) []*TreeNode {
	var dfs func(l, r int) []*TreeNode

	dfs = func(l, r int) []*TreeNode {
		if l == r {
			el := &TreeNode{Val: l}

			return []*TreeNode{el}
		}

		if l > r {
			return []*TreeNode{nil}
		}

		result := make([]*TreeNode, 0)

		for i := l; i <= r; i++ {
			for _, left := range dfs(l, i-1) {
				for _, right := range dfs(i+1, r) {
					root := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}

					result = append(result, root)
				}
			}
		}

		return result
	}

	return dfs(1, n)
}
