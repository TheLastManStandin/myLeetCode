package sliding_window

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s1: "abc",
				s2: "lecabee",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s1: "abc",
				s2: "ioiioaabbac",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s1: "abc",
				s2: "abb",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s1: "ab",
				s2: "abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
