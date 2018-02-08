package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if(root == nil) {return 0}
    return drill(root, 1)
}
func drill(walk *TreeNode, lvl int ) int{
    if(walk == nil) {return 2147483647}
    if(walk.Left == nil && walk.Right == nil) {return lvl}
    return Min(drill(walk.Left, lvl+1), drill(walk.Right, lvl+1))
}
func Min( a int, b int) int{
    if a < b {
        return a
    }
    return b
}