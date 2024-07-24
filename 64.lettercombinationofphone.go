package main

func letterCombinations(digits string) []string {
	keys := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return nil
	}

	if len(digits) == 1 {
		k, _ := keys[digits]
		return k
	}

	var merge func(a []string, b []string) []string

	merge = func(a []string, b []string) []string {
		newA := make([]string, 0)

		for j := 0; j < len(a); j++ {
			aa := a[j]

			for i := 0; i < len(b); i++ {
				bb := b[i]

				concat := aa + bb

				newA = append(newA, concat)
			}
		}

		return newA
	}

	result, _ := keys[string(digits[0])]

	for i := 1; i < len(digits); i++ {
		curr, _ := keys[string(digits[i])]

		result = merge(result, curr)
	}

	return result
}
