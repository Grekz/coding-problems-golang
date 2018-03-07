package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if ( head == nil ) {
        return head
    }
    fh := &ListNode{Val:-1}
    fh.Next = head
    cur := fh
    for cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        }else{
            cur = cur.Next
        }
    }
    return fh.Next
}