package main

func averageOfLevels(root *TreeNode) []float64 {
	sum := make([]int, 0)
	count := make([]int, 0)

	db := []*TreeNode{root}

	for len(db) > 0 {
		s := len(db)
		su := 0
		co := 0

		for i := 0; i < s; i++ {
			c := db[0]
			db = db[1:]

			su += c.Val
			co += 1

			if c.Left != nil {
				db = append(db, c.Left)
			}

			if c.Right != nil {
				db = append(db, c.Right)
			}
		}

		sum = append(sum, su)
		count = append(count, co)
	}

	res := make([]float64, 0)

	for i := 0; i < len(sum); i++ {
		avg := float64(sum[i]) / float64(count[i])

		res = append(res, avg)
	}

	return res
}
