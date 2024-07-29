package main

func pathSum3(root *TreeNode, targetSum int) int {
	count := 0

	var dfs func(r *TreeNode, prev []int)

	dfs = func(r *TreeNode, prev []int) {
		if r == nil {
			return
		}

		bbb := 0

		prev = append(prev, r.Val)

		for i := len(prev) - 1; i >= 0; i-- {
			bbb += prev[i]

			if bbb == targetSum {
				count++
			}
		}

		dfs(r.Left, prev)
		dfs(r.Right, prev)
	}

	dfs(root, []int{})

	return count
}
