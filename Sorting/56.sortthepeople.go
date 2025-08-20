package Sorting

import (
	"cmp"
	"slices"
)

func sortPeople(names []string, heights []int) []string {
	heightMap := make(map[int]int)
	namesMap := make(map[int]string)

	for k, v := range names {
		namesMap[k] = v
	}

	for k, v := range heights {
		heightMap[k] = v
	}

	slices.SortFunc(heights, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	for index, height := range heights {
		for k, v := range heightMap {
			if height == v {
				a, _ := namesMap[k]
				names[index] = a
				break
			}
		}
	}

	return names
}
