package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fn = Fn-1 + Fn-2
// fibonocci(0) = 0
// fibonocci(1) = 1
// n = 2 => fibonocci(2) -> fibonocci(n-1) + fibonocci(n-2) -> fibonocci(2-1) + fibonocci(2-2) -> fibonocci(1) + fibonocci(0) -> 1+0 = 1
// n = 3 => fibonocci(3) -> fibonocci(n-1) + fibonocci(n-2) -> fibonocci(3-1) + fibonocci(3-2) -> fibonocci(2) + fibonocci(1) -> 1+1 = 2
// fibonocci(4) = 3
// fibonocci(5) = 5
// fibonocci(6) = 8

func Test_Fibonacci(t *testing.T) {

	t.Run("F0 should be return 0", func(t *testing.T) {
		result := Fibonacci(0)
		expected := 0

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F1 should be return 1", func(t *testing.T) {
		result := Fibonacci(1)
		expected := 1

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F2 should be return 2", func(t *testing.T) {
		result := Fibonacci(2)
		expected := 1

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F3 should be return 3", func(t *testing.T) {
		result := Fibonacci(3)
		expected := 2

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 4", func(t *testing.T) {
		result := Fibonacci(4)
		expected := 3

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 5", func(t *testing.T) {
		result := Fibonacci(5)
		expected := 5

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 6", func(t *testing.T) {
		result := Fibonacci(6)
		expected := 8

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 7", func(t *testing.T) {
		result := Fibonacci(7)
		expected := 13

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 14", func(t *testing.T) {
		result := Fibonacci(14)
		expected := 377

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 30", func(t *testing.T) {
		result := Fibonacci(30)
		expected := 832040

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 35", func(t *testing.T) {
		result := Fibonacci(35)
		expected := 9227465

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 40", func(t *testing.T) {
		result := Fibonacci(40)
		expected := 102334155

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 45", func(t *testing.T) {
		result := Fibonacci(45)
		expected := 1134903170

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})

	t.Run("F4 should be return 50", func(t *testing.T) {
		result := Fibonacci(50)
		expected := 12586269025

		assert.Equal(t, expected, result, fmt.Sprintf("expected: %d, result: %d. should be equal", expected, result))
	})
}
