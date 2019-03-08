package medium

// @author grekz
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return 
    }
    // get to the middle, reverse second part, mix parts
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    fast = slow.Next
    for fast.Next != nil {
        cur := fast.Next
        fast.Next = cur.Next
        cur.Next = slow.Next
        slow.Next = cur
    }
    fast = slow.Next
    mid := slow
    slow = head
    for slow != mid {
        mid.Next = fast.Next
        fast.Next = slow.Next
        slow.Next = fast
        slow = fast.Next
        fast = mid.Next
    }
}