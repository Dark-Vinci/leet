package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	h := height(root)

	result := make([][]int, 0)

	for i := 0; i < h; i++ {
		result = append(result, level(root, i))
	}

	return result
}

func level(a *TreeNode, l int) []int {
	if a == nil {
		return []int{}
	}

	res := make([]int, 0)

	if l == 0 {
		res = append(res, a.Val)
	} else {
		le := level(a.Left, l-1)
		ri := level(a.Right, l-1)

		res = append(res, le...)
		res = append(res, ri...)
	}

	return res
}

func height(a *TreeNode) int {
	if a == nil {
		return 0
	}

	l := height(a.Left)
	r := height(a.Right)

	if l > r {
		return l + 1
	}

	return r + 1
}
