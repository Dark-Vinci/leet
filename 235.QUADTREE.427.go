package main

type NodeQ struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *NodeQ
	TopRight    *NodeQ
	BottomLeft  *NodeQ
	BottomRight *NodeQ
}

// true is 1
func construct(grid [][]int) *NodeQ {
	var dfs func(is, ie, js, je int) *NodeQ

	dfs = func(is, ie, js, je int) *NodeQ {
		mp := make(map[int]struct{})

		for j := js; j <= je; j++ {
			for i := is; i <= ie; i++ {
				mp[grid[j][i]] = struct{}{}
			}
		}

		if len(mp) == 1 {
			_, ok1 := mp[1]

			n := &NodeQ{
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}

			if ok1 {
				n.Val = true
			}

			return n
		}

		im, jm := is+((ie-is)/2), js+((je-js)/2)

		tl, tr := dfs(is, im, js, jm), dfs(im+1, ie, js, jm)
		bl, br := dfs(is, im, jm+1, je), dfs(im+1, ie, jm+1, je)

		return &NodeQ{
			IsLeaf:      false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}

	return dfs(0, len(grid[0])-1, 0, len(grid[0])-1)
}
