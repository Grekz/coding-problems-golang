package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if(head == nil || n < 1) {
        return head
    }
    i, size, walk := 0, 0, head
    for(walk.Next != nil){
        size++
        walk = walk.Next
    }
    index := size - (n - 1)
    if(index < 1){
        head = head.Next
        return head
    }
    walk = head
    for(walk.Next != nil){
        i++
        if(i == index){
            walk.Next = walk.Next.Next
            break
        }
        walk = walk.Next
    }
    return head
}