package main

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    var recurse func(prev, curr *ListNode)

    recurse = func(prev, curr *ListNode) {
        if curr == nil {
            return
        }

        if prev != nil && prev.Val == curr.Val {
            prev.Next = curr.Next
            recurse(prev, prev.Next)
            return
        }

        prev = curr
        curr = curr.Next

        recurse(prev, curr)
    }

    recurse(nil, head)

    return head
}