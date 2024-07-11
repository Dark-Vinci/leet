package main

func removeElements(head *ListNode, val int) *ListNode {
    var r func (a, prev *ListNode)

    r = func (a, prev *ListNode) {
        if a == nil {
            return
        }

        if a.Val == val && prev == nil {
            head = a.Next
            r(a.Next, nil)
        } else if a.Val == val && prev != nil {
            prev.Next = a.Next
            r(a.Next, prev)
        } else {
            r(a.Next, a)
        }
    }

    r(head, nil)

    return head
}
