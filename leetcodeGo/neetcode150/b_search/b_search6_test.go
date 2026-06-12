package b_search

import "testing"

func Test_bSearchVal(t *testing.T) {
	type args struct {
		timeMap   []timeStamp
		timestamp int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
				},
				timestamp: 7,
			},
			want: "7",
		},
		{
			name: "case 2",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
				},
				timestamp: 8,
			},
			want: "8",
		},
		{
			name: "case 3",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
				},
				timestamp: 9,
			},
			want: "9",
		},
		{
			name: "case 4",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
				},
				timestamp: 5,
			},
			want: "5",
		},
		{
			name: "case 5",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
					{10, "10"},
				},
				timestamp: 7,
			},
			want: "7",
		},
		{
			name: "case 6",
			args: args{
				timeMap: []timeStamp{
					{5, "5"},
					{6, "6"},
					{7, "7"},
					{8, "8"},
					{9, "9"},
					{10, "10"},
				},
				timestamp: 8,
			},
			want: "8",
		},
		{
			name: "case 7",
			args: args{
				timeMap: []timeStamp{
					{10, "5"},
					{20, "6"},
					{30, "7"},
				},
				timestamp: 15,
			},
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchVal(tt.args.timeMap, tt.args.timestamp); got != tt.want {
				t.Errorf("bSearchVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
