package main

import (
	"fmt"
	"math"
	"strconv"
)

func grayCodeMORE_ELEGANT(n int) (result []int) {

	for i := 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	return result
}

func grayCode(n int) []int {
	bins := make([]string, 0)
	pow := math.Pow(2.0, float64(n))
	l := len(intToBin(int(pow - 1)))

	for i := 0; i <= int(pow)-1; i++ {
		bin := intToBin(i)
		diff := l - len(bin)

		for j := 0; j < diff; j++ {
			bin = "0" + bin
		}

		bins = append(bins, bin)
	}

	for i := 0; i < len(bins); i++ {
		newVal := ""

		for j := 0; j < len(bins[i]); j++ {
			var gn byte

			if j == 0 {
				gn = bins[i][j] - '0'
			} else {
				gn = (bins[i][j] - '0') ^ (bins[i][j-1] - '0')
			}

			newVal += fmt.Sprintf("%v", gn)
		}

		bins[i] = newVal
	}

	result := make([]int, len(bins))

	for i := 0; i < len(bins); i++ {
		val, _ := strconv.ParseInt(bins[i], 2, 64)
		result[i] = int(val)
	}

	return result
}

func intToBin(n int) string {
	result := ""

	for n != 0 {
		result = fmt.Sprintf("%v", n%2) + result
		n = n / 2
	}

	return result
}
