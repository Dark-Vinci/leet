package main

func isCompleteTree(root *TreeNode) bool {
    if root == nil {
        return true
    }

    q := []*TreeNode{root}

    isNull := false

    for len(q) > 0 {
        c := q[0]

        if isNull && c != nil {
            return false
        }

        q = q[1:]

        if c == nil {
            isNull = true
        } else {
            q = append(q, c.Left)
            q = append(q, c.Right)
        }
    }

    return true
}
