package maths

func addBinary(a string, b string) string {
	var (
		result = ""
		i      = 0
		carry  = false
	)

	for {
		var (
			aIndex = len(a) - 1 - i
			bIndex = len(b) - 1 - i
			aLast  string
			bLast  string
		)

		if aIndex < 0 && bIndex < 0 {
			break
		}

		if aIndex >= 0 {
			aLast = string(a[aIndex])
		}

		if bIndex >= 0 {
			bLast = string(b[bIndex])
		}

		// this means be is not
		if aLast == "" {
			if carry {
				if bLast == "1" {
					result = "0" + result
					carry = true
				} else {
					result = "1" + result
					carry = false
				}
			} else {
				result = bLast + result
			}

			i++
			continue
		}

		if bLast == "" {
			if carry {
				if aLast == "1" {
					result = "0" + result
					carry = true
				} else {
					result = "1" + result
					carry = false
				}
			} else {
				result = aLast + result
			}

			i++
			continue
		}

		// case one is empty, handle
		if aLast == "0" && bLast == "0" {
			if carry {
				result = "1" + result
				carry = false
			} else {
				result = "0" + result
				carry = false
			}
		} else if aLast == "0" || bLast == "0" {
			if carry {
				result = "0" + result
				carry = true
			} else {
				carry = false
				result = "1" + result
			}
		} else if aLast == "1" && bLast == "1" {
			if carry {
				carry = true
				result = "1" + result
			} else {
				carry = true
				result = "0" + result
			}
		}

		i++
	}

	if carry {
		result = "1" + result
	}

	return result
}
