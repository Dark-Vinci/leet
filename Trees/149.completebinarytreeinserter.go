package Trees

type CBTInserter struct {
	root *TreeNode
}

func constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root: root}
}

func (this *CBTInserter) Insert(val int) int {
	var (
		q    = []*TreeNode{this.root}
		last = 0
		node = &TreeNode{Val: val}
	)

	for len(q) != last {
		l := len(q)

		for i := last; i < l; i++ {
			c := q[i]

			if c.Left != nil {
				q = append(q, c.Left)
			}

			if c.Right != nil {
				q = append(q, c.Right)
			}
		}

		last = l
	}

	index := len(q) + 1

	i := (index / 2) - 1

	if index%2 == 0 {
		q[i].Left = node
	} else {
		q[i].Right = node
	}

	return q[i].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
