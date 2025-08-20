package combinatorics

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// right answer
func getPermutation(n int, k int) string {
	var (
		pool   = make([]int, n)
		result strings.Builder
	)

	for i := 1; i <= n; i++ {
		pool[i-1] = i
	}

	for i := n; i > 0; i-- {
		countPerBlock := factorial(i - 1)
		containedBlock := int(math.Ceil(float64(k) / float64(countPerBlock)))

		result.WriteString(fmt.Sprintf("%d", pool[containedBlock-1]))
		pool = slices.Delete(pool, containedBlock-1, containedBlock)

		k -= countPerBlock * (containedBlock - 1)
	}

	return result.String()
}

// EXCEEDS TIME LIMIT(I DON TAYA)[STILL WORTH IT THO]
func getTimeLimitPermutation(n int, k int) string {
	p := make([]string, 0)

	var permute func(prePerm string, i int)
	permute = func(prePerm string, i int) {
		if i > n {
			return
		}

		i++

		for j := 0; j <= len(prePerm); j++ {
			nSlice := strings.Split(prePerm, "")
			nSlice = slices.Insert(nSlice, j, fmt.Sprintf("%d", i-1))

			postPerm := strings.Join(nSlice, "")

			if i-1 == n {
				p = append(p, postPerm)
			} else {
				permute(postPerm, i)
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
