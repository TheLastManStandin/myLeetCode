package main

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				heights: []int{3, 4, 5, 0},
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				heights: []int{3, 5, 4, 6, 2},
			},
			want: 12,
		},
		{
			name: "case 3",
			args: args{
				heights: []int{7, 1, 7, 2, 2, 4},
			},
			want: 8,
		},
		{
			name: "case 4",
			args: args{
				heights: []int{2, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
