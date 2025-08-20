package dynamicProgramming

type tract struct {
	val   string
	count int
}

func countTexts(pressedKeys string) int {
	var (
		keys = map[string]int{
			"2": 3,
			"3": 3,
			"4": 3,
			"5": 3,
			"6": 3,
			"7": 4,
			"8": 3,
			"9": 4,
		}
		split  = make([]tract, 0)
		t      = tract{val: string(pressedKeys[0]), count: 1}
		result = 1
		dfs    func(l, n, mx int) int
		memo   = make(map[int]int)
	)

	for i := 1; i < len(pressedKeys); i++ {
		current := string(pressedKeys[i])

		if i == len(pressedKeys)-1 {
			if t.val == current {
				t.count++
			} else {
				split = append(split, tract{val: current, count: 1}) // last
			}

			split = append(split, t)
			break
		}

		if current == t.val {
			t.count++
		} else {
			split = append(split, t)

			t = tract{val: current, count: 1}
		}
	}

	dfs = func(l, n, mx int) int {
		if val, ok := memo[n]; ok {
			return val
		}

		if n == l {
			memo[n] = 1
			return 1
		}

		result := 0

		for i := 1; i <= mx; i++ {
			if nn := n + i; nn <= l {
				result += dfs(l, nn, mx)
			}
		}

		memo[n] = result % 1000000007

		return result % 1000000007
	}

	for _, value := range split {
		count, val := value.count, value.val

		clear(memo)

		result *= dfs(count, 0, keys[val])

		result %= 1000000007
	}

	return result
}
