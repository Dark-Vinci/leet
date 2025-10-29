package arrays

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	multiple := make([]int, 0)
	str := make([]string, 0)
	curr, num := "", ""

	for i := 0; i < len(s); i++ {
		e := string(s[i])

		if unicode.IsDigit(rune(s[i])) {
			num += e
			continue
		}

		if e == "[" {
			m, _ := strconv.Atoi(num)
			multiple, str = append(multiple, m), append(str, curr)

			curr, num = "", ""
		} else if e == "]" {
			curr = str[len(str)-1] + (strings.Repeat(curr, multiple[len(multiple)-1]))

			str, multiple = str[:len(str)-1], multiple[:len(multiple)-1]
		} else {
			curr += e
		}
	}

	return curr
}
