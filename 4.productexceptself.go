package main

func productExceptSelf(nums []int) []int {
    l := len(nums)

    lArr := make([]int, l)
    rArr := make([]int, l)
    
    left := 1
    right := 1

    for i := 0; i < l; i++ {
        j := l - 1 - i

        lArr[i] = left
        rArr[j] = right

        left *= nums[i]
        right *= nums[j]
    }

    for i := 0; i < l; i++ {
        a := lArr[i]
        b := rArr[i]

        lArr[i] = a * b
    }

    return lArr
}
