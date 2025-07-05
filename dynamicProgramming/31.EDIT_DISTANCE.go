package dynamicProgramming

//https://leetcode.com/problems/edit-distance/
//medium

func minDistance(word1 string, word2 string) int {
	var (
		dp     func(i int, j int, str1, str2 string, memo map[[2]int]int) int
		l1, l2 = len(word1), len(word2)
	)

	dp = func(i int, j int, str1, str2 string, memo map[[2]int]int) int {
		if j < 0 {
			return i + 1
		}

		if i < 0 {
			return j + 1
		}

		key := [2]int{i, j}

		if v, ok := memo[key]; ok {
			return v
		}

		if str1[i] == str2[j] {
			memo[key] = dp(i-1, j-1, str1, str2, memo)
			return memo[key]
		}

		insert := dp(i, j-1, str1, str2, memo)
		delet := dp(i-1, j, str1, str2, memo)
		replace := dp(i-1, j-1, str1, str2, memo)

		result := min(insert, delet, replace) + 1

		memo[key] = result

		return memo[key]
	}

	return dp(l1-1, l2-1, word1, word2, map[[2]int]int{})
}
