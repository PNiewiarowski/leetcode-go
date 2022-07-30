package solution

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type test struct {
		n    int
		want bool
		name string
	}

	tests := []test{
		{
			n:    155551,
			want: true,
			name: "IsPalindrome 001",
		},
		{
			n:    1556551,
			want: true,
			name: "IsPalindrome 002",
		},
		{
			n:    1110,
			want: false,
			name: "IsPalindrome 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %t\nWant:\t %t", got, tc.want)
			}
		})
	}
}
