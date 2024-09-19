package main

func TopKElements(nums []int, k int) []int {
	count := make(map[int]int)

	freq := make([][]int, len(nums)+1)

	var res []int

	for _, num := range nums {
		count[num] += 1
	}

	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	for i := len(freq) - 1; i > 0; i-- {
		res = append(res, freq[i]...)

		if len(res) == k {
			return res
		}
	}

	return res
}
