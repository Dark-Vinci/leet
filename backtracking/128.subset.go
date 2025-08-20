package backtracking

import (
	"fmt"
	"slices"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	var (
		result = make([][]int, 0)
		mp     = make(map[string]struct{})
		dfs    func(i int, db []int)
	)

	dfs = func(i int, db []int) {
		if i >= len(nums) {
			var str strings.Builder

			slices.Sort(db)

			for j := 0; j < len(db); j++ {
				str.WriteString(fmt.Sprintf("%v", db[j]))
			}

			if _, ok := mp[str.String()]; !ok {
				mp[str.String()] = struct{}{}
				result = append(result, db)
			}

			return
		}

		cAdd := append(slices.Clone(db), nums[i])
		cNot := slices.Clone(db)

		dfs(i+1, cAdd)
		dfs(i+1, cNot)
	}

	dfs(0, []int{})

	return result
}
