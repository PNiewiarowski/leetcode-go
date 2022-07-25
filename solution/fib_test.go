package solution

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	type test struct {
		n    int
		want int
		name string
	}

	tests := []test{
		{
			n:    0,
			want: 0,
			name: "Fib 001",
		},
		{
			n:    2,
			want: 1,
			name: "Fib 002",
		},
		{
			n:    7,
			want: 13,
			name: "Fib 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Fib(tc.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %d\nWant:\t %d", got, tc.want)
			}
		})
	}
}
