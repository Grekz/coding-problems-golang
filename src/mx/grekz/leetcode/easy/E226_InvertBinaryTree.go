package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if ( root == nil ) { return root }
    tmp := invertTree(root.Left)
    root.Left = invertTree(root.Right)
    root.Right = tmp
    return root 
}