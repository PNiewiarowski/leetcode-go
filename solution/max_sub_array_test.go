package solution

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type test struct {
		val  []int
		want int
		name string
	}

	tests := []test{
		{
			val:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
			name: "MaxSubArray 001",
		},
		{
			val:  []int{1},
			want: 1,
			name: "MaxSubArray 002",
		},
		{
			val:  []int{5, 4, -1, 7, 8},
			want: 23,
			name: "MaxSubArray 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxSubArray(tc.val); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %v\nWant:\t %v", got, tc.want)
			}
		})
	}
}
