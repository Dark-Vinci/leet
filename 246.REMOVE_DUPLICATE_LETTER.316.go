package main

func removeDuplicateLetters(s string) string {
	var (
		result    string
		seen      = [26]bool{}
		stk       = make([]int, 0)
		lastIndex = [26]int{}
	)

	for i := 0; i < len(s); i++ {
		lastIndex[int(s[i]-'a')] = i
	}

	for i := 0; i < len(s); i++ {
		char := int(s[i] - 'a')

		if seen[char] {
			continue
		}

		for len(stk) > 0 && stk[len(stk)-1] > char && i < lastIndex[stk[len(stk)-1]] {
			val := stk[len(stk)-1]

			stk = stk[:len(stk)-1]

			seen[val] = false
		}

		seen[char] = true
		stk = append(stk, char)
	}

	for j := len(stk) - 1; j >= 0; j-- {
		result = string([]byte{byte(stk[j]) + 'a'}) + result
	}

	return result
}
