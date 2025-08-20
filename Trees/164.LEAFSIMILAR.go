package Trees

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(r *TreeNode, p *[]int)
	r1, r2 := make([]int, 0), make([]int, 0)

	dfs = func(r *TreeNode, p *[]int) {
		if r == nil {
			return
		}

		if r.Left == nil && r.Right == nil {
			*p = append(*p, r.Val)
			return
		}

		dfs(r.Left, p)
		dfs(r.Right, p)
	}

	dfs(root1, &r1)
	dfs(root2, &r2)

	if len(r1) != len(r2) {
		return false
	}

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
