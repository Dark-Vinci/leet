package main

import "math"

func maxCardScore(cardPoints []int, k int) int {
	mx := math.MinInt

	for i := k; i >= 0; i-- {
		rIndex := len(cardPoints) - k + i
		l, r := sum(cardPoints[0:i]), sum(cardPoints[rIndex:])

		mx = max(mx, l+r)
	}

	return mx
}
