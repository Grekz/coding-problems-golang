package medium

// @author grekz
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func numComponents(head *ListNode, G []int) int {
    s, res := make(map[int]bool), 0
    for _, x := range G {
        s[x] = true
    }
    for head != nil {
        if _, ok := s[head.Val]; ok {
            if head.Next == nil {
                res += 1
            } else if _, dok := s[head.Next.Val]; !dok {
                res += 1
            }
        }
        head = head.Next
    }
    return res
}