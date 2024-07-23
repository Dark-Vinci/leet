package main

func masDepth(root *Node) int {
	var dfs func(r *Node, depth int)
	var maxDepth int

	dfs = func(r *Node, depth int) {
		if r == nil {
			return
		}

		depth++

		if maxDepth < depth {
			maxDepth = depth
		}

		for _, child := range r.Children {
			dfs(child, depth)
		}
	}

	dfs(root, 0)

	return maxDepth
}
