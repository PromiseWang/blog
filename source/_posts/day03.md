---
title: 代码随想录算法训练营第三天| 203.移除链表元素、707.设计链表、206.反转链表。
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 链表
categories: 刷题
abbrlink: bd069431
date: 2023-09-23 00:08:21
---

# 代码随想录算法训练营第三天| 203.移除链表元素、707.设计链表、206.反转链表。

## 203. 移除链表元素

> 题目链接：[leetcode题目链接](https://leetcode.cn/problems/remove-linked-list-elements/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0203.%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0.html)
>
> 视频讲解：[链表基础操作| LeetCode：203.移除链表元素](https://www.bilibili.com/video/BV18B4y1s7R9)
>
> 状态：AC

### 思路

删除元素还是比较简单的，假设`q = p.Next`，如果删除`q`则是`p.Next = p.Next.Next`，考虑下`p.Next.Next`是否存在即可

### 代码

``` go
func removeElements(head *ListNode, val int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	p := newHead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return newHead.Next
}
```



## 707. 设计链表

> 题目链接：[leetcode题目链接](https://leetcode.cn/problems/design-linked-list/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0707.%E8%AE%BE%E8%AE%A1%E9%93%BE%E8%A1%A8.html)
>
> 视频讲解：[帮你把链表操作学个通透！LeetCode：707.设计链表](https://www.bilibili.com/video/BV1FU4y1X7WD)
>
> 状态：半AC（AC了但是也没完全AC）

### 思路

- `(this *MyLinkedList)get(index int) int`: 先判断index是否合法，所以需要引入一个新的成员对象`size`在`MyLinkedList`中
- `(this *MyLinkedList)AddAtHead(val int)`: 头插法只需要在新的头之后接入一个新的节点，然后`size++`
- `(this *MyLinkedList)AddAtTail(val int)`: 尾插法只需要在整个链表之后接入一个新的节点，然后`size++`
- `(this *MyLinkedList)DeleteAtIndex(index int)`: 删除一个元素类似上一题，只不过要判断index是否合法
- `(this *MyLinkedList)AddAtIndex(index int, val)`: 先判断index是否合法，找到正确位置进行插入

### 代码

``` go

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

```




## 206. 反转链表

> 题目链接：[leetcode题目链接](https://leetcode.cn/problems/reverse-linked-list/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0206.%E7%BF%BB%E8%BD%AC%E9%93%BE%E8%A1%A8.html)
>
> 视频讲解：[帮你拿下反转链表 | LeetCode：206.反转链表](https://www.bilibili.com/video/BV1nB4y1i7eL)
>
> 状态：AC

### 思路

1. 建立新的头结点`newHead` ，使其`newHead.Next = head`。
2. 定义指针`p = head.Next`(先判断head是否为单节点)，然后手动断链，`head.Next = nil`
3. 指针`p`不断向后，每遍历到一个元素，将这个元素保存到新的节点`newNode`，并且不断头插到`newHead`

### 代码

``` go
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
```



## 小结

- leetcode上的链表的题都是无头链表，所谓的 `head` 被叫做 `虚拟头`。我感觉还是叫做无头链表好一些，新建一个头先链接无头链表

- 以前写过的都是有头链表，换到了无头链表刚开始有点不知所措、插入时候感觉怪怪的，其实都一样

- `Go`的`new`相关或者初始化成员变量有了新的理解

    ``` go
    type SSS struct {
        Val int
    }
    
    //初始化的方式1
    a := new SSS(1)
    a.Val = 1
    
    //初始化的方式2
    a := &SSS{
        Val: 1
    }
    
    ```

- `Go`的构造函数看起来还有点懵
- 最近事情太多了，周六上午招聘会、下午学校组织的求职训练营、晚上继续leetcode；周日上午训练营结课、下午再重写下链表、搞搞课题；周一还有个小活动。加油！
