package main

func ProductExceptSelf(nums []int) []int {
	// [1, 2, 3, 4]
	leftToRight := make([]int, len(nums))
	rightToLeft := make([]int, len(nums))
	ret := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			leftToRight[i] = num
		} else {
			leftToRight[i] = num * leftToRight[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightToLeft[i] = nums[i]
		} else {
			rightToLeft[i] = rightToLeft[i+1] * nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ret[i] = rightToLeft[i+1]
		} else if i == len(nums)-1 {
			ret[i] = leftToRight[i-1]
		} else {
			ret[i] = leftToRight[i-1] * rightToLeft[i+1]
		}
	}

	return ret
}
