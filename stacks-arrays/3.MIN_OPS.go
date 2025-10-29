package arrays

func minimumOperations(num string) int {
	var (
		l     = len(num)
		zeros = 0
		fives = 0
	)

	for i := len(num) - 1; i >= 0; i-- {
		curr := num[i]

		if curr == '0' {
			zeros++

			if zeros == 2 {
				return l - i - 2
			}
		} else if curr == '5' {
			fives++

			if zeros == 1 {
				return l - i - 2
			}
		} else if curr == '2' || curr == '7' {
			if fives > 0 {
				return l - i - 2
			}
		}
	}

	return l - zeros
}
