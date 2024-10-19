package main

import (
	"math"
	"strings"
)

func bestClosingTime(customers string) int {
	l, fCount, tCount := len(customers), 0, 0
	pref, post := make([]int, l+1), make([]int, l+1)

	for i := 0; i <= l; i++ {
		j := l - i

		if j < l && customers[j] == 'Y' {
			tCount++
		}

		pref[i] = fCount
		post[j] = tCount

		if i == l {
			break
		}

		if customers[i] == 'N' {
			fCount++
		}
	}

	mn, val := -1, math.MaxInt

	for i := 0; i <= l; i++ {
		pref[i] += post[i]

		if pref[i] < val {
			val = pref[i]
			mn = i
		}
	}

	return mn
}

func bestClosingTime0(customers string) int {
	if customers == strings.Repeat("Y", len(customers)) {
		return len(customers)
	}

	if customers == strings.Repeat("N", len(customers)) {
		return 0
	}

	result, mn := 0, math.MaxInt

	for i := 0; i <= len(customers); i++ {
		count := 0

		for j := 0; j < len(customers); j++ {
			if customers[j] == 'Y' && i-1 < j {
				count++
			}

			if customers[j] == 'N' && i-1 >= j {
				count++
			}
		}

		if count < mn {
			mn = count
			result = i
		}
	}

	return result
}
