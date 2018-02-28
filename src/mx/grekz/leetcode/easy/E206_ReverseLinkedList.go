package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var res, cur *ListNode = nil, nil
    for(head!=nil){
        cur=head;
        head=head.Next;
        cur.Next=res;
        res=cur;
    }
    return res;
}