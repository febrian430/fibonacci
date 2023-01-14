package fibonacci_test

import (
	"recursion-demo/fibonacci"
	"testing"
)

const nthFib = 30

func BenchmarkTestRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Recursive(nthFib)
	}
}

func BenchmarkTestMemoized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Memoized(nthFib)
	}
}
