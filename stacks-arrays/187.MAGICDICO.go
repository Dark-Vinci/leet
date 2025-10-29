package arrays

type MagicDictionary struct {
	db []string
}

func Constructor0() MagicDictionary {
	return MagicDictionary{
		db: []string{},
	}
}

func (t *MagicDictionary) BuildDict(dictionary []string) {
	t.db = dictionary
}

func (t *MagicDictionary) Search(searchWord string) bool {
	l := len(searchWord)

	for i := 0; i < len(t.db); i++ {
		w := t.db[i]

		if len(w) != l {
			continue
		}

		count := 0

		for j := 0; j < len(w); j++ {
			if w[j] == searchWord[j] {
				count++
			}
		}

		if count+1 == len(w) {
			return true
		}
	}

	return false
}
