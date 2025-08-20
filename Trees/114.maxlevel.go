package Trees

import (
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		q      = []*TreeNode{root}
		result = make([]int, 0)
	)

	for len(q) > 0 {
		l := len(q)
		maxValue := math.MinInt

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c.Val > maxValue {
				maxValue = c.Val
			}

			if c.Left != nil {
				q = append(q, c.Left)
			}

			if c.Right != nil {
				q = append(q, c.Right)
			}
		}

		result = append(result, maxValue)
	}

	return result
}
