package arrays

import (
	"cmp"
	"new_apps/rust_build/melonnu/Trees"
)

func pacq(heights [][]int) [][]int {
	result := make([][]int, 0)
	l1, l2 := len(heights), len(heights[0])
	cache := make(map[[2]int]bool)

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		if i >= l1 || j >= l2 {
			return true
		}

		if v, ok := cache[[2]int{i, j}]; ok {
			return v
		}

		var left, right, up, down bool

		// RIGHT
		if i+1 == l2 {
			right = true
		} else if heights[j][i] >= heights[j][i+1] {
			right = dfs(i+1, j)
		}

		// DOWN
		if j+1 == l1 {
			down = true
		} else if heights[j][i] >= heights[j+1][i] {
			right = dfs(i, j+1)
		}

		if !right && !down {
			cache[[...]int{i, j}] = false
			return false
		}

		// UP
		if j-1 < 0 {
			up = true
		} else if heights[j][i] >= heights[j-1][i] {
			up = dfs(i, j-1)
		}

		// LEFT
		if i-1 < 0 {
			left = true
		} else if heights[j][i] >= heights[j][i-1] {
			left = dfs(i-1, j)
		}

		res := cmp.Or(right && up, down && left)

		cache[[...]int{i, j}] = res

		return res
	}

	for j := 0; j < l1; j++ {
		for i := 0; i < l2; i++ {
			if dfs(i, j) {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func lievelOrder(root *Trees.TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := []*Trees.TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var nums []int
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
