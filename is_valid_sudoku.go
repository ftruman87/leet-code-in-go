package main

import (
	"slices"
	"strconv"
)

func IsValidSudoku(board [][]string) bool {
	rows := make([][]string, 9)
	columns := make([][]string, 9)
	subgrids := make(map[string][]string)

	for i, r := range board {
		for j, num := range r {
			if num == "." {
				continue
			}
			subgridKey := strconv.Itoa(i/3) + strconv.Itoa(j/3)
			if slices.Contains(rows[i], num) ||
				slices.Contains(columns[j], num) ||
				slices.Contains(subgrids[subgridKey], num) {
				return false
			} else {
				rows[i] = append(rows[i], num)
				columns[j] = append(columns[j], num)
				subgrids[subgridKey] = append(subgrids[subgridKey], num)
			}
		}
	}

	return true
}
