package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if( l1 == nil ){ return l2 }
    if( l2 == nil ){ return l1 }
    temp := &ListNode{Val: -1}
    walk := temp
    for(l1 != nil && l2 != nil){
        if(l1.Val < l2.Val){
            walk.Next = l1
            l1 = l1.Next
        }else{
            walk.Next = l2
            l2 = l2.Next
        }
        walk = walk.Next
    }
    if(l1 == nil){ walk.Next = l2 }
    if(l2 == nil){ walk.Next = l1 }
    return temp.Next
}