package stacks_arrays

import "math"

type p [2]int // val, is_back

func minimumJumps1(forbidden []int, a int, b int, x int) int {
	fMap, used := make(map[int]struct{}), make(map[p]struct{})

	mx := forbidden[0]

	for _, v := range forbidden {
		mx = max(mx, v)
		fMap[v] = struct{}{}
	}

	boundary, result := x+a+b+mx, 0

	q := []p{{0, 0}}

	for len(q) > 0 {
		l := len(q)

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			if c[0] == x {
				return result
			}

			next := [2]int{c[0] + a, 0}
			if _, ok := fMap[next[0]]; !ok {
				if _, ko := used[next]; !ko {
					if next[0] < boundary {
						used[next] = struct{}{}
						q = append(q, next)
					}
				}
			}

			back := [2]int{c[0] - b, 1}
			if _, ok := fMap[back[0]]; !ok && c[1] == 0 {
				if _, ko := used[back]; !ko {
					if back[0] >= 0 {
						used[back] = struct{}{}
						q = append(q, back)
					}
				}
			}
		}

		result++
	}

	return -1
}

// WHY DO DFS NOT WORK FOR SOME QUESTIONS
func minimumJumps_(forbidden []int, a int, b int, x int) int {
	mn := math.MaxInt

	var dfs func(back bool, val int, count int)

	fMap, used := make(map[int]struct{}), make(map[int]struct{})

	mx := forbidden[0]

	for _, v := range forbidden {
		mx = max(mx, v)
		fMap[v] = struct{}{}
	}

	dfs = func(back bool, val int, count int) {
		if val < 0 || count > mn {
			return
		}

		if _, ok := fMap[val]; ok {
			return
		}

		if val >= mx+x+a+b {
			return
		}

		if val == x {
			mn = min(mn, count)
			return
		}

		if _, ok := used[val]; ok {
			return
		}

		used[val] = struct{}{}

		if !back {
			dfs(true, val-b, count+1)
		}

		dfs(false, val+a, count+1)
	}

	dfs(false, 0, 0)

	if mn == math.MaxInt {
		return -1
	}

	return mn
}
