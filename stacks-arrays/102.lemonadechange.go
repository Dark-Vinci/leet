package stacks_arrays

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
			continue
		}

		change := bills[i] - 5
		if change == 5 {
			if five <= 0 {
				return false
			}

			five--
			ten++
		}

		if change == 15 {
			if (five*5 + ten*10) < 15 {
				return false
			}

			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five = five - 3
			} else {
				return false
			}
		}
	}

	return true
}
