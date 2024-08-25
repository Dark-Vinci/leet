package main

func findTarget(root *TreeNode, k int) bool {
	var (
		db  = make([]int, 0)
		dfs func(r *TreeNode)
	)

	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		db = append(db, r.Val)

		dfs(r.Left)
		dfs(r.Right)
	}

	dfs(root)

	for i := 0; i < len(db); i++ {
		for j := 0; j < len(db); j++ {
			if i != j && db[i]+db[j] == k {
				return true
			}
		}
	}

	return false
}
