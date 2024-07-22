package main

func maxxDepth(root *Node) int {
    var dfs func(r *Node, depth int)
    var max int

    dfs = func(r *Node, depth int) {
        if r == nil {
            return
        }

        depth++

        if max < depth {
            max = depth
        }

        for _, child := range r.Children {
            dfs(child, depth)
        }
    }

    dfs(root, 0)

    return max
}
