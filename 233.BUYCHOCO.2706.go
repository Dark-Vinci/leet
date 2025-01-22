package main

import "slices"

func buyChoco(prices []int, money int) int {
	slices.Sort(prices)

	a, b := prices[0], prices[1]

	if a+b > money {
		return money
	}

	return money - a - b
}
