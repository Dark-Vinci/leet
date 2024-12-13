package main

import "slices"

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return n
	}

	list := toAdj(trust)
	res, count := -1, 0

	for key, val := range list {
		if len(val) == 0 {
			res = key
			break
		}
	}

	if res < 0 {
		return res
	}

	for _, val := range list {
		if slices.Contains(val, res) {
			count++
		}
	}

	if count == n-1 {
		return res
	}

	return -1
}

func toAdj(trust [][]int) map[int][]int {
	result := make(map[int][]int)

	for _, tr := range trust {
		u, v := tr[0], tr[1] // trustee, trusted

		if _, ok := result[u]; !ok {
			result[u] = make([]int, 0)
		}

		if _, ok := result[v]; !ok {
			result[v] = make([]int, 0)
		}

		result[u] = append(result[u], v)
	}

	return result
}
