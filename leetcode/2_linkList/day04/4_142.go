package day04

func detectCycle(head *ListNode) *ListNode {
    fast := head
    slow := head
    p := head
    // 这层for寻找环
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
            if slow == fast {  // 两指针相遇说明有环
            for p != fast {  // 这层for执行的是公式中x=z的一步
                p = p.Next
                fast = fast.Next
            }
            return p
            }
        } else {
            break
        }        
    }
    return nil
}