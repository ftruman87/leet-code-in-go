package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tcs := []struct {
		strs     []string
		expected [][]string
	}{
		{strs: []string{"act", "pots", "tops", "cat", "stop", "hat"}, expected: [][]string{
			{"hat"}, {"act", "cat"}, {"pots", "tops", "stop"},
		}},
		{strs: []string{"x"}, expected: [][]string{
			{"x"},
		}},
		{strs: []string{""}, expected: [][]string{{""}}},
	}

	for _, tc := range tcs {
		res := GroupAnagram(tc.strs)

		assert.ElementsMatch(t, tc.expected, res)
	}
}
