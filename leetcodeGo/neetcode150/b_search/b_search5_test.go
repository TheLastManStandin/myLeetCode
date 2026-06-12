package b_search

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
				nums:   []int{3, 4, 5, 6, 1, 2},
				target: 1,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{3, 5, 6, 0, 1, 2},
				target: 4,
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{2, 3, 4, 5, 6, 1},
				target: 1,
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{2},
				target: 2,
			},
			want: 0,
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{2, 3, 4, 5, 6, 1},
				target: 2,
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
