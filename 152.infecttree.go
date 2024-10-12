package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
	infected bool
	val      *TreeNode
	par      *node
	left     *node
	right    *node
}

func amountOfTime(root *TreeNode, start int) int {
	q := make([]*node, 0)

	var dfs func(parent *node, current *TreeNode) *node
	dfs = func(parent *node, current *TreeNode) *node {
		if current == nil {
			return nil
		}

		n := node{
			val: current,
			par: parent,
		}

		if current.Val == start {
			n.infected = true
			q = append(q, &n)
		}

		n.left, n.right = dfs(&n, current.Left), dfs(&n, current.Right)

		return &n
	}

	_ = dfs(nil, root)
	count := -1

	for len(q) > 0 {
		nextQ := make([]*node, 0)

		for i := 0; i < len(q); i++ {
			c := q[i]

			if c.par != nil && c.par.infected == false {
				c.par.infected = true
				nextQ = append(nextQ, c.par)
			}

			if c.left != nil && c.left.infected == false {
				c.left.infected = true
				nextQ = append(nextQ, c.left)
			}

			if c.right != nil && c.right.infected == false {
				c.right.infected = true
				nextQ = append(nextQ, c.right)
			}
		}

		q = nextQ
		count++
	}

	return count
}

func amountOfTime0(root *TreeNode, start int) int {
	if root.Left == nil && root.Right == nil {
		if root.Val == start {
			return 0
		} else {
			return 1
		}
	}

	type node struct {
		infected bool
		val      *TreeNode
		par      *node
		left     *node
		right    *node
	}

	nodes := make([]*node, 0)

	var dfs func(parent *node, current *TreeNode) *node
	dfs = func(parent *node, current *TreeNode) *node {
		if current == nil {
			return nil
		}

		n := node{
			val: current,
			par: parent,
		}

		l, r := dfs(&n, current.Left), dfs(&n, current.Right)

		n.left, n.right = l, r

		if current.Val == start {
			n.infected = true
		}

		nodes = append(nodes, &n)

		return &n
	}

	_ = dfs(nil, root)

	count := -1
	dealth := make(map[int]struct{})

	for len(dealth) < len(nodes) {
		infInd := make(map[int]struct{})

		for i := 0; i < len(nodes); i++ {
			if _, ok := dealth[i]; nodes[i].infected == true && !ok {
				infInd[i] = struct{}{}
			}
		}

		for k := range infInd {
			curr := nodes[k]

			if curr.par != nil {
				curr.par.infected = true
			}

			if curr.left != nil {
				curr.left.infected = true
			}

			if curr.right != nil {
				curr.right.infected = true
			}

			dealth[k] = struct{}{}
		}
		count++
	}

	return count
}
