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
 func maxAncestorDiff(root *TreeNode) int {
    return dfs(root, root.Val, root.Val)
}
func dfs( node *TreeNode, ma, mi int) int {
    if node == nil {
        return ma - mi
    }
    ma = max(ma, node.Val)
    mi = min(mi, node.Val)
    maxL := dfs(node.Left, ma, mi)
    maxR := dfs(node.Right, ma, mi)
    return max(maxL, maxR)
}
func max( a, b int) int {
    if a < b {
        return b
    }
    return a
}
func min( a, b int) int {
    if a > b {
        return b
    }
    return a
}