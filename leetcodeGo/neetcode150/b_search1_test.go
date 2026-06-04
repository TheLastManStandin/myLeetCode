package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{-1, 0, 2, 4, 6, 8},
				target: 4,
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{-1, 0, 2, 4, 6, 8},
				target: 6,
			},
			want: 4,
		},
		{
			name: "case 6",
			args: args{
				nums:   []int{-1, 0, 2, 4, 6, 8},
				target: 3,
			},
			want: -1,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{0, 2},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{0, 2, 5},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 7",
			args: args{
				nums:   []int{0, 2, 5, 7},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 8",
			args: args{
				nums:   []int{0, 2, 5, 7, 9, 10, 11},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 9",
			args: args{
				nums:   []int{0, 2, 5, 7, 9, 10, 11},
				target: 10,
			},
			want: 5,
		},

		{
			name: "case 10",
			args: args{
				nums:   []int{0, 2, 5, 7, 9, 10, 11, 12},
				target: 10,
			},
			want: 5,
		},

		{
			name: "case 11",
			args: args{
				nums:   []int{0, 2, 5, 7, 9, 10, 11, 12, 13, 15},
				target: 10,
			},
			want: 5,
		},
		{
			name: "case 12",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
