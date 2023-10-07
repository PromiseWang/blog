package day04

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 定义两个新的头，指向传参的头
    newHeadA := &ListNode {
        Val: 0,
        Next: headA,
    }
    newHeadB := &ListNode {
        Val: 0,
        Next: headB,
    }
    // 求两个链表各自的长度
    p := newHeadA.Next
    lenA := 0
    q := newHeadB.Next
    lenB := 0
    for p != nil {
        p = p.Next
        lenA++
    }
    for q != nil {
        q = q.Next
        lenB++
    }
    
    // 定义两个新的指针，进行“对齐”操作
    p1 := newHeadA.Next
    q1 := newHeadB.Next
    if lenA < lenB {
        for i := 0; i < lenB - lenA; i++ {
            q1 = q1.Next
        }
    } else {
        for i := 0; i < lenA - lenB; i++ {
            p1 = p1.Next
        }
    }
    // 对齐后向前查找，找到了返回结点，否则为nil
    for p1 != nil && q1 != nil {
        if p1 == q1 {
            return p1
        } else {
            p1 = p1.Next
            q1 = q1.Next
        }
    }
    return nil
}