package b_search

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				piles: []int{1, 4, 3, 2},
				h:     9,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				piles: []int{25, 10, 23, 4},
				h:     4,
			},
			want: 25,
		},
		{
			name: "case 3",
			args: args{
				piles: []int{228228},
				h:     228227,
			},
			want: 2,
		},
		{
			name: "case 4",
			args: args{
				piles: []int{192929},
				h:     292929,
			},
			want: 1,
		},

		{
			name: "case 4",
			args: args{
				piles: []int{192929, 192929},
				h:     292929,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
