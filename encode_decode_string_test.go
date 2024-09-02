package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeDecodeStrings(t *testing.T) {
	tcs := []struct {
		strs     []string
		expected []string
	}{
		{strs: []string{"neet", "code", "love", "you"}, expected: []string{"neet", "code", "love", "you"}},
		{strs: []string{"we", "say", ":", "yes"}, expected: []string{"we", "say", ":", "yes"}},
	}

	for _, tc := range tcs {
		actual := EncodeDecodeStrings(tc.strs)

		assert.Equal(t, tc.expected, actual)
	}
}
