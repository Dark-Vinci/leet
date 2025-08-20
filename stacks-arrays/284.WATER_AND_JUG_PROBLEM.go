package stacks_arrays

func canMeasureWater(x int, y int, target int) bool {
	var (
		helper func(x, y int) bool
		seen   = make(map[[2]int]struct{})
	)

	helper = func(xx, yy int) bool {
		key := [2]int{xx, yy}

		if _, ok := seen[key]; ok {
			return false
		}

		seen[key] = struct{}{}

		if xx == target || yy == target || xx+yy == target {
			return true
		}

		// fill x,
		if xx == 0 {
			if helper(x, yy) {
				return true
			}
		}

		// fill y
		if yy == 0 {
			if helper(xx, y) {
				return true
			}
		}

		// empty x
		if xx > 0 {
			if helper(0, yy) {
				return true
			}
		}

		// empty y
		if yy > 0 {
			if helper(xx, 0) {
				return true
			}
		}

		// pour some y into x
		if yy > 0 && xx < x {
			if xx+yy <= x {
				if helper(xx+yy, 0) {
					return true
				}
			} else {
				if helper(x, xx+yy-x) {
					return true
				}
			}
		}

		// pour some x into y
		if xx > 0 && yy < y {
			if xx+yy <= y {
				if helper(0, xx+yy) {
					return true
				}
			} else {
				if helper(xx+yy-y, y) {
					return true
				}
			}
		}

		return false
	}

	return helper(0, 0)
}
