package yandex

import "testing"

func TestMinSumAE(t *testing.T) {
	type args struct {
		n int
		k int
		a []int64
		b []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				n: 3,
				k: 0,
				a: []int64{1, 2, 3},
				b: []int64{4, 5, 7},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				n: 3,
				k: 1,
				a: []int64{1, 2, 3},
				b: []int64{4, 5, 7},
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				n: 4,
				k: 1,
				a: []int64{1, 2, 1, 2},
				b: []int64{5, 6, 7, 8},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSumAE(tt.args.n, tt.args.k, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinSumAE() = %v, want %v", got, tt.want)
			}
		})
	}
}
