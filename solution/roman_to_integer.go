package solution

func RomanToInt(s string) int {
	var resolver = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	r := 0
	i := 0
	s = reverse(s)
	for {
		if i == len(s) {
			return r
		}

		v := resolver[string(s[i])]
		if i < len(s)-1 {
			if vNext := resolver[string(s[i+1])]; v > vNext {
				r, i = r-vNext, i+1
			}
		}

		r, i = r+v, i+1
	}
}

func reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
