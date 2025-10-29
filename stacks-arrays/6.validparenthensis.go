package arrays

func isValid1(s string) bool {
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	t := make([]string, 0)
	j := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	for _, v := range s {
		switch string(v) {
		case "}", "]", ")":
			if len(t) == 0 || t[len(t)-1] != j[string(v)] {
				return false
			} else {
				t = t[:len(t)-1]
			}
		default:
			t = append(t, string(v))
		}
	}

	return len(t) == 0
}
