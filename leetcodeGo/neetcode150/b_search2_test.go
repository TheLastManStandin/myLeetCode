package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 9,
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 55,
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 16,
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 34,
			},
			want: true,
		},
		{
			name: "case 6",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 55,
			},
			want: true,
		},
		{
			name: "case 7",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 204,
			},
			want: true,
		},
		{
			name: "case 8",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: 12,
			},
			want: false,
		},
		{
			name: "case 9",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
					{55, 93, 95, 99},
					{102, 114, 163, 204},
				},
				target: -1,
			},
			want: false,
		},
		{
			name: "case 10",
			args: args{
				matrix: [][]int{
					{-100, -98, -83, -50},
					{-45, -43, -30, -24},
				},
				target: -30,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
