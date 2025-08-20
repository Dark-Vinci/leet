package maths

func xorAllNums(nums1 []int, nums2 []int) int {
	result := 0

	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			result ^= v
		}
	}

	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			result ^= v
		}
	}

	return result
}

func xorAllNums0(nums1 []int, nums2 []int) int {
	n3 := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			n3 = append(n3, nums1[i]^nums2[j])
		}
	}

	result := n3[0]

	for i := 1; i < len(n3); i++ {
		result ^= n3[i]
	}

	return result
}
