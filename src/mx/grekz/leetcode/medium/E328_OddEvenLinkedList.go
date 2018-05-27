package medium

// @author grekz
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head != nil && head.Next != nil && head.Next.Next != nil {
        od, ev, evHe := head, head.Next, head.Next
        for ev != nil && ev.Next != nil {
            od.Next = od.Next.Next
            ev.Next = ev.Next.Next
            od = od.Next
            ev = ev.Next
        }
        od.Next = evHe
    }
    return head
}