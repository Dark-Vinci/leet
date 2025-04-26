package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}

	sum := 0

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if c.Left != nil {
			q = append(q, c.Left)

			if c.Left.Left == nil && c.Left.Right == nil {
				sum += c.Left.Val
			}
		}

		if c.Right != nil {
			q = append(q, c.Right)
		}
	}

	return sum
}
