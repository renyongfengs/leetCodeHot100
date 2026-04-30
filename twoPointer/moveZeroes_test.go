package twoPointers

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{nums: []int{0, 1, 0, 3, 12}},
		}, {
			name: "Test Case 2",
			args: args{nums: []int{0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
