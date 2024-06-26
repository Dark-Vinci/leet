package main

func lievelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		nums := []int{}
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			nums = append(nums, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, nums)
	}

	return result
}