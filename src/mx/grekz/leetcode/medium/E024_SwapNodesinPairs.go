package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    res := new(ListNode)
    res.Next = head
    walk := res
    for( walk.Next != nil && walk.Next.Next != nil ) {
        a := walk.Next
        b := walk.Next.Next
        a.Next = b.Next
        b.Next = a
        walk.Next = b
        walk = a
    }        
    return res.Next
}