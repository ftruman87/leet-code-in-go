package main

func TwoIntegerSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	resMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if _, ok := resMap[num]; ok {
			return []int{resMap[num], i}
		}
		resMap[diff] = i
	}
	return []int{0, 1}
}
