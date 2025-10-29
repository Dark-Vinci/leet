package arrays

import (
	"slices"
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	left, count := 0, [26]int{}

	for _, v := range s1 {
		count[int(v-'a')]++
	}

	for right, val := range s2 {
		count[int(val-'a')]--

		if count == [26]int{} {
			return true
		}

		if right+1 >= len(s1) {
			count[int(s2[left]-'a')]++
			left++
		}
	}

	return false
}

// exceed time complexity
func checkInclusionWORSTTIME(s1 string, s2 string) bool {
	var (
		perms = make([]string, 0)
		dfs   func([]string, int)
	)

	if len(s1) > len(s2) {
		return false
	}

	dfs = func(above []string, index int) {
		if index == len(s1) {
			perms = append(perms, strings.Join(above, ""))
			return
		}

		for i := 0; i <= len(above); i++ {
			cp := slices.Clone(above)

			cp = slices.Insert(cp, i, string(s1[index]))

			dfs(cp, index+1)
		}
	}

	dfs([]string{}, 0)

	for _, v := range perms {
		if strings.Contains(s2, v) {
			return true
		}
	}

	return false
}
