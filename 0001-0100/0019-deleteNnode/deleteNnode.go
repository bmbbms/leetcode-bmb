package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	P := head
	length := 0
	for P != nil {
		P = P.Next
		length++
	}
	P = head
	for i := 0; i < length-n; i++ {
		P = P.Next
	}
	P.Next = P.Next.Next
	P.Val = P.Next.Val
	return head

}
func main() {

}
