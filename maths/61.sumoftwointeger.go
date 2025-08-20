package maths

func getSum(a int, b int) int {
	var leftOver int

	for b != 0 {
		leftOver = a & b

		a = a ^ b

		b = leftOver << 1
	}

	return a
}
