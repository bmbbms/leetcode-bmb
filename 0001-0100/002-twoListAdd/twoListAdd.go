package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if result == nil {
			result = &ListNode{}
			result.Val = sum

			tail = result
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}

	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}

	return result
}

func main() {
	node := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	node1 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	fmt.Printf("%#v", addTwoNumbers(node, node1))

	//nodes := []ListNode{{Val: 4}, {Val: 8}}
}
