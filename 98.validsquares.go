package main

import "math"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	arrPoints := make([][]int, 4)

	arrPoints[0] = p1
	arrPoints[1] = p2
	arrPoints[2] = p3
	arrPoints[3] = p4

	dist := make([]float64, 0)

	for i := 0; i < len(arrPoints); i++ {
		for j := 0; j < len(arrPoints); j++ {
			if i != j {
				if arrPoints[i][0] == arrPoints[j][0] && arrPoints[i][1] == arrPoints[j][1] {
					return false
				}

				dist = append(dist, distance(arrPoints[i], arrPoints[j]))
			}
		}
	}

	m := make(map[float64]struct{})

	for _, v := range dist {
		m[v] = struct{}{}
	}

	return len(m) == 2
}

func distance(a, b []int) float64 {
	return math.Sqrt(math.Pow(float64(a[0])-float64(b[0]), 2) + math.Pow(float64(a[1])-float64(b[1]), 2))
}
