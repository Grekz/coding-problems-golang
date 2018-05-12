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
var cur int = 0
func convertBST(root *TreeNode) *TreeNode {
    cur = 0
    got(root)
    return root
}
func got(x *TreeNode) {
    if x == nil {return}
    got(x.Right)
    x.Val += cur
    cur = x.Val
    got(x.Left)
}