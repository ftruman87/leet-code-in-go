package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tcs := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 20, 4, 10, 3, 4, 5}, expected: 4},
		{nums: []int{0, 3, 2, 5, 4, 6, 1, 1}, expected: 7},
	}

	for _, tc := range tcs {
		actual := LongestConsecutive(tc.nums)

		assert.Equal(t, tc.expected, actual)
	}
}
