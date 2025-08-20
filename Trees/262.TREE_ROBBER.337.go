package Trees

import (
	"fmt"
	"slices"
)

func rob(root *TreeNode) int {
	var dfs func(r *TreeNode) [2]int

	dfs = func(r *TreeNode) [2]int {
		if r == nil {
			return [2]int{0, 0}
		}

		left, right := dfs(r.Left), dfs(r.Right)

		res := [2]int{}

		res[0] = max(left[0], left[1]) + max(right[0], right[1])
		res[1] = r.Val + left[0] + right[0]

		return res
	}

	res := dfs(root)

	return max(res[0], res[1])
}

// even better solution i found somewhere
func rob__(root *TreeNode) int {
	res := robtraversal(root)
	return slices.Max(res)
}

func robtraversal(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}

	left := robtraversal(cur.Left)
	right := robtraversal(cur.Right)

	rob := cur.Val + left[0] + right[0]
	notRob := slices.Max(left) + slices.Max(right)

	return []int{notRob, rob}
}

// time limit exceeded
func rob_(root *TreeNode) int {
	var (
		dfs  func(r *TreeNode) int
		memo = make(map[*TreeNode]int)
	)

	dfs = func(r *TreeNode) int {
		if root == nil {
			return 0
		}

		if v, ok := memo[r]; ok {
			return v
		}

		result := 0

		if root.Left != nil {
			result += rob(root.Left.Left) + rob(root.Left.Right)
		}

		if root.Right != nil {
			result += rob(root.Right.Right) + rob(root.Right.Left)
		}

		memo[r] = max(result+root.Val, rob(root.Left)+rob(root.Right))

		fmt.Println(memo)

		return memo[r]
	}

	return dfs(root)
}
