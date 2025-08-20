package stacks_arrays

import "sort"

func canAttendMeetings(a [][]int) bool {
	if len(a) < 2 {
		return true
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	for i := 1; i < len(a); i++ {

		f := a[i][0]
		b := a[i-1][0]

		if b < f {
			return false
		}
	}

	return true
}
