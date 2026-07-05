package b_search

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				nums1: []int{1, 2, 3, 5, 7},
				nums2: []int{2, 4, 6, 8, 16},
			},
			want: 4.5,
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{1, 2, 3, 5, 7},
				nums2: []int{8, 9, 10, 11, 12},
			},
			want: 7.5,
		},
		{
			name: "case 3",
			args: args{
				nums1: []int{8, 9, 10, 11, 12},
				nums2: []int{1, 2, 3, 5, 7},
			},
			want: 7.5,
		},
		{
			name: "case 4",
			args: args{
				nums2: []int{2, 4, 8, 16},
				nums1: []int{1, 2, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "case 5",
			args: args{
				nums1: []int{1, 2, 3, 5, 7},
				nums2: []int{108, 109, 110, 111, 112},
			},
			want: 57.5,
		},
		{
			name: "case 6",
			args: args{
				nums1: []int{1, 2, 3, 5, 7},
				nums2: []int{2, 4, 6, 8, 16, 17},
			},
			want: 5,
		},
		{
			name: "case 7",
			args: args{
				nums1: []int{1, 2, 2, 2, 7},
				nums2: []int{2, 2, 2, 2, 16},
			},
			want: 2,
		},
		{
			name: "case 8",
			args: args{
				nums2: []int{2, 4 /* */, 8, 16},
				nums1: []int{1, 2, 3, 5, 7, 9, 10},
			},
			want: 5,
		},
		{
			name: "case 9",
			args: args{
				nums1: []int{1, 2, 3, 5, 7, 9, 10, 15},
				nums2: []int{2, 4, 8, 16},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
