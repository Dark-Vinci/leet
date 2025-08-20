package stacks_arrays

import (
	"math"
	"new_apps/rust_build/melonnu/combinatorics"
)

func maxCardScore(cardPoints []int, k int) int {
	mx := math.MinInt

	for i := k; i >= 0; i-- {
		rIndex := len(cardPoints) - k + i
		l, r := combinatorics.sum(cardPoints[0:i]), combinatorics.sum(cardPoints[rIndex:])

		mx = max(mx, l+r)
	}

	return mx
}
