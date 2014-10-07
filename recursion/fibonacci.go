package recursion

/*
import (
	"fmt"
)
*/

func FibonacciRecursive(n int64) int64 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		// fmt.Printf("%v ", 1)
		return 1
	}

	fib := FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	// fmt.Printf("%v ", fib)
	return fib
}

func FibonacciSequential(n int64) int64 {
	if n == 0 {
		return 0
	}

	a, b := int64(1), int64(1)
	for i := int64(3); i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}

	return b
}
