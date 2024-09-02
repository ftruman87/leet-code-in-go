package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tcs := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{1, 2, 4, 6}, expected: []int{48, 24, 12, 8}},
		{nums: []int{-1, 0, 1, 2, 3}, expected: []int{0, -6, 0, 0, 0}},
	}

	for _, tc := range tcs {
		actual := ProductExceptSelf(tc.nums)

		assert.Equal(t, tc.expected, actual)
	}
}
