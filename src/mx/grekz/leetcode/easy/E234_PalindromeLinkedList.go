package easy

// @author grekz
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if(head == nil || head.Next == nil) {return true}
    if(head.Next.Next == nil) {return head.Next.Val == head.Val}
    fast := head
    slow := head
    var rev *ListNode
    for ( fast != nil && fast.Next != nil ) {
        fast = fast.Next.Next
        tmp := slow
        slow = slow.Next
        tmp.Next = rev
        rev = tmp
    }
    // if is odd
    if ( fast != nil ) {
        slow = slow.Next
    }
    for ( slow != nil && rev.Val == slow.Val ){
        rev = rev.Next
        slow = slow.Next
    }
    return rev == nil
}