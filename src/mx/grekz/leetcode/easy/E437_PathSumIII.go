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
func pathSum(root *TreeNode, sum int) int {
    if root == nil { return 0 }
    return doit(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
func doit( n *TreeNode, s int ) int {
    if n == nil { return 0 }
    return inte( n.Val == s ) + doit(n.Left, s - n.Val ) + doit(n.Right, s - n.Val )
}
func inte( a bool ) int {
    if a {
        return 1
    }
    return 0
}