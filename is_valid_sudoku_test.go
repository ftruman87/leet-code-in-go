package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	tcs := []struct {
		board    [][]string
		expected bool
	}{
		{
			board: [][]string{
				{"1", "2", ".", ".", "3", ".", ".", ".", "."},
				{"4", ".", ".", "5", ".", ".", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", ".", "3"},
				{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
				{".", ".", ".", "8", ".", "3", ".", ".", "5"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", ".", ".", ".", ".", ".", "2", ".", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "8"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			}, expected: true,
		},
		{
			board: [][]string{
				{"1", "2", ".", ".", "3", ".", ".", ".", "."},
				{"4", ".", ".", "5", ".", ".", ".", ".", "."},
				{".", "9", "1", ".", ".", ".", ".", ".", "3"},
				{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
				{".", ".", ".", "8", ".", "3", ".", ".", "5"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", ".", ".", ".", ".", ".", "2", ".", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "8"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			}, expected: false,
		},
	}

	for _, tc := range tcs {
		actual := IsValidSudoku(tc.board)

		assert.Equal(t, tc.expected, actual)
	}
}
