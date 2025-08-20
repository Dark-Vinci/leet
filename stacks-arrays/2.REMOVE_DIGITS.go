package stacks_arrays

import (
	"cmp"
	"slices"
)

func removeDigit(number string, digit byte) string {
	options := make([]string, 0)

	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			nums := slices.Delete([]byte(number), i, i+1)

			options = append(options, string(nums))
		}
	}

	slices.SortFunc(options, func(a, b string) int {
		return cmp.Compare(b, a)
	})

	return options[0]
}
