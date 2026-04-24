package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				strs: []string{"cat", "tac"},
			},
			want: [][]string{{"cat", "tac"}},
		},
		{
			name: "case 2",
			args: args{
				strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			},
			want: [][]string{{"hat"}, {"act", "cat"}, {"stop", "tops", "pots"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
