package easy

// @author grekz
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left != nil &&
        root.Left.Val != root.Val ||
        !isUnivalTree(root.Left) {
        return false
    }
    if root.Right != nil &&
        root.Right.Val != root.Val ||
        !isUnivalTree(root.Right) {
        return false
    }
    return true
}