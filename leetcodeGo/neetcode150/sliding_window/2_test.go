package sliding_window

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "zxyzxyz",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "zzzzz",
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				s: "z",
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "case 6",
			args: args{
				s: "qwererwq",
			},
			want: 4,
		},
		{
			name: "case 7",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "case 8",
			args: args{
				s: "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsovert",
				//            i  14
			},
			want: 17,
		},
		{
			name: "case 9",
			args: args{
				s: "pwwkewka",
			},
			want: 4,
		},
		{
			name: "case 10",
			args: args{
				s: "pwwkewwa",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
