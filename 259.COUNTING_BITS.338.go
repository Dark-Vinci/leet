package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countBits(n int) []int {
	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		a := fmt.Sprintf("%v", strconv.FormatInt(int64(i), 2))

		r := strings.Count(a, "1")

		arr[i] = r
	}

	return arr
}

// Learnt: Hamming weight or pop count;
// uses; (res[i>>1] + i&1) -> Count Number of ones
// zeros[i >> 1] + (!(i & 1)) -> Count Number of zeros
