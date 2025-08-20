package Trees

type node1 struct {
	curr  *TreeNode
	count int
	left  *node1
	right *node1
	par   *node1
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var (
		dfs     func(par *node1, curr *TreeNode) *node1
		q       = make([]*node1, 0)
		visited = make(map[*node1]struct{})
		result  = make([]int, 0)
	)

	dfs = func(par *node1, curr *TreeNode) *node1 {
		if curr == nil {
			return nil
		}

		n := &node1{
			curr: curr,
			par:  par,
		}

		n.left, n.right = dfs(n, curr.Left), dfs(n, curr.Right)

		if curr == target {
			q = append(q, n)
			visited[n] = struct{}{}
		}

		return n
	}

	_ = dfs(nil, root)

	for len(q) > 0 {
		for i := 0; i < len(q); i++ {
			c := q[0]
			q = q[1:]

			if c.count == k {
				result = append(result, c.curr.Val)
				continue
			}

			if _, ok := visited[c.left]; !ok && c.left != nil {
				c.left.count = c.count + 1
				q = append(q, c.left)
				visited[c.left] = struct{}{}
			}

			if _, ok := visited[c.right]; !ok && c.right != nil {
				c.right.count = c.count + 1
				q = append(q, c.right)
				visited[c.right] = struct{}{}
			}

			if _, ok := visited[c.par]; !ok && c.par != nil {
				c.par.count = c.count + 1
				q = append(q, c.par)
				visited[c.par] = struct{}{}
			}
		}
	}

	return result
}
