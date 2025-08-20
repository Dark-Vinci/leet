package Trees

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	nodeValues := make([]int, 0)

	for len(q) > 0 {
		l := len(q)

		nodeValues = append(nodeValues, q[l-1].Val)

		for i := 0; i < l; i++ {
			currentNode := q[0]

			q = q[1:]

			if currentNode.Left != nil {
				q = append(q, currentNode.Left)
			}

			if currentNode.Right != nil {
				q = append(q, currentNode.Right)
			}
		}
	}

	return nodeValues
}
