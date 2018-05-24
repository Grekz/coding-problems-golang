package medium

// @author grekz
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count := countit(root.Left) + 1
    if k == count {
        return root.Val
    }
    if k < count {
        return kthSmallest(root.Left, k)
    }
    return kthSmallest(root.Right, k - count)
}
func countit(n *TreeNode) int {
    if n == nil {
        return 0
    }
    return 1 + countit(n.Left) + countit(n.Right)
}