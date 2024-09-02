package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "racecar", t: "carrace", expected: true},
		{s: "jar", t: "jam", expected: false},
	}

	for _, testCase := range testCases {
		r := IsAnagram(testCase.s, testCase.t)
		assert.Equal(t, testCase.expected, r)
	}
}
