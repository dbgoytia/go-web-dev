package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFibonacci(t *testing.T) {

	testGrid := map[uint]uint64{
		0:  0,
		1:  1,
		2:  1,
		5:  5,
		38: 39088169,
		50: 12586269025,
		78: 8944394323791464,
		85: 259695496911122585,
		90: 2880067194370816120,
	}

	for n, fibo := range testGrid {
		res, _ := Fibonacci(n)
		if res != fibo {
			t.Fatalf(`Fibonacci(%v)= %v, want %v, error`, n, res, fibo)
		}
		t.Logf(`Fibonacci(%v)= %v, correct`, n, res)
	}
}
