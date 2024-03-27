package reverselinkedlist

/*
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/

Reverse a singly linked list.
*/

type ListNode struct {
	key  int
	next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.next
		head.next = prev
		prev = head
		head = tmp
	}
	return prev
}

func reverseLinkedList1(head *ListNode) *ListNode {
	if head == nil || head.next == nil { // For empty or 1 node linked lists
		return head
	}

	// We set the end node as new head node
	// previously end node had next node as nil
	// now set end node's next node (i.e current head node's next node's next node) as current head node to reverse
	// set current head node's next node as nil to make it end node
	newHead := reverseLinkedList1(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}
