package main

import (
	"fmt"
	"strconv"
)

func GroupAnagram(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{strs}
	}

	temp := make(map[string][]string)
	var res [][]string
	for _, str := range strs {
		var keyString string

		keySlice := make([]int, 26)

		strRuneSlice := []rune(str)
		for _, r := range strRuneSlice {
			keySlice[r-'a'] += 1
		}

		for _, count := range keySlice {
			keyString += strconv.Itoa(count)
		}
		temp[keyString] = append(temp[keyString], str)

	}
	fmt.Println(temp)

	for _, s := range temp {
		res = append(res, s)
	}

	return res
}
