package solution

import (
	"math"
	"reflect"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	type test struct {
		val  int
		want int
		name string
	}

	tests := []test{
		{
			val:  123,
			want: 321,
			name: "Reverse Integer 001",
		},
		{
			val:  -456,
			want: -654,
			name: "Reverse Integer 002",
		},
		{
			val:  int(math.Pow(2, 33)),
			want: 0,
			name: "Reverse Integer 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ReverseInteger(tc.val); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %d\nWant:\t %d", got, tc.want)
			}
		})
	}
}
