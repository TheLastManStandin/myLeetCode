package sliding_window

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "OUZODYXAZV",
				t: "XYZ",
			},
			want: "YXAZ",
		},
		{
			name: "case 2",
			args: args{
				s: "OUZODYX",
				t: "XYZ",
			},
			want: "ZODYX",
		},
		{
			name: "case 3",
			args: args{
				s: "ooooxooooyoooooozoox",
				t: "xyz",
			},
			want: "yoooooozoox",
		},
		{
			name: "case 4",
			args: args{
				s: "xyoooyoooooozoox",
				t: "xyz",
			},
			want: "yoooooozoox",
		},
		{
			name: "case 5",
			args: args{
				s: "xy",
				t: "xy",
			},
			want: "xy",
		},
		{
			name: "case 6",
			args: args{
				s: "xy",
				t: "xyz",
			},
			want: "",
		},
		{
			name: "case 7",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
