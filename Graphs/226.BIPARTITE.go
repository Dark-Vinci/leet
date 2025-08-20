package Graphs

func isBipartite(graph [][]int) bool {
	db := make(map[int]string) //red | blue

	type node struct {
		val   int
		color string
	}

	var bfs func(node node) bool

	bfs = func(n node) bool {
		q := []node{n}

		for len(q) > 0 {
			l := len(q)

			for i := 0; i < l; i++ {
				c := q[0]
				q = q[1:]

				for _, val := range graph[c.val] {
					k := node{val: val}
					color, ok := db[val]

					if c.color == "RED" && ok && color == "RED" {
						return false
					} else if c.color == "RED" && ok && color == "BLUE" {
						continue
					}

					if c.color == "BLUE" && ok && color == "BLUE" {
						return false
					} else if c.color == "RED" && ok && color == "BLUE" {
						continue
					}

					if c.color == "RED" {
						k.color = "BLUE"
					} else {
						k.color = "RED"
					}

					db[val] = k.color
					q = append(q, k)
				}
			}
		}

		return true
	}

	for i := range len(graph) {
		if _, ok := db[i]; !ok {
			db[i] = "RED"
			if !bfs(node{val: i, color: "RED"}) {
				return false
			}
		}
	}

	return true
}
