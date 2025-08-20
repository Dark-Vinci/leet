package maths

import (
	"slices"
	"strconv"
	"strings"
)

func isSameAfterReversals(num int) bool {
	firstReversal := _reverse(num)
	secondReversal := _reverse(firstReversal)

	return num == secondReversal
}

func _reverse(num int) int {
	str := strconv.Itoa(num)

	arrStr := strings.Split(str, "")

	slices.Reverse(arrStr)

	ne := strings.Join(arrStr, "")

	num, _ = strconv.Atoi(ne)

	return num
}
