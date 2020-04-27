func fib(N int) int {
	var fibs [31]int
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i <= N; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	return fibs[N]
}