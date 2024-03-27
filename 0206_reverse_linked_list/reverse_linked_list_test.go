package reverselinkedlist

import (
	"reflect"
	"testing"
)

// https://github.com/keep-practicing/leetcode-go/blob/master/solutions/0206_reverse_linked_list/reverse_linked_list_test.go
// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0206.reverse-linked-list/reverse-linked-list.go
func Test_reverseLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLinkedList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLinkedList1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLinkedList1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseLinkedList1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createSinglyLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	currnt := head

	for _, num := range nums {
		currnt.next = &ListNode{key: num}
		currnt = currnt.next
	}
	return head.next
}
