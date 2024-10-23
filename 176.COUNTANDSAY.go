package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	mOne := countAndSay(n - 1)

	i, j := 0, 0

	var result strings.Builder

	for j < len(mOne) && i < len(mOne) {
		if j+1 < len(mOne) && mOne[j+1] == mOne[j] {
			j++
			continue
		}

		numb, _ := strconv.Atoi(string(mOne[i]))

		result.WriteString(fmt.Sprintf("%v%v", j-i+1, numb))
		i = j + 1
		j++
	}

	return result.String()
}
