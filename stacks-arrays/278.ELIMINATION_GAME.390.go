package stacks_arrays

import "slices"

func lastRemaining(n int) int {
	if n == 1 {
		return n
	}

	return 2 * (1 + (n / 2) - lastRemaining(n/2))
}

// exceed time limit
func lastRemaining_(n int) int {
	var (
		helper func(arr []int, forw bool)
		arr    = make([]int, 0)
		result = -1
	)

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	helper = func(arr []int, forw bool) {
		if len(arr) == 1 {
			result = arr[0]
			return
		}

		nn := make([]int, 0)

		ll := len(arr) - 1

		for i := 0; i < len(arr); i++ {
			if forw {
				if i%2 != 0 {
					nn = slices.Insert(nn, len(nn), arr[i])
				}
			} else {
				j := len(arr) - i - 1

				if ll%2 == 0 {
					if j%2 != 0 {
						nn = slices.Insert(nn, 0, arr[j])
					}
				} else {
					if j%2 == 0 {
						nn = slices.Insert(nn, 0, arr[j])
					}
				}
			}
		}

		helper(nn, !forw)
	}

	helper(arr, true)

	return result
}
