package main

import "slices"

func findLongestWord(s string, dictionary []string) string {
	result := ""

	slices.Sort(dictionary)

	for i := 0; i < len(dictionary); i++ {
		curr := dictionary[i]
		k, count := 0, 0

		for j := 0; j < len(curr); j++ {
			for ; k < len(s); k++ {
				if curr[j] == s[k] {
					count++
					k++
					break
				}
			}
		}

		if count == len(curr) && len(result) < len(curr) {
			result = curr
		}
	}

	return result
}
