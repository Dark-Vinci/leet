package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func splitIntoFibonacci(num string) []int {
	var (
		dfs    func(a, b string, i int, db []int) bool
		result = make([]int, 0)
	)

	dfs = func(a, b string, i int, db []int) bool {
		if i == len(num) {
			result = slices.Clone(db)
			return true
		}

		aInt, _ := strconv.Atoi(a)
		bInt, _ := strconv.Atoi(b)

		if aInt+bInt > math.MaxInt32 {
			return false
		}

		strSum := fmt.Sprintf("%v", aInt+bInt)

		if !strings.HasPrefix(num[i:], strSum) {
			return false
		}

		db = append(db, aInt+bInt)

		return dfs(b, strSum, i+len(strSum), db)
	}

	for i := 1; i <= len(num)-2; i++ {
		a := num[0:i]
		aInt, _ := strconv.Atoi(a)

		if a[0] == '0' && len(a) > 1 {
			continue
		}

		for j := i + 1; j < len(num); j++ {
			b := num[i:j]

			if b[0] == '0' && len(b) > 1 {
				continue
			}

			bInt, _ := strconv.Atoi(b)

			if dfs(a, b, j, []int{aInt, bInt}) {
				return result
			}
		}
	}

	return result
}
