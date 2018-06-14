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
 func isValidBST(root *TreeNode) bool {
    return root == nil || dfs(root, 2147483648, -2147483649)
}
func dfs( x *TreeNode, max, min int64) bool{
    return x == nil || !(int64(x.Val) >= max || int64(x.Val) <= min) && dfs(x.Left, int64(x.Val), min) && dfs(x.Right, max, int64(x.Val))
}