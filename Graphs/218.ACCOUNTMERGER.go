package Graphs

import (
	"slices"
)

func accountsMerge(accounts [][]string) [][]string {
	var (
		db         = make(map[string]int) // email -> index
		u          = NewUnionFind(len(accounts))
		result     = make([][]string, 0)
		newAccount = make(map[int][]string) //parent -> email list
	)

	for i, val := range accounts {
		for j, email := range val {
			if j > 0 {
				if _, ok := db[email]; !ok {
					db[email] = i
				} else {
					u.Union(i, db[email])
				}
			}
		}
	}

	for email, index := range db {
		parent := u.Find(index)

		if _, ok := newAccount[parent]; !ok {
			newAccount[parent] = []string{accounts[parent][0]}
		}

		newAccount[parent] = append(newAccount[parent], email)
	}

	for _, val := range newAccount {
		slices.Sort(val[1:])
		result = append(result, val)
	}

	return result
}
