package arrays

import (
	"cmp"
	"fmt"
)

func baseNeg2(n int) string {
	var result string

	for n != 0 {
		r := n % -2
		n /= -2

		if r < 0 {
			r += 2
			n += 1
		}

		result = fmt.Sprintf("%v", r) + result
	}

	return cmp.Or(result, "0")
}
