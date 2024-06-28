package main

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }

    children := make([]int, 0)

    for _, v := range root.Children {
        children = append(children, postorder(v)...)
    }

    children = append(children, root.Val)

    return children
}
