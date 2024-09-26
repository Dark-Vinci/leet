package main

import (
	"math"
	"slices"
)

func maximumTastiness(price []int, k int) int {
	slices.Sort(price)

	left, right := 0, math.MaxInt

	for left <= right {
		mid := (left + right) / 2

		if isTastinessValid(price, k, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func isTastinessValid(price []int, k int, x int) bool {
	last, count, i := price[0], 1, 1

	for i < len(price) {
		if price[i]-last >= x {
			last = price[i]
			count++
		}
		i++

		if count >= k {
			return true
		}
	}

	return count >= k
}
