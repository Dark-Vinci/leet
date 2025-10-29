package arrays

func minJumps(arr []int) int {
	var (
		used   = make(map[int]struct{})
		arrMap = make(map[int][]int)
	)

	for i, v := range arr {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = make([]int, 0)
		}

		arrMap[v] = append(arrMap[v], i)
	}

	q := [][2]int{{0, 0}}

	used[0] = struct{}{}

	for len(q) > 0 {
		l := len(q)

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]

			index, count := c[0], c[1]
			indexP1, indexM1 := index+1, index-1
			nextCount := count + 1

			if indexP1 == len(arr) {
				return count
			}

			for _, j := range arrMap[arr[index]] {
				if _, ok := used[j]; !ok && j != index {
					used[j] = struct{}{}
					q = append(q, [2]int{j, nextCount})
				}
			}

			delete(arrMap, arr[index])

			if _, ok := used[indexP1]; !ok && indexP1 < len(arr) {
				used[indexP1] = struct{}{}
				q = append(q, [2]int{indexP1, nextCount})
			}

			if _, ok := used[indexM1]; !ok && indexM1 > 0 {
				used[indexM1] = struct{}{}
				q = append(q, [2]int{indexM1, nextCount})
			}
		}
	}

	return -1
}
