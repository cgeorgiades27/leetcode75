package asteroidcollision

import (
	"slices"
	"testing"
)

var tests = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{5, 10, -5},
		expected: []int{5, 10},
	},
	{
		input:    []int{8, -8},
		expected: []int{},
	},
	{
		input:    []int{10, 2, -5},
		expected: []int{10},
	},
}

func TestAsteroidCollision(t *testing.T) {
	for i, test := range tests {
		actual := asteroidCollision(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("Test %d failed. Expected %v but got %v", i+1, test.expected, actual)
			continue
		}
		slices.Sort(actual)
		slices.Sort(test.expected)

		for j := 0; j < len(actual); j++ {
			if actual[j] != test.expected[j] {
				t.Errorf("Test %d failed. Expected %v but got %v", i+1, test.expected, actual)
				break
			}
		}
	}
}
