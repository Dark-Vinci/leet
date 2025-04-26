package dynamicProgramming

import "fmt"

func checkValidString(s string) bool {
	var (
		dfs  func(i int, opCount int) bool
		memo = make(map[string]bool)
	)

	dfs = func(i int, opCount int) bool {
		key := fmt.Sprintf("%v-%v", i, opCount)

		if v, ok := memo[key]; ok {
			return v
		}

		if i >= len(s) {
			memo[key] = opCount == 0
			return memo[key]
		}

		if s[i] == '*' {
			if dfs(i+1, opCount+1) {
				memo[key] = true
				return true
			}

			if dfs(i+1, opCount) {
				memo[key] = true
				return true
			}

			if opCount > 0 {
				if dfs(i+1, opCount) {
					memo[key] = true
					return true
				}
			}
		}

		if s[i] == '(' {
			if dfs(i+1, opCount+1) {
				memo[key] = true
				return true
			}
		} else if opCount > 0 {
			if dfs(i+1, opCount-1) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false

		return memo[key]
	}

	return dfs(0, 0)
}

// EXCEED TIME LIMMIT
func checkValidString_(s string) bool {
	var dfs func(i int, ss string) bool
	cst := [3]string{"(", ")", ""}
	validMemo := make(map[string]bool)
	memo := make(map[string]bool)

	dfs = func(i int, ss string) bool {
		key := fmt.Sprintf("%v-%v", i, ss)

		fmt.Println(key)

		if v, ok := memo[key]; ok {
			return v
		}

		if i >= len(s) {
			if v, ok := validMemo[ss]; ok {
				return v
			}

			validMemo[ss] = isValid(ss)
			return validMemo[ss]
		}

		if s[i] == '*' {
			for _, v := range cst {
				if dfs(i+1, ss+v) {
					memo[key] = true
					return true
				}
			}
		}

		memo[key] = dfs(i+1, ss+string(s[i]))
		return memo[key]
	}

	return dfs(0, "")
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	count := 0

	for _, v := range s {
		if v == '(' {
			count++
			continue
		}

		if count <= 0 {
			return false
		}

		count--
	}

	return count == 0
}
