package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args1 struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args1
		want *ListNode
	}{
		{
			"test1",
			args1{
				&ListNode{1, &ListNode{2, nil}},
				&ListNode{1, &ListNode{2, nil}},
			},
			&ListNode{2, &ListNode{4, nil}},
		},
		{
			"test2",
			args1{
				&ListNode{9, &ListNode{9, nil}},
				&ListNode{1, &ListNode{2, nil}},
			},
			&ListNode{0, &ListNode{2, &ListNode{1, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
