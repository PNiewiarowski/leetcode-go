package solution

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type test struct {
		val  string
		want int
		name string
	}

	tests := []test{
		{
			val:  "III",
			want: 3,
			name: "Roman To Integer 001",
		},
		{
			val:  "IX",
			want: 9,
			name: "Roman To Integer 002",
		},
		{
			val:  "XIII",
			want: 13,
			name: "Roman To Integer 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RomanToInt(tc.val); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %d\nWant:\t %d", got, tc.want)
			}
		})
	}
}
