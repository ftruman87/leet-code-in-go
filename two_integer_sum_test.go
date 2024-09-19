package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoIntegerSum(t *testing.T) {
	tcs := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{3, 4, 5, 6}, target: 7, expected: []int{0, 1}},
		{nums: []int{4, 5, 6}, target: 10, expected: []int{0, 2}},
		{nums: []int{5, 5}, target: 10, expected: []int{0, 1}},
	}

	for _, tc := range tcs {
		res := TwoIntegerSum(tc.nums, tc.target)

		assert.Equal(t, tc.expected, res)
	}
}
