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
 func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0)
}
func dfs(n *TreeNode, c int) int {
    if n == nil {
        return 0
    }
    c = c * 2 + n.Val
    if n.Left == nil && n.Right == nil {
        return c
    }
    return dfs(n.Left, c) + dfs(n.Right, c)
}