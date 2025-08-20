package Trees

func numTrees(n int) int {
	db := make([]int, n+1)

	db[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			db[i] += db[j] * db[i-j-1]
		}
	}

	return db[n]
}
