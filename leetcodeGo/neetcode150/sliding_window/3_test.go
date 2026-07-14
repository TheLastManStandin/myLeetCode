package sliding_window

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "abcdfghij",
				k: 0,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				s: "abcdffghij",
				k: 0,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				s: "abba",
				k: 2,
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				s: "AAABABB",
				k: 1,
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				s: "ABABABA",
				k: 3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
