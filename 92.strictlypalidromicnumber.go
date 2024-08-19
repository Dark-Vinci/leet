package main

import "fmt"

func isStrictlyPalindromic(n int) bool {
	curr := 2

	for curr <= n-2 {
		numb := convertToBase(n, curr)

		fmt.Println(numb)

		if !isPalindrome(numb) {
			return false
		}

		curr++
	}

	return true
}

func convertToBase(n int, base int) string {
	result := ""

	for n > 0 {
		remainder := n % base

		result = fmt.Sprintf("%v", remainder) + result

		n = n / base
	}

	return result
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
