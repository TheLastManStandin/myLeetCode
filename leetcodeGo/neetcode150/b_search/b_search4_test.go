package b_search

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 4, 5, 6, 1, 2},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{2, 3, 4, 5, 6, 1},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
