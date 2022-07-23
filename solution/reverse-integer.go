package solution

import "math"

func ReverseInteger(num int) int {
	reversed := 0
	maxVal := int(math.Pow(2, 31)) - 1
	minVal := -int(math.Pow(2, 31))

	for num != 0 {
		pop := num % 10
		num /= 10
		reversed = (reversed * 10) + pop

		if reversed < minVal || reversed > maxVal {
			return 0
		}
	}

	return reversed
}
