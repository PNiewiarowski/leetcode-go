package solution

func Fib(n int) int {
	if n < 0 {
		return 0
	}

	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1

	return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		return val
	}

	val := fibHelper(n-1, memo) + fibHelper(n-2, memo)

	memo[n] = val
	return val
}
