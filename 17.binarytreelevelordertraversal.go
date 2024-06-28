package main

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    q := []*TreeNode{root}
    result := make([][]int, 0)

    for len(q) > 0 {
        s := len(q)
        b := make([]int, s)
        for i := 0; i < s; i++ {
            c := q[0]
            q = q[1:]

            b[i] = c.Val

            if c.Left != nil {
                q = append(q, c.Left)
            }
            if c.Right != nil {
                q = append(q, c.Right)
            }
        }
        result = append([][]int{b}, result...)
    }

    return result
}
