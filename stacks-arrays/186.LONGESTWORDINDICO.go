package stacks_arrays

import (
	"cmp"
	"slices"
)

func longestWord(words []string) string {
	result, maxCount := "", 0

	slices.SortFunc(words, func(a, b string) int {
		return cmp.Or(cmp.Compare(len(a), len(b)), cmp.Compare(a, b))
	})

	db := make(map[string]struct{})

	for i := 0; i < len(words); i++ {
		db[words[i]] = struct{}{}
	}

	for i := 0; i < len(words); i++ {
		li, word, count := len(words[i]), words[i], 0

		for j := 0; j < li; j++ {
			mSome := []byte(word)[:j+1]

			if _, ok := db[string(mSome)]; ok {
				count++
			} else {
				break
			}
		}

		if (count == maxCount && result > word) || maxCount < count {
			maxCount = count
			result = word
		}
	}

	return result
}
