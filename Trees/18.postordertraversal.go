package Trees

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	l := postorderTraversal(root.Left)
	r := postorderTraversal(root.Right)

	result := make([]int, 0)

	result = append(result, l...)
	result = append(result, r...)
	result = append(result, root.Val)

	return result
}
