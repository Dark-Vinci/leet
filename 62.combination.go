package main

import (
	"fmt"
	"slices"
	"strings"
)

func combination(s string) int {
	solution := make([][]string, 0)

	var dfs func(index int, prev []string)

	dfs = func(index int, prev []string) {
		if index == len(s) {
			solution = append(solution, prev)
			return
		}

		dfs(index+1, prev)

		prev = append(prev, string(s[index]))

		dfs(index+1, prev)
	}

	dfs(0, []string{})

	fmt.Println(solution)

	return len(solution)
}

//func main() {
//	fmt.Println(combination("1010"))
//}

func formula(s string) int {
	num := factorial(len(s))

	fac := 0

	for i := 1; i <= len(s); i++ {
		den := factorial(i) * factorial(len(s)-i)

		fac += num / den
	}

	return fac + 1
}

func factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func newPermutation(s string) []string {
	var result = make([]string, 0)

	var recurse func(res string)
	recurse = func(res string) {
		if len(res) == len(s) {
			result = append(result, res)
			return
		}

		for i := 0; i < len(s); i++ {
			if !strings.Contains(res, string(s[i])) {
				newRes := res + string(s[i])
				recurse(newRes)
			}
		}
	}

	recurse("")

	return result
}

func permutation(s string) []string {
	result := make([]string, 0)

	var dfs func(index int, prev string)

	dfs = func(index int, prev string) {
		if index == len(s) {
			result = append(result, prev)
			return
		}

		for i := 0; i <= len(prev); i++ {
			slicedPrev := strings.Split(prev, "")
			slicedPrev = slices.Insert(slicedPrev, i, string(s[index]))
			newPrev := strings.Join(slicedPrev, "")

			dfs(index+1, newPrev)
		}
	}

	dfs(0, "")

	fmt.Println(result)

	return result
}
