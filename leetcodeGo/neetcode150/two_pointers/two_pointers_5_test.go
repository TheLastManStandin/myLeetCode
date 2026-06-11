package two_pointers

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				height: []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			},
			want: 9,
		},
		{
			name: "case 2",
			args: args{
				height: []int{0, 2, 0},
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				height: []int{5, 4, 1, 2},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				height: []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3},
			},
			want: 83,
		},
		{
			name: "case 5",
			args: args{
				height: []int{5, 3, 2, 1, 4},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
