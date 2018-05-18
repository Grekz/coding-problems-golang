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
var res int = 0
func diameterOfBinaryTree(root *TreeNode) int {
    res = 0
    doit(root)
    return res
}
func doit(n *TreeNode) int {
    if n == nil {
        return 0
    }
    le := doit(n.Left)
    ri := doit(n.Right)
    res = max(res, le+ri)
    return 1 + max(le,ri)
}
func max( a, b int) int {
    if a > b {
        return a
    }
    return b
}