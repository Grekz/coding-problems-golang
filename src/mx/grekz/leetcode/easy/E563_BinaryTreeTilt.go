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
var sum int
func findTilt(root *TreeNode) int {
    sum = 0
    doit(root)
    return sum
}
func doit(node *TreeNode) int {
    if node == nil { return 0 }
    left := doit( node.Left ) 
    right := doit( node.Right )
    sum += abs( left - right )
    return node.Val + left + right
}
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}