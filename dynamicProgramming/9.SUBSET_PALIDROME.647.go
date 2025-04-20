package dynamicProgramming

func countSubstrings(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if isPalidrome(s[i : j+1]) {
				result++
			}
		}
	}

	return result
}

func isPalidrome(s string) bool {
	for i := 0; i < (len(s)/2)+1; i++ {
		j := len(s) - 1 - i

		if s[i] != s[j] {
			return false
		}
	}

	return true
}
