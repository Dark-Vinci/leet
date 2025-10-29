package arrays

func longestValidParentheses(s string) int {
	var (
		stk = []int{-1}
		mx  = 0
	)

	for i, v := range s {
		if v == '(' {
			stk = append(stk, i)
			continue
		}

		if len(stk) > 0 {
			stk = stk[:len(stk)-1]

			if len(stk) == 0 {
				stk = append(stk, i)
			} else {
				mx = max(mx, i-stk[len(stk)-1])
			}
		}
	}

	return mx
}
