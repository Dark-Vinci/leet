package backtracking

import (
	"fmt"
	"slices"
)

func smallestNumber(s string) string {
	var dfs func(curr int, prev int, str string, used map[int]struct{}) []string

	dfs = func(curr int, prev int, str string, used map[int]struct{}) []string {
		if curr == len(s)+1 {
			return []string{str}
		}

		result := make([]string, 0)

		for i := 1; i <= 9; i++ {
			if _, ok := used[i]; ok {
				continue
			}

			if prev == 0 || (s[curr-1] == 'D' && i < prev) || (s[curr-1] == 'I' && i > prev) {
				used[i] = struct{}{}

				strr := str + fmt.Sprintf("%v", i)

				result = append(result, dfs(curr+1, i, strr, used)...)

				delete(used, i)
			}
		}

		return result
	}

	result := dfs(0, 0, "", map[int]struct{}{})

	slices.Sort(result)

	return result[0]
}

// ALSO CORRECT
func smallestNumber1(s string) string {
	var dfs func(curr int, prev int, str string, used map[int]struct{}) []string

	dfs = func(curr int, prev int, str string, used map[int]struct{}) []string {
		if curr == len(s)+1 {
			return []string{str}
		}

		result := make([]string, 0)

		if prev == 0 {
			for i := 1; i <= 9; i++ {
				used[i] = struct{}{}

				strr := str + fmt.Sprintf("%v", i)

				result = append(result, dfs(curr+1, i, strr, used)...)

				delete(used, i)
			}
		} else {
			f := s[curr-1]

			for i := 1; i <= 9; i++ {
				if _, ok := used[i]; ok {
					continue
				}

				if (f == 'D' && i < prev) || (f == 'I' && i > prev) {
					used[i] = struct{}{}

					strr := str + fmt.Sprintf("%v", i)

					result = append(result, dfs(curr+1, i, strr, used)...)

					delete(used, i)
				}
			}
		}

		return result
	}

	result := dfs(0, 0, "", map[int]struct{}{})

	slices.Sort(result)

	return result[0]
}
