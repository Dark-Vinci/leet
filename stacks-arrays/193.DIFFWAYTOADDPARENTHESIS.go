package stacks_arrays

import "strconv"

func diffWaysToCompute(expression string) []int {
	var (
		dfs  func(exp string) []int
		memo = make(map[string][]int)
	)

	dfs = func(exp string) []int {
		if v, ok := memo[exp]; ok {
			return v
		}
		if len(exp) <= 2 {
			v, _ := strconv.Atoi(exp)
			return []int{v}
		}

		ans := make([]int, 0)

		for i, v := range exp {
			if v == '+' || v == '-' || v == '*' {
				left, right := dfs(exp[:i]), dfs(exp[i+1:])
				for _, l := range left {
					for _, r := range right {
						if v == '+' {
							ans = append(ans, l+r)
						} else if v == '-' {
							ans = append(ans, l-r)
						} else {
							ans = append(ans, l*r)
						}
					}
				}
			}
		}

		memo[exp] = ans
		return ans
	}

	return dfs(expression)
}
