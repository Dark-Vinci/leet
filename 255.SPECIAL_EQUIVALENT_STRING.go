package main

import (
	"cmp"
	"slices"
)

func numSpecialEquivGroups(words []string) int {
	db := make(map[string]struct{})

	for _, val := range words {
		var odd, even []rune

		for i, v := range val {
			if i%2 == 0 {
				even = append(even, v)
			} else {
				odd = append(odd, v)
			}
		}

		slices.SortFunc(odd, func(a, b rune) int {
			return cmp.Compare(a, b)
		})

		slices.SortFunc(even, func(a, b rune) int {
			return cmp.Compare(a, b)
		})

		db[string(even)+string(odd)] = struct{}{}
	}

	return len(db)
}
