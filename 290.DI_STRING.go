package main

func diStringMatch(s string) []int {
	var (
		result = make([]int, len(s)+1)
		l, r   = 0, len(s)
	)

	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			result[i] = r
			r--
		} else {
			result[i] = l
			l++
		}
	}

	result[len(s)] = l

	return result
}
