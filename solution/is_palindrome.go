package solution

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for y < x {
		pop := x % 10
		x = x / 10
		y = (y * 10) + pop
	}

	return y == x || y/10 == x
}
