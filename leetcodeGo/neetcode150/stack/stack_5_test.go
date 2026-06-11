package stack

import "testing"

func Test_carFleet(t *testing.T) {
	type args struct {
		target   int
		position []int
		speed    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				target:   10,
				position: []int{1, 4},
				speed:    []int{3, 2},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				target:   10,
				position: []int{8, 7, 6},
				speed:    []int{2, 6, 2},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				target:   10,
				position: []int{4, 1, 0, 7},
				speed:    []int{2, 2, 1, 1},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				target:   12,
				position: []int{10, 8, 0, 5, 3},
				speed:    []int{2, 4, 1, 1, 3},
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				target:   10,
				position: []int{8, 3, 7, 4, 6, 5},
				speed:    []int{4, 4, 4, 4, 4, 4},
			},
			want: 6,
		},
		{
			name: "case 6",
			args: args{
				target:   10,
				position: []int{0, 4, 2},
				speed:    []int{2, 1, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.args.target, tt.args.position, tt.args.speed); got != tt.want {
				t.Errorf("carFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}
