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
var minDiff int = 2147483647
var prev *TreeNode = nil
func getMinimumDifference(root *TreeNode) int {
    minDiff = 2147483647
    prev = nil
    doit(root)
    return minDiff
}
func doit(root *TreeNode ) {
    if (root == nil) { return }
    doit(root.Left)
    if (prev != nil) {
        minDiff = Min(minDiff, root.Val - prev.Val)
    }
    prev = root
    doit(root.Right)
}
func Min(a, b int) int {
    if a < b { return a}
    return b
}