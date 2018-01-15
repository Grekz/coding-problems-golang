package hard


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    return merge( lists, 0, len(lists) - 1 )
}
func merge(lists []*ListNode, ini int, las int) *ListNode {
    if ( ini > las ) { return nil }
    if ( ini == las ) { return lists[ini] }
    mid := ( ini + las ) / 2;
    a := merge(lists, ini, mid)
    b := merge(lists, mid+1, las)
    tmp := new(ListNode)
    cur := tmp
    for ( a != nil && b != nil ) {
        if ( a.Val < b.Val ) {
            cur.Next = a;
            a = a.Next;
        }else{
            cur.Next = b;
            b = b.Next;                
        }
        cur = cur.Next;
    }
    if a != nil {
        cur.Next = a        
    }else{
        cur.Next = b
    }
    return tmp.Next
}