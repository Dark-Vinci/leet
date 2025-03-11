package main

func lenLongestFibSubseq(arr []int) int {
	var (
		result int
		dfs    func(x1, x2, count int)
		db     = make(map[int]int) // value | index
		l      = len(arr)
	)

	for k, val := range arr {
		db[val] = k
	}

	dfs = func(x1, x2, count int) {
		result = max(result, count+1)
		sum := arr[x2] + arr[x1]

		if _, ok := db[sum]; !ok {
			return
		}

		index, _ := db[sum]

		dfs(x2, index, count+1)
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum := arr[i] + arr[j]

			if ind, ok := db[sum]; ok {
				result = max(result, 3)
				dfs(j, ind, 2)
			}
		}
	}

	return result
}
