package day04

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := newHead
	for p.Next != nil && p.Next.Next != nil {
		tempNode1 := &ListNode{
			Val:  p.Next.Next.Val,
			Next: nil,
		}
		tempNode2 := &ListNode{
			Val:  p.Next.Val,
			Next: p.Next.Next.Next,
		}
		tempNode1.Next = tempNode2
		p.Next = tempNode1
		p = p.Next.Next
	}
	return newHead.Next
}