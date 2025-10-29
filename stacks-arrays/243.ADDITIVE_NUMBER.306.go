package arrays

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	var dfs func(ind int, a, b string) bool

	dfs = func(ind int, a, b string) bool {
		if ind == len(num) {
			return true
		}

		aVal, _ := strconv.Atoi(a)
		bVal, _ := strconv.Atoi(b)

		sum := aVal + bVal

		sumStr := fmt.Sprintf("%v", sum)

		if !strings.HasPrefix(num[ind:], sumStr) {
			return false
		}

		return dfs(ind+len(sumStr), b, sumStr)
	}

	for i := 1; i < len(num); i++ {
		a := num[0:i]

		if a[0] == '0' && len(a) > 1 {
			continue
		}

		for j := i + 1; j < len(num); j++ {
			b := num[i:j]

			if b[0] == '0' && len(b) > 1 {
				continue
			}

			if dfs(j, a, b) {
				return true
			}
		}
	}

	return false
}

/// ONLY ONE TEST CASE DID NOT PASS
// func isAdditiveNumber(num string) bool {
//     if len(num) < 3 {
//         return false
//     }

//     var (
//         dfs func(i int, prev string, prevVal int) bool
//         l = len(num)
//         first, _ = strconv.Atoi(string(num[0]))
//     )

//     dfs = func(i int, prev string, prevVal int) bool {
//         if i >= len(num) {
//             return true
//         }

//         for j := i + 1; j <= l; j++ {
//             curr := string(num[i: j])
//             currVal, _ := strconv.Atoi(curr)

//             if len(curr) > 0 && curr[0] == '0' && currVal != 0 {
//                 continue
//             }

//             for k := j + 1; k <= l; k++ {
//                 next := string(num[j : k])
//                 nextVal, _ := strconv.Atoi(next)

//                 if len(next) > 0 && next[0] == '0' && nextVal != 0 {
//                     continue
//                 }

//                 if nextVal > currVal + prevVal {
//                     break
//                 }

//                 if prevVal + currVal == nextVal {
//                     fmt.Println(prev, curr, next)
//                     if dfs(j, curr, currVal) {
//                         return true
//                     }

//                     if k == l {
//                         return true
//                     }
//                 }
//             }
//         }

//         return false
//     }

//     return dfs(1, string(num[0]), first)
// }
