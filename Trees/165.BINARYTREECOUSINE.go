package Trees

func replaceValueInTree(root *TreeNode) *TreeNode {
	type node struct {
		par     *TreeNode
		curr    *TreeNode
		prevVal int
	}

	q, ind := []*node{{par: nil, curr: root, prevVal: root.Val}}, 0

	for {
		l := len(q)

		if l == ind {
			break
		}

		for i := ind; i < l; i++ {
			c := q[i]

			if c.curr.Left != nil {
				q = append(q, &node{par: c.curr, curr: c.curr.Left, prevVal: c.curr.Left.Val})
			}

			if c.curr.Right != nil {
				q = append(q, &node{par: c.curr, curr: c.curr.Right, prevVal: c.curr.Right.Val})
			}

			val := 0

			for j := ind; j < l; j++ {
				if c != q[j] && c.par != q[j].par {
					val += q[j].prevVal
				}
			}

			c.curr.Val = val
		}

		ind = l
	}

	return root
}
