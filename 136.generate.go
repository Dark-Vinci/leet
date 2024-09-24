package main

import (
	"math/rand"
	"slices"
)

type RandomizedSet struct {
	db []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		db: []int{},
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if !slices.Contains(r.db, val) {
		r.db = append(r.db, val)
		return true
	}

	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	ind := slices.Index(r.db, val)

	if ind < 0 {
		return false
	}

	r.db = slices.Delete(r.db, ind, ind+1)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	ind := int(rand.Float64() * float64((len(r.db))))

	return r.db[ind]
}
