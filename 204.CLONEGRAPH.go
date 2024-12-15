package main

type GNode struct {
	Val       int
	Neighbors []*GNode
}

func cloneGraph(node *GNode) *GNode {
	return cloneHelper(node, map[int]*GNode{})
}

func cloneHelper(node *GNode, db map[int]*GNode) *GNode {
	if node == nil {
		return node
	}

	n := &GNode{
		Val:       node.Val,
		Neighbors: make([]*GNode, len(node.Neighbors)),
	}

	db[node.Val] = n

	for i, nei := range node.Neighbors {
		if val, ok := db[nei.Val]; ok {
			n.Neighbors[i] = val
		} else {
			n.Neighbors[i] = cloneHelper(nei, db)
		}
	}

	return n
}
