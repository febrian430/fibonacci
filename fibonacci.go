package fibonacci

func Memoized(nth int) int {
	fib := []int{0, 1}
	if nth <= 2 {
		return fib[nth-1]
	}

	for i := 2; i <= nth; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}

	return fib[nth-1]
}

func Recursive(nth int) int {
	if nth <= 2 {
		return nth - 1
	}

	return Recursive(nth-2) + Recursive(nth-1)
}
