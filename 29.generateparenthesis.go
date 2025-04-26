package main

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var dfs func(current string, l, r int)

	dfs = func(current string, l, r int) {
		if (l == 0 && r == 0) || r < l {
			return
		}

		if l == 0 {
			for i := 0; i < r; i++ {
				current += ")"
			}

			result = append(result, current)
			return
		}

		lCurrent := current + "("
		rCurrent := current + ")"

		dfs(lCurrent, l-1, r)
		dfs(rCurrent, l, r-1)
	}

	dfs("", n, n)
	return result
}
