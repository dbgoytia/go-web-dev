package main

import (
	"fmt"
)

// fibonacci returns the nth fibonacci number
func Fibonacci(number uint) (uint64, error) {

	if number < 1 {
		return uint64(0), nil
	}

	// Prevent int overflow
	if number > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", number)
	}

	var prev, n0, n1 uint64 = 0, 0, 1

	for i := uint(2); i <= number; i++ {
		prev = n1
		n1 = n1 + n0
		n0 = prev
	}

	return n1, nil
}
