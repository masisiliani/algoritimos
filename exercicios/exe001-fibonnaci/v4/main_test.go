package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fibonacci(t *testing.T) {
	tests := []struct {
		value    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{14, 377},
		{30, 832040},
		{35, 9227465},
		{40, 102334155},
		{45, 1134903170},
		{50, 12586269025},
	}

	for _, testCase := range tests {
		t.Run(fmt.Sprintf("F%d should be return %d", testCase.value, testCase.expected), func(t *testing.T) {
			result := Fibonacci(testCase.value)
			assert.Equal(t, testCase.expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", testCase.expected, result))
		})
	}
}
