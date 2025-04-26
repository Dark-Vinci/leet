package dynamicProgramming

func numWays(words []string, target string) int {
	var (
		dfs  func(i int, n int) int
		l1   = len(words[0])
		memo = make(map[[2]int]int)
		frq  = make([][26]int, l1)
	)

	for _, word := range words {
		for i, w := range word {
			frq[i][w-'a']++
		}
	}

	dfs = func(i int, n int) int {
		if i == len(target) {
			return 1
		}

		if n == l1 {
			return 0
		}

		key := [2]int{i, n}

		if v, ok := memo[key]; ok {
			return v
		}

		res := dfs(i, n+1)

		chr := target[i] - 'a'
		count := frq[n][chr]

		if count > 0 {
			res += count * dfs(i+1, n+1) % 1000000007
			res %= 1000000007
		}

		memo[key] = res

		return memo[key]
	}

	return dfs(0, 0)
}

// EXCEED TIME LIMIT
func numWays_(words []string, target string) int {
	var (
		dfs  func(i int, n int, s string) int
		l    = len(words)
		l1   = len(words[0])
		memo = make(map[[2]int]int)
	)

	dfs = func(i int, n int, s string) int {
		key := [2]int{i, n}

		if v, ok := memo[key]; ok {
			return v
		}

		if i >= len(target) || n >= l1 {
			if s == target {
				return 1
			}

			return 0
		}

		res := 0

		for j := 0; j < l; j++ {
			word := words[j]

			for k := n; k < l1; k++ {
				ss := s + string(word[k])

				if word[k] == target[i] {
					res += dfs(i+1, k+1, ss)
				}
			}
		}

		memo[key] = res % 1000000007

		return memo[key]
	}

	return dfs(0, 0, "")
}
