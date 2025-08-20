package stacks_arrays

import (
	"slices"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	slices.Sort(sSlice)
	slices.Sort(tSlice)

	s = strings.Join(sSlice, "")
	t = strings.Join(tSlice, "")

	return s == t
}

func anotherValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	store := [26]int{}

	for _, ch := range s {
		store[ch-'a'] += 1
	}

	for _, ch := range t {
		store[ch-'a'] -= 1
	}

	for _, v := range store {
		if v != 0 {
			return false
		}
	}

	return true
}
