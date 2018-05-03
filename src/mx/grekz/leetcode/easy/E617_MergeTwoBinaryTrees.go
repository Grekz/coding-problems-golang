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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    c := TreeNode{Val: t1.Val + t2.Val }
    c.Left = mergeTrees(t1.Left, t2.Left)
    c.Right = mergeTrees(t1.Right, t2.Right)
    return &c
}