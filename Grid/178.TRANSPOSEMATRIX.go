package Grid

func transpose(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))

	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(matrix))
	}

	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[0]); i++ {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}
