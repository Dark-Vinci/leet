package main

import (
	"slices"
	"strings"
)

func permutationCombination(s string) int {
	var (
		result = make([]string, 0)
		dfs    func(prev []string, index int)
	)

	dfs = func(prev []string, index int) {
		if index == len(s) {
			str := strings.Join(prev, "")

			if !slices.Contains(result, str) && str != "" {
				result = append(result, str)
			}
			return
		}

		dfs(prev, index+1)

		for i := 0; i <= len(prev); i++ {
			cp := slices.Clone(prev)
			cp = slices.Insert(cp, i, string(s[index]))

			dfs(cp, index+1)
		}
	}

	dfs([]string{}, 0)

	return len(result)
}

//func main() {
//	a := permutationCombination("ACCCCBC")
//	fmt.Println(a)
//}
