package day03

func addAtHead(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	newNode := &ListNode{
		Val:  val,
		Next: newHead.Next,
	}
	newHead.Next = newNode
	return newHead.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head.Next
	head.Next = nil
	for p != nil {
		head = addAtHead(head, p.Val)
		p = p.Next
	}
	return head
}