package main

import (
	"strconv"
)

func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)

    res = append(res, []int{nums[0]})

    if len(nums) == 1 {
        return res
    }

    for j := 1; j < len(nums); j++ {
        updatedPermutation := make([][]int, 0)

        for _, permutation := range res {
            for i := 0; i <= len(permutation); i++ {
                b := permutation[:i]
                a := permutation[i:]

                newPermutation := make([]int, 0)

                newPermutation = append(newPermutation, b...)
                newPermutation = append(newPermutation, nums[j])
                newPermutation = append(newPermutation, a...)

                updatedPermutation = append(updatedPermutation, newPermutation)
            }
        }

        res = updatedPermutation
    }

    return removeDouble(res)
}

func removeDouble(a [][]int) [][]int {
    check := make(map[string]struct{})
    result := make([][]int, 0)

    for _, v := range a {
        a := ""

        for _, vv := range v {
            s := strconv.Itoa(vv)

            a += s
        }

        if _, ok := check[a]; !ok {
            result = append(result, v)
            check[a] = struct{}{}
        }
    }

    return result
}