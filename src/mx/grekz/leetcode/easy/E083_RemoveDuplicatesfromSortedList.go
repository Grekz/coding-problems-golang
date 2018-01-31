package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil ) { return head }
    walk := head
    for(walk.Next != nil){
        if(walk.Val == walk.Next.Val){
            walk.Next = walk.Next.Next   
        } else {
            walk = walk.Next
        }
    }
    return head
}