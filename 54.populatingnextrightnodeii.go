package main

func connect1(root *HNode) *HNode {
	if root == nil {
        return root
    }

    q := []*HNode{root}

    for len(q) > 0 {
        length := len(q)

        for i := 0; i < length; i++ {
            currentNode := q[0]
            q = q[1:]

            if i != length - 1 {
                currentNode.Next = q[0]
            } else {
                currentNode.Next = nil
            }

            if currentNode.Left != nil {
                q = append(q, currentNode.Left)
            }

            if currentNode.Right != nil {
                q = append(q, currentNode.Right)
            }
        }
    }

    return root
}