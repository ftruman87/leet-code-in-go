package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKElements(t *testing.T) {
	tcs := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{2, 3},
		},
		{
			nums:     []int{7, 7},
			k:        1,
			expected: []int{7},
		},
	}

	for _, tc := range tcs {
		res := TopKElements(tc.nums, tc.k)

		assert.ElementsMatch(t, tc.expected, res)
	}
}
