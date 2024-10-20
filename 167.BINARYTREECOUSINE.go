package main

func isCousins(root *TreeNode, x int, y int) bool {
	type node struct {
		par  *TreeNode
		curr *TreeNode
	}

	q := []*node{{par: nil, curr: root}}

	for len(q) > 0 {
		l := len(q)
		var first *node

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c.curr.Left != nil {
				q = append(q, &node{par: c.curr, curr: c.curr.Left})
			}

			if c.curr.Right != nil {
				q = append(q, &node{par: c.curr, curr: c.curr.Right})
			}

			if c.curr.Val == x || c.curr.Val == y {
				if first == nil {
					first = c
				} else if first.par != c.par {
					return true
				}
			}
		}
	}

	return false
}
