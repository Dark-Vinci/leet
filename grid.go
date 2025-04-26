package main

import (
	"strings"
)

func showConstruct(s string, strs []string) [][]string {
	table := make([][][]string, len(s)+1)

	for i := 0; i < len(strs); i++ {
		table[i] = make([][]string, 0)
	}

	table[0] = [][]string{{}}

	for i := 0; i <= len(s); i++ {
		sliceStr := table[i]

		if len(sliceStr) == 0 {
			continue
		}

		for _, val := range strs {
			for _, v := range sliceStr {
				inner := strings.Join(v, "")

				ss := inner + val

				if strings.HasPrefix(s, ss) && i+len(val) <= len(s) {
					v = append(v, val)
					table[i+len(val)] = append(table[i+len(val)], v)
				}
			}
		}
	}

	return table[len(s)]
}

func canSum(num int, nums []int) [][]int {
	result := make([][][]int, num+1)

	for i := range result {
		result[i] = make([][]int, 0)
	}

	result[0] = [][]int{{}}

	for i := 0; i <= num; i++ {
		if len(result[i]) == 0 {
			continue
		}

		for _, ops := range nums {
			if i+ops > num {
				continue
			}

			for _, val := range result[i] {
				val = append(val, ops)

				result[i+ops] = append(result[i+ops], val)
			}
		}
	}

	return result[num]
}

func gridTraveller(m, n int) int {
	grid := make([][]int, m+1)

	for i := range grid {
		row := make([]int, n+1)

		grid[i] = row
	}

	grid[1][1] = 1

	for j := 0; j <= m; j++ {
		for i := 0; i <= n; i++ {
			current := grid[j][i]

			if j+1 <= m {
				grid[j+1][i] += current
			}

			if i+1 <= n {
				grid[j][i+1] += current
			}
		}
	}

	return grid[m][n]
}

// func main() {
// 	res := showConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"})

// 	fmt.Println(res)
// }
