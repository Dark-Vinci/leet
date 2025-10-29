package arrays

import (
	"cmp"
	"slices"
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var (
		db     = make([][2]int, 0)
		result = make([]int, 0)
	)

	for i := 0; i < len(restaurants); i++ {
		res := restaurants[i]
		id, rating, vegan, price, dist := res[0], res[1], res[2], res[3], res[4]

		if vegan >= veganFriendly && price <= maxPrice && dist <= maxDistance {
			db = append(db, [...]int{id, rating})
		}
	}

	slices.SortFunc(db, func(a, b [2]int) int {
		return cmp.Or(cmp.Compare(b[1], a[1]), cmp.Compare(b[0], a[0]))
	})

	for i := 0; i < len(db); i++ {
		result = append(result, db[i][0])
	}

	return result
}
