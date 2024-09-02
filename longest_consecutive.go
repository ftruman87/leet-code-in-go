package main

func LongestConsecutive(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num] += 1
	}

	c := 1

	for key, _ := range numMap {
		if _, ok2 := numMap[key-1]; !ok2 {
			length := 1
			for {
				if _, ok := numMap[key+length]; !ok {
					break
				} else {
					length += 1
				}
			}
			c = max(c, length)
		}
	}

	return c
}
