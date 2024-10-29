package main

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor2() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}

		current = current.children[char]
	}

	current.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}

		current = current.children[char]
	}

	return current.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}

		current = current.children[char]
	}

	return true
}

type Trie0 struct {
	val      string
	str      string
	children []*Trie0
}

func Constructor1() Trie0 {
	return Trie0{
		val:      "",
		str:      "",
		children: []*Trie0{},
	}
}

func (this *Trie0) Insert(word string) {
	dfsInsert(this, word, 0)
}

func dfsInsert(p *Trie0, word string, index int) {
	if index >= len(word) {
		p.children = append(p.children, nil)
		return
	}

	v := &Trie0{
		val:      string(word[index]),
		str:      p.str + string(word[index]),
		children: []*Trie0{},
	}

	if len(p.children) == 0 {
		p.children = append(p.children, v)
		dfsInsert(p.children[0], word, index+1)
		return
	}

	var seen bool

	for i := 0; i < len(p.children); i++ {
		if p.children[i].val == string(word[index]) {
			seen = true
			dfsInsert(p.children[i], word, index+1)
		}

		if i == len(p.children)-1 && !seen {
			p.children = append(p.children, v)
			dfsInsert(p.children[i+1], word, index+1)
		}
	}
}

func dfsSearch(w *Trie0, word string, index int) bool {
	if index >= len(word) {
		return true
	}

	// if w == nil || w.val != string(word[index]) && w.val != "" {
	//     return false
	// }

	if w.str == word && len(w.children) >= 1 {
		for i := 0; i < len(w.children); i++ {
			if w.children[0].val == "." {
				return true
			}
		}
	}

	for i := 0; i < len(w.children); i++ {
		var r bool

		if w.val == "" {
			r = dfsSearch(w.children[i], word, index)
		} else {
			r = dfsSearch(w.children[i], word, index+1)
		}

		if r {
			return true
		}
	}

	return false
}

func (this *Trie0) Search(word string) bool {
	return dfsSearch(this, word, 0)
}

func dfsStartsWith(a *Trie0, word string) bool {
	if a.str == word {
		return true
	}

	for i := 0; i < len(a.children); i++ {
		if dfsStartsWith(a.children[i], word) {
			return true
		}
	}

	return false
}

func (this *Trie0) StartsWith(prefix string) bool {
	return dfsStartsWith(this, prefix)
}
