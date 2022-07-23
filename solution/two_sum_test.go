package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   []int
		name   string
	}

	tests := []test{
		{
			nums:   []int{1, 5, 6, 3, 2},
			target: 9,
			want:   []int{2, 3},
			name:   "Two Sum 001",
		},
		{
			nums:   []int{5, 23, 15, 14, 100},
			target: 123,
			want:   []int{1, 4},
			name:   "Two Sum 002",
		},
		{
			nums:   []int{5, 23, 15, 14, 100},
			target: 122,
			want:   []int{0, 0},
			name:   "Two Sum 003",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := TwoSum(tc.nums, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nGot:\t %v\nWant:\t %v", got, tc.want)
			}
		})
	}
}
