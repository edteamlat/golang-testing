package clase2

func fibonacciFor(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}

	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}

	return b
}

func fibonacciRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRec(n-1) + fibonacciRec(n-2)
}


var fibonacciData = []int{0, 1}

func fibonacciRecMem(n int) int {
	if n == 0 {
		return fibonacciData[0]
	}

	if len(fibonacciData) >= n+1 {
		return fibonacciData[n]
	}

	newData := fibonacciRecMem(n-1) + fibonacciRecMem(n-2)
	fibonacciData = append(fibonacciData, newData)

	return newData
}