package arrays

func findAnagrams(s string, p string) []int {
	var (
		result = make([]int, 0)
		lp     = len(p)
		ls     = len(s)
		pa     = [26]int{}
	)

	for i := 0; i < len(p); i++ {
		pa[int(p[i]-'a')]++
	}

	for i := 0; i < ls; i++ {
		cp := pa

		if i+lp <= ls {
			for j := i; j < lp+i; j++ {
				cp[int(s[j]-'a')]--
			}
		}

		if cp == [26]int{} {
			result = append(result, i)
		}
	}

	return result
}
