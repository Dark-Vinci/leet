package Trees

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	var path func(a *TreeNode, b []int)

	path = func(a *TreeNode, b []int) {
		if a == nil {
			return
		}

		b = append(b, a.Val)

		if a.Left == nil && a.Right == nil {
			sum := 0

			for _, v := range b {
				sum += v
			}

			if sum == targetSum {
				result = append(result, b)
			}

			return
		}

		cp := make([]int, len(b))

		copy(cp, b)

		path(a.Left, cp)
		path(a.Right, cp)

		return
	}

	path(root, []int{})

	return result
}
