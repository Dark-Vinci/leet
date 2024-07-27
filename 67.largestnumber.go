package main

import (
	"slices"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsStr := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		numsStr[i] = strconv.Itoa(nums[i])
	}

	slices.SortFunc(numsStr, func(a, b string) int {
		if a+b > b+a {
			return -1
		}

		return 1
	})

	if numsStr[0] == "0" {
		return "0"
	}

	return strings.Join(numsStr, "")
}
