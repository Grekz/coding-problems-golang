package easy

// @author grekz
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    return getKth(head, getSize(head, 0) / 2, 0)
}
func getSize(node *ListNode, n int) int {
    if node == nil {
        return n
    }
    return getSize(node.Next, n + 1)
}
func getKth(node *ListNode, k, n int) *ListNode {
    if k == n {
        return node
    }
    return getKth(node.Next, k, n + 1)
}