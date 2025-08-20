package Trees

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	l := preorderTraversal(root.Left)
	r := preorderTraversal(root.Right)

	res := make([]int, 0)

	res = append(res, root.Val)
	res = append(res, l...)
	res = append(res, r...)

	return res
}
