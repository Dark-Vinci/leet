package main

func removeAnagrams(words []string) []string {
	result := make([]string, 0)

	for i := 0; i < len(words); i++ {
		if i == 0 || (i > 0 && !isAnagram(words[i], words[i-1])) {
			result = append(result, words[i])
		}
	}

	return result
}
