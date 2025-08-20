package Trees

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}

	odd := false

	for len(q) > 0 {
		l := len(q)

		if odd {
			for i, j := 0, l-1; i <= j; i, j = i+1, j-1 {
				q[i].Val, q[j].Val = q[j].Val, q[i].Val
			}
		}

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c.Left != nil {
				q = append(q, c.Left)
			}

			if c.Right != nil {
				q = append(q, c.Right)
			}
		}

		odd = !odd
	}

	return root
}
