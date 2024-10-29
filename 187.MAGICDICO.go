package main

type MagicDictionary struct {
	db []string
}

func Constructor0() MagicDictionary {
	return MagicDictionary{
		db: []string{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.db = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	l := len(searchWord)

	for i := 0; i < len(this.db); i++ {
		w := this.db[i]

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
