package main

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var (
		result = make([][]string, 0)
		m      = make(map[string][]string)
	)

	for i := 0; i < len(strs); i++ {
		sSlices := strings.Split(strs[i], "")
		slices.Sort(sSlices)

		s := strings.Join(sSlices, "")

		if _, ok := m[s]; ok {
			m[s] = append(m[s], strs[i])
		} else {
			m[s] = []string{strs[i]}
		}
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
