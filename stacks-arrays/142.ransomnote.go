package arrays

func canConstruct(ransomNote string, magazine string) bool {
	db := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		db[magazine[i]] += 1
	}

	for i := 0; i < len(ransomNote); i++ {
		if val, ok := db[ransomNote[i]]; ok && val > 0 {
			db[ransomNote[i]]--
		} else {
			return false
		}
	}

	return true
}
