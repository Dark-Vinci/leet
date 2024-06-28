package main

func levelOrder2(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }

    q := []*Node{root}

    result := make([][]int, 0)

    for len(q) > 0 {
        level := make([]int, 0)
        s := len(q)

        for i := 0; i < s; i++ {
            c := q[0]
            q = q[1:]

            level = append(level, c.Val)

            for _, v := range c.Children {
                q = append(q, v)
            }
        }

        result = append(result, level)
    }

    return result
}
