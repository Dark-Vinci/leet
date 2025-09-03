package greedy

import "slices"

func maximumGain(input string, x int, y int) int {
	stk := make([]byte, 0)
	result := 0

	f, fs := "ab", x
	s, ss := "ba", y

	if ss > fs {
		f, fs, s, ss = s, ss, f, fs
	}

	for i := 0; i < len(input); i++ {
		if len(stk) > 0 && string(stk[len(stk)-1])+string(input[i]) == f {
			stk = stk[:len(stk)-1]
			result += fs
		} else {
			stk = append(stk, input[i])
		}
	}

	clone := slices.Clone(stk)
	stk = []byte{}

	for i := 0; i < len(clone); i++ {
		if len(stk) > 0 && string(stk[len(stk)-1])+string(clone[i]) == s {
			stk = stk[:len(stk)-1]
			result += ss
		} else {
			stk = append(stk, clone[i])
		}
	}

	return result
}
