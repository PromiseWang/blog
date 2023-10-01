package day04


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := newHead
	for i := 0; i <= n; i++ {
		p = p.Next
	}
	q := newHead
	for p != nil {
		p = p.Next
		q = q.Next
	}
	if q.Next.Next != nil {
		q.Next = q.Next.Next
	} else {
		q.Next = nil
	}

	return newHead.Next
}