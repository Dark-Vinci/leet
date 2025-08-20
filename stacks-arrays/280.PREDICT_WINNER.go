package stacks_arrays

import (
	"cmp"
	"slices"
)

func predictTheWinner(nums []int) bool {
	var helper func(arr []int, val int, first bool) bool

	helper = func(arr []int, val int, first bool) bool {
		if len(arr) == 0 {
			return val >= 0
		}

		l := len(arr)

		a, b := arr[0], arr[l-1]
		aArr, bArr := slices.Clone(arr)[1:], slices.Clone(arr)[:l-1]

		if first {
			a1, b1 := helper(aArr, val+a, false), helper(bArr, val+b, false)

			return cmp.Or(a1, b1)
		}

		a1, b1 := helper(aArr, val-a, true), helper(bArr, val-b, true)

		return a1 && b1
	}

	return helper(nums, 0, true)
}
