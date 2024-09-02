package main

import (
	"strconv"
)

func EncodeDecodeStrings(strs []string) []string {
	encoded := encode(strs)
	return decode(encoded)
}

func encode(strs []string) string {
	var retString string
	for _, str := range strs {
		length := len(str)
		retString += strconv.Itoa(length) + "#" + str
	}
	return retString
}

func decode(str string) []string {
	var ret []string
	for i := 0; i < len(str); i++ {
		j := i
		for {
			if string(str[j]) == "#" {
				break
			}
			j++
		}
		length, _ := strconv.Atoi(str[i:j])
		eow := j + 1 + length
		word := str[j+1 : eow]

		ret = append(ret, word)
		i = j + length
	}
	return ret
}
