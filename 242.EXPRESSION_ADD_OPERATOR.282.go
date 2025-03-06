package main

import (
	"fmt"
	"strconv"
)

func addOperators(num string, target int) []string {
	var (
		result = make([]string, 0)
		dfs    func(num, prefix string, val, last int)
	)

	dfs = func(num, prefix string, val, last int) {
		if len(num) == 0 {
			if val == target {
				result = append(result, prefix)
			}

			return
		}

		for i := 1; i <= len(num); i++ {
			ith := num[:i]
			ithVal, _ := strconv.Atoi(ith)

			if num[0] == '0' && i != 1 {
				continue
			}

			if len(prefix) == 0 {
				dfs(num[i:], ith, ithVal, ithVal) // first dfs recursive call
				continue
			}

			dfs(num[i:], fmt.Sprintf("%v%v%v", prefix, "+", ith), val+ithVal, ithVal)
			dfs(num[i:], fmt.Sprintf("%v%v%v", prefix, "-", ith), val-ithVal, -ithVal)
			dfs(num[i:], fmt.Sprintf("%v%v%v", prefix, "*", ith), val-last+last*ithVal, last*ithVal) // multiplication is a wierd beast
		}
	}

	dfs(num, "", 0, 0)

	return result
}
