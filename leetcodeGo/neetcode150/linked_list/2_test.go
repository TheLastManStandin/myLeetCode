package linked_list

import (
	"reflect"
	"testing"
)

// listToSlice переводит связанный список в срез чисел для удобного сравнения
func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int // Ожидаем плоский срез вместо сложной структуры
	}{
		{
			name: "case1",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
			want: []int{1, 1, 2, 3, 4, 5},
		},
		{
			name: "case2",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			gotSlice := listToSlice(got) // Конвертируем результат в срез

			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
