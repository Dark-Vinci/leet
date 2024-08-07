package main

func allPossibleFBT(n int) []*TreeNode {
	cache := make([][]*TreeNode, n+1)

	var dfs func(n int) []*TreeNode

	dfs = func(n int) []*TreeNode {
		if len(cache[n]) > 0 {
			return cache[n]
		}

		if n == 1 {
			return []*TreeNode{&TreeNode{Val: 0}}
		}

		result := make([]*TreeNode, 0)

		for i := 1; i < n-1; i++ {
			j := n - i - 1

			for _, left := range dfs(i) {
				for _, right := range dfs(j) {
					ans := &TreeNode{
						Val:   0,
						Left:  left,
						Right: right,
					}

					result = append(result, ans)
				}
			}
		}

		cache[n] = result

		return result
	}

	return dfs(n)
}
