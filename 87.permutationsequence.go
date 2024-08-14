package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// EXCEEDS TIME LIMIT(I DON TAYA)
func getPermutation(n int, k int) string {
	p := make([]string, 0)

	var permute func(db string, i int)
	permute = func(db string, i int) {
		if i > n {
			return
		}

		i++

		for j := 0; j <= len(db); j++ {
			nSlice := strings.Split(db, "")
			nSlice = slices.Insert(nSlice, j, fmt.Sprintf("%d", i-1))

			str := strings.Join(nSlice, "")

			if i-1 == n {
				p = append(p, str)
			} else {
				permute(str, i)
			}
		}
	}

	permute("", 1)

	slices.SortFunc(p, func(a, b string) int {
		numA, _ := strconv.Atoi(a)
		numB, _ := strconv.Atoi(b)

		return cmp.Compare(numA, numB)
	})

	return p[k-1]
}
