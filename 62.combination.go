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
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
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

//func main() {
//	//a := combination("122")
//	//b := formula("CCABBA")
//	_ = permutation("ABC")
//
//	//fmt.Println(c)
//}
