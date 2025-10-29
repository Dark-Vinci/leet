package arrays

import (
	"fmt"
	"maps"
)

func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 { // this can be proved
		return []string{}
	}

	var (
		h, m         = make(map[int]struct{}), make(map[int]struct{})
		minute, hour = []int{1, 2, 4, 8, 16, 32}, []int{1, 2, 4, 8}
		db           = make(map[string]struct{})
		result       = make([]string, 0)
		dfs          func(count int, h, m map[int]struct{})
	)

	dfs = func(count int, h, m map[int]struct{}) {
		if count == turnedOn {
			mm, hh := 0, 0

			for k := range m {
				mm += k
			}

			for k := range h {
				hh += k
			}

			if isValid_(mm, hh) {
				paddedMM := fmt.Sprintf("%v", mm)

				if len(paddedMM) < 2 {
					paddedMM = fmt.Sprintf("0%v", paddedMM)
				}

				db[fmt.Sprintf("%v:%v", hh, paddedMM)] = struct{}{}
			}

			return
		}

		count++

		for _, val := range minute {
			if _, ok := m[val]; !ok {
				mp := maps.Clone(m)
				mp[val] = struct{}{}

				if findValidity(h, mp) {
					dfs(count, h, mp)
				}
			}
		}

		for _, val := range hour {
			if _, ok := h[val]; !ok {
				mp := maps.Clone(h)
				mp[val] = struct{}{}

				if findValidity(mp, m) {
					dfs(count, mp, m)
				}
			}
		}
	}

	dfs(0, h, m)

	for k := range db {
		result = append(result, k)
	}

	return result
}

func findValidity(h, m map[int]struct{}) bool {
	hh, mm := 0, 0

	for k := range h {
		hh += k

		if hh >= 12 {
			return false
		}
	}

	for k := range m {
		mm += k

		if mm >= 60 {
			return false
		}
	}

	return true
}

func isValid_(m int, h int) bool {
	if m >= 60 || m < 0 {
		return false
	}

	if h >= 12 || h < 0 {
		return false
	}

	return true
}
