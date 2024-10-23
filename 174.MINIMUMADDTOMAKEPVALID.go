package main

func minAddToMakeValid(s string) int {
	q := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if len(q) == 0 {
			q = append(q, s[i])
			continue
		}

		if q[len(q)-1] == '(' && s[i] == ')' {
			q = q[:len(q)-1]
			continue
		}

		q = append(q, s[i])
	}

	return len(q)
}
