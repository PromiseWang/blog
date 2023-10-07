package day03

type LinkList struct {
	Val  int
	Next *LinkList
}

type MyLinkedList struct {
	size    int
	newHead *LinkList
}

func Constructor() MyLinkedList {
	//return MyLinkedList{}
	node := &LinkList{
		Val:  0,
		Next: nil,
	}
	return MyLinkedList{
		newHead: node,
		size:    0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size || this == nil {
		return -1
	}
	p := this.newHead.Next
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &LinkList{
		Val:  val,
		Next: this.newHead.Next,
	}
	this.newHead.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &LinkList{
		Val:  val,
		Next: nil,
	}
	p := this.newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.size {
		return
	}
	newNode := &LinkList{
		Val:  val,
		Next: nil,
	}
	p := this.newHead
	for i := 0; i < index; i++ {
		p = p.Next
	}
	newNode.Next = p.Next
	p.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	p := this.newHead
	for i := 0; i < index; i++ {
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	}
	this.size--
}
