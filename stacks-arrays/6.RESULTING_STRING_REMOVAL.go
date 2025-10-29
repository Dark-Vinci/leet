package arrays

import "slices"

func resultingString(s string) string {
	str := []byte{s[0]}

	for i := 1; i < len(s); i++ {
		if len(str) == 0 {
			str = append(str, s[i])
			continue
		}

		curr, last := s[i], str[len(str)-1]

		if isValid(last, curr) {
			str = str[:len(str)-1]
			continue
		}

		str = append(str, curr)
	}

	return string(str)
}

// TLE
func resultingString1(s string) string {
	byStr := []byte(s)

	for {
		last := len(byStr)

		for i := 1; i < len(byStr); i++ {
			if isValid(byStr[i-1], byStr[i]) {
				byStr = slices.Delete(byStr, i-1, i+1)
				break
			}
		}

		if len(byStr) == last {
			break
		}
	}

	return string(byStr)
}

func isValid(a, b byte) bool {
	if a == 'a' && b == 'z' || a == 'z' && b == 'a' {
		return true
	}

	return (a-'a')-(b-'a') == 1 || (b-'a')-(a-'a') == 1
}
