package solution

import (
	"math"
	"testing"
)

type testCase struct {
	val  int
	want int
}

func TestReverseInteger(t *testing.T) {
	testCases := []testCase{
		{
			val:  123,
			want: 321,
		},
		{
			val:  -456,
			want: -654,
		},
		{
			val:  int(math.Pow(2, 33)),
			want: 0,
		},
	}

	for _, tc := range testCases {
		got := ReverseInteger(tc.val)
		if got != tc.want {
			t.Errorf("\nGot:\t %d\nWant:\t %d", got, tc.want)
		}
	}
}
