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
 func flipEquiv(a *TreeNode, b *TreeNode) bool {
    if a == nil || b == nil { return a == b } 
    return a.Val == b.Val && (flipEquiv(a.Left, b.Left) && flipEquiv(a.Right, b.Right) || flipEquiv(a.Right, b.Left) && flipEquiv(a.Left, b.Right))
}