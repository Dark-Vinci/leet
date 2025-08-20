package Trees

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}

	var depth int

	for len(q) > 0 {
		length := len(q)

		changed := false

		for i := 0; i < length; i++ {
			current := q[0]

			if !changed {
				depth = q[0].Val
				changed = true
			}

			q = q[1:]

			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
		}
	}

	return depth
}
